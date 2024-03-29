package tests

import (
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/test-network-function/cnfcert-tests-verification/tests/globalhelper"
	"github.com/test-network-function/cnfcert-tests-verification/tests/globalparameters"
	"github.com/test-network-function/cnfcert-tests-verification/tests/utils/deployment"
	"github.com/test-network-function/cnfcert-tests-verification/tests/utils/nodes"

	tshelper "github.com/test-network-function/cnfcert-tests-verification/tests/lifecycle/helper"
	tsparams "github.com/test-network-function/cnfcert-tests-verification/tests/lifecycle/parameters"
)

var _ = Describe("lifecycle-pod-recreation", Serial, func() {
	var randomNamespace string
	var randomReportDir string
	var randomTnfConfigDir string

	BeforeEach(func() {
		By("Enable intrusive tests")
		err := os.Setenv("TNF_NON_INTRUSIVE_ONLY", "false")
		Expect(err).ToNot(HaveOccurred())

		if globalhelper.IsKindCluster() {
			By("Make masters schedulable")
			err := nodes.EnableMasterScheduling(globalhelper.GetAPIClient().Nodes(), true)
			Expect(err).ToNot(HaveOccurred())
		}

		// Create random namespace and keep original report and TNF config directories
		randomNamespace, randomReportDir, randomTnfConfigDir = globalhelper.BeforeEachSetupWithRandomNamespace(tsparams.LifecycleNamespace)

		By("Define TNF config file")
		err = globalhelper.DefineTnfConfig(
			[]string{randomNamespace},
			[]string{tsparams.TestPodLabel},
			[]string{tsparams.TnfTargetOperatorLabels},
			[]string{},
			[]string{}, randomTnfConfigDir)
		Expect(err).ToNot(HaveOccurred())

		if globalhelper.GetConfiguration().General.DisableIntrusiveTests == strings.ToLower("true") {
			Skip("Intrusive tests are disabled via config")
		}
	})

	AfterEach(func() {
		By("Disable intrusive tests")
		err := os.Setenv("TNF_NON_INTRUSIVE_ONLY", "true")
		Expect(err).ToNot(HaveOccurred())

		globalhelper.AfterEachCleanupWithRandomNamespace(randomNamespace, randomReportDir, randomTnfConfigDir, tsparams.WaitingTime)
	})

	// 47405
	It("One deployment with PodAntiAffinity, replicas are less than schedulable nodes", func() {
		schedulableNodes, err := nodes.GetNumOfReadyNodesInCluster(globalhelper.GetAPIClient().Nodes())
		Expect(err).ToNot(HaveOccurred())

		if schedulableNodes == 0 {
			Skip("The cluster does not have enough schedulable nodes.")
		}
		// at least one "clean of any resource" worker is needed.
		maxPodsPerDeployment := schedulableNodes - 1

		By("Define deployment")
		deploymenta, err := tshelper.DefineDeployment(maxPodsPerDeployment, 1, tsparams.TestDeploymentName, randomNamespace)
		Expect(err).ToNot(HaveOccurred())

		deployment.RedefineWithPodAntiAffinity(deploymenta, tsparams.TestTargetLabels)

		By("Create deployment")
		err = globalhelper.CreateAndWaitUntilDeploymentIsReady(deploymenta, tsparams.WaitingTime)
		Expect(err).ToNot(HaveOccurred())

		By("Assert deployment has pod anti-affinity")
		runningDeployment, err := globalhelper.GetRunningDeployment(deploymenta.Namespace, deploymenta.Name)
		Expect(err).ToNot(HaveOccurred())
		Expect(runningDeployment.Spec.Template.Spec.Affinity.PodAntiAffinity).ToNot(BeNil())

		By("Start lifecycle-pod-recreation test")
		err = globalhelper.LaunchTests(tsparams.TnfPodRecreationTcName,
			globalhelper.ConvertSpecNameToFileName(CurrentSpecReport().FullText()), randomReportDir, randomTnfConfigDir)
		Expect(err).ToNot(HaveOccurred())

		By("Verify test case status in Claim report")
		err = globalhelper.ValidateIfReportsAreValid(tsparams.TnfPodRecreationTcName, globalparameters.TestCasePassed, randomReportDir)
		Expect(err).ToNot(HaveOccurred())
	})

	// 47406
	It("Two deployments with PodAntiAffinity, replicas are less than schedulable nodes", func() {
		schedulableNodes, err := nodes.GetNumOfReadyNodesInCluster(globalhelper.GetAPIClient().Nodes())
		Expect(err).ToNot(HaveOccurred())

		if schedulableNodes < 4 {
			Skip("The cluster does not have enough schedulable nodes.")
		}
		// at least one "clean of any resource" worker is needed.
		maxPodsPerDeployment := (schedulableNodes / 2) - 1

		By("Define and create first deployment")
		deploymenta, err := tshelper.DefineDeployment(maxPodsPerDeployment, 1, tsparams.TestDeploymentName, randomNamespace)
		Expect(err).ToNot(HaveOccurred())

		deployment.RedefineWithPodAntiAffinity(deploymenta, tsparams.TestTargetLabels)

		By("Create deployment 1")
		err = globalhelper.CreateAndWaitUntilDeploymentIsReady(deploymenta, tsparams.WaitingTime)
		Expect(err).ToNot(HaveOccurred())

		By("Assert deploymenta has pod anti-affinity")
		runningDeployment, err := globalhelper.GetRunningDeployment(deploymenta.Namespace, deploymenta.Name)
		Expect(err).ToNot(HaveOccurred())
		Expect(runningDeployment.Spec.Template.Spec.Affinity.PodAntiAffinity).ToNot(BeNil())

		By("Define and create second deployment")
		deploymentb, err := tshelper.DefineDeployment(maxPodsPerDeployment, 1, "lifecycle-dpb", randomNamespace)
		Expect(err).ToNot(HaveOccurred())

		deployment.RedefineWithPodAntiAffinity(deploymentb, tsparams.TestTargetLabels)

		By("Create deployment 2")
		err = globalhelper.CreateAndWaitUntilDeploymentIsReady(deploymentb, tsparams.WaitingTime)
		Expect(err).ToNot(HaveOccurred())

		By("Assert deploymentb has pod anti-affinity")
		runningDeployment, err = globalhelper.GetRunningDeployment(deploymentb.Namespace, deploymentb.Name)
		Expect(err).ToNot(HaveOccurred())
		Expect(runningDeployment.Spec.Template.Spec.Affinity.PodAntiAffinity).ToNot(BeNil())

		By("Start lifecycle-pod-recreation test")
		err = globalhelper.LaunchTests(tsparams.TnfPodRecreationTcName,
			globalhelper.ConvertSpecNameToFileName(CurrentSpecReport().FullText()), randomReportDir, randomTnfConfigDir)
		Expect(err).ToNot(HaveOccurred())

		By("Verify test case status in Claim report")
		err = globalhelper.ValidateIfReportsAreValid(tsparams.TnfPodRecreationTcName, globalparameters.TestCasePassed, randomReportDir)
		Expect(err).ToNot(HaveOccurred())
	})

	// 47407
	It("One deployment with PodAntiAffinity, replicas are equal to schedulable nodes [negative]", func() {
		schedulableNodes, err := nodes.GetNumOfReadyNodesInCluster(globalhelper.GetAPIClient().Nodes())
		Expect(err).ToNot(HaveOccurred())

		if schedulableNodes == 0 {
			Skip("The cluster does not have enough schedulable nodes.")
		}

		By("Define deployment")
		deploymenta, err := tshelper.DefineDeployment(schedulableNodes, 1, tsparams.TestDeploymentName, randomNamespace)
		Expect(err).ToNot(HaveOccurred())

		deployment.RedefineWithPodAntiAffinity(deploymenta, tsparams.TestTargetLabels)

		By("Create deployment")
		err = globalhelper.CreateAndWaitUntilDeploymentIsReady(deploymenta, tsparams.WaitingTime)
		Expect(err).ToNot(HaveOccurred())

		By("Assert deployment has pod anti-affinity")
		runningDeployment, err := globalhelper.GetRunningDeployment(deploymenta.Namespace, deploymenta.Name)
		Expect(err).ToNot(HaveOccurred())
		Expect(runningDeployment.Spec.Template.Spec.Affinity.PodAntiAffinity).ToNot(BeNil())

		By("Start lifecycle-pod-recreation test")
		err = globalhelper.LaunchTests(tsparams.TnfPodRecreationTcName,
			globalhelper.ConvertSpecNameToFileName(CurrentSpecReport().FullText()), randomReportDir, randomTnfConfigDir)
		Expect(err).ToNot(HaveOccurred())

		By("Verify test case status in Claim report")
		err = globalhelper.ValidateIfReportsAreValid(tsparams.TnfPodRecreationTcName, globalparameters.TestCaseFailed, randomReportDir)
		Expect(err).ToNot(HaveOccurred())
	})

	// 47408
	It("Two deployments with PodAntiAffinity, replicas are equal to schedulable nodes [negative]", func() {
		schedulableNodes, err := nodes.GetNumOfReadyNodesInCluster(globalhelper.GetAPIClient().Nodes())
		Expect(err).ToNot(HaveOccurred())

		// all nodes will be scheduled with a pod.
		if schedulableNodes < 2 {
			Skip("The cluster does not have enough schedulable nodes.")
		}
		maxPodsPerDeploymentPerFirstDeployment := (schedulableNodes / 2)
		maxPodsPerDeploymentPerSecondDeployment := schedulableNodes - maxPodsPerDeploymentPerFirstDeployment

		By("Define and create first deployment")
		deploymenta, err := tshelper.DefineDeployment(maxPodsPerDeploymentPerFirstDeployment, 1,
			tsparams.TestDeploymentName, randomNamespace)
		Expect(err).ToNot(HaveOccurred())

		deployment.RedefineWithPodAntiAffinity(deploymenta, tsparams.TestTargetLabels)

		By("Create deployment 1")
		err = globalhelper.CreateAndWaitUntilDeploymentIsReady(deploymenta, tsparams.WaitingTime)
		Expect(err).ToNot(HaveOccurred())

		By("Assert deploymenta has pod anti-affinity")
		runningDeployment, err := globalhelper.GetRunningDeployment(deploymenta.Namespace, deploymenta.Name)
		Expect(err).ToNot(HaveOccurred())
		Expect(runningDeployment.Spec.Template.Spec.Affinity.PodAntiAffinity).ToNot(BeNil())

		By("Define and create second deployment")
		deploymentb, err := tshelper.DefineDeployment(maxPodsPerDeploymentPerSecondDeployment, 1,
			"lifecycle-dpb", randomNamespace)
		Expect(err).ToNot(HaveOccurred())

		deployment.RedefineWithPodAntiAffinity(deploymentb, tsparams.TestTargetLabels)

		By("Create deployment 2")
		err = globalhelper.CreateAndWaitUntilDeploymentIsReady(deploymentb, tsparams.WaitingTime)
		Expect(err).ToNot(HaveOccurred())

		By("Assert deploymentb has pod anti-affinity")
		runningDeployment, err = globalhelper.GetRunningDeployment(deploymentb.Namespace, deploymentb.Name)
		Expect(err).ToNot(HaveOccurred())
		Expect(runningDeployment.Spec.Template.Spec.Affinity.PodAntiAffinity).ToNot(BeNil())

		By("Start lifecycle-pod-recreation test")
		err = globalhelper.LaunchTests(tsparams.TnfPodRecreationTcName,
			globalhelper.ConvertSpecNameToFileName(CurrentSpecReport().FullText()), randomReportDir, randomTnfConfigDir)
		Expect(err).ToNot(HaveOccurred())

		By("Verify test case status in Claim report")
		err = globalhelper.ValidateIfReportsAreValid(tsparams.TnfPodRecreationTcName, globalparameters.TestCaseFailed, randomReportDir)
		Expect(err).ToNot(HaveOccurred())
	})
})
