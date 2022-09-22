package tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"

	"github.com/test-network-function/cnfcert-tests-verification/tests/globalhelper"
	"github.com/test-network-function/cnfcert-tests-verification/tests/globalparameters"
	tshelper "github.com/test-network-function/cnfcert-tests-verification/tests/lifecycle/helper"
	tsparams "github.com/test-network-function/cnfcert-tests-verification/tests/lifecycle/parameters"
	"github.com/test-network-function/cnfcert-tests-verification/tests/utils/deployment"
	"github.com/test-network-function/cnfcert-tests-verification/tests/utils/namespaces"
	"github.com/test-network-function/cnfcert-tests-verification/tests/utils/persistentvolume"
	"github.com/test-network-function/cnfcert-tests-verification/tests/utils/persistentvolumeclaim"
)

var _ = Describe("lifecycle-persistent-volume-reclaim-policy", func() {

	BeforeEach(func() {
		err := tshelper.WaitUntilClusterIsStable()
		Expect(err).ToNot(HaveOccurred())

		By("Clean namespace before each test")
		err = namespaces.Clean(tsparams.LifecycleNamespace, globalhelper.APIClient)
		Expect(err).ToNot(HaveOccurred())

		By("Delete all existing PVs")
		err = tshelper.DeletePVs()
		Expect(err).ToNot(HaveOccurred())

	})

	// // 54201
	// It("One deployment, one pod with a volume that uses a reclaim policy of delete", func() {

	// 	persistentVolume := persistentvolume.RedefineWithPVReclaimPolicy(
	// 		persistentvolume.DefinePersistentVolume(tsparams.TestPVName, tsparams.LifecycleNamespace), corev1.PersistentVolumeReclaimDelete)

	// 	err := tshelper.CreatePersistentVolume(persistentVolume, tsparams.WaitingTime)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	pvc := persistentvolumeclaim.DefinePersistentVolumeClaim(tsparams.TestPVCName, tsparams.LifecycleNamespace)

	// 	err = tshelper.CreateAndWaitUntilPVCIsBound(pvc, tsparams.LifecycleNamespace, tsparams.WaitingTime, persistentVolume.Name)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Define deployment")
	// 	dep := deployment.DefineDeployment(tsparams.TestDeploymentName, tsparams.LifecycleNamespace,
	// 		globalhelper.Configuration.General.TestImage, tsparams.TestTargetLabels)

	// 	dep = deployment.RedefineWithPVC(dep, tsparams.TestVolumeName, tsparams.TestPVCName)

	// 	err = globalhelper.CreateAndWaitUntilDeploymentIsReady(dep, tsparams.WaitingTime)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Start lifecycle-persistent-volume-reclaim-policy test")
	// 	err = globalhelper.LaunchTests(tsparams.TnfPersistentVolumeReclaimPolicy,
	// 		globalhelper.ConvertSpecNameToFileName(CurrentSpecReport().FullText()))
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Verify test case status in Junit and Claim reports")
	// 	err = globalhelper.ValidateIfReportsAreValid(tsparams.TnfPersistentVolumeReclaimPolicy, globalparameters.TestCasePassed)
	// 	Expect(err).ToNot(HaveOccurred())
	// })

	// // 54202
	// It("One pod with a volume that uses a reclaim policy of delete", func() {

	// 	persistentVolume := persistentvolume.RedefineWithPVReclaimPolicy(
	// 		persistentvolume.DefinePersistentVolume(tsparams.TestPVName, tsparams.LifecycleNamespace), corev1.PersistentVolumeReclaimDelete)

	// 	err := tshelper.CreatePersistentVolume(persistentVolume, tsparams.WaitingTime)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	pvc := persistentvolumeclaim.DefinePersistentVolumeClaim(tsparams.TestPVCName, tsparams.LifecycleNamespace)

	// 	err = tshelper.CreateAndWaitUntilPVCIsBound(pvc, tsparams.LifecycleNamespace, tsparams.WaitingTime, persistentVolume.Name)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Define pod")
	// 	put := pod.RedefineWithPVC(pod.DefinePod(tsparams.TestPodName, tsparams.LifecycleNamespace,
	// 		globalhelper.Configuration.General.TestImage), tsparams.TestVolumeName, tsparams.TestPVCName)

	// 	put = pod.RedefinePodWithLabel(put, tsparams.TestTargetLabels)

	// 	err = globalhelper.CreateAndWaitUntilPodIsReady(put, tsparams.WaitingTime)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Start lifecycle-persistent-volume-reclaim-policy test")
	// 	err = globalhelper.LaunchTests(tsparams.TnfPersistentVolumeReclaimPolicy,
	// 		globalhelper.ConvertSpecNameToFileName(CurrentSpecReport().FullText()))
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Verify test case status in Junit and Claim reports")
	// 	err = globalhelper.ValidateIfReportsAreValid(tsparams.TnfPersistentVolumeReclaimPolicy, globalparameters.TestCasePassed)
	// 	Expect(err).ToNot(HaveOccurred())

	// })

	// // 54203
	// It("One replicaSet with a volume that uses a reclaim policy of delete", func() {

	// 	persistentVolume := persistentvolume.RedefineWithPVReclaimPolicy(
	// 		persistentvolume.DefinePersistentVolume(tsparams.TestPVName, tsparams.LifecycleNamespace), corev1.PersistentVolumeReclaimDelete)

	// 	err := tshelper.CreatePersistentVolume(persistentVolume, tsparams.WaitingTime)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	pvc := persistentvolumeclaim.DefinePersistentVolumeClaim(tsparams.TestPVCName, tsparams.LifecycleNamespace)

	// 	err = tshelper.CreateAndWaitUntilPVCIsBound(pvc, tsparams.LifecycleNamespace, tsparams.WaitingTime, persistentVolume.Name)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Define replicaSet")
	// 	rs := replicaset.RedefineWithPVC(replicaset.DefineReplicaSet(tsparams.TestReplicaSetName, tsparams.LifecycleNamespace,
	// 		globalhelper.Configuration.General.TestImage, tsparams.TestTargetLabels), tsparams.TestVolumeName, tsparams.TestPVCName)

	// 	err = tshelper.CreateAndWaitUntilReplicaSetIsReady(rs, tsparams.WaitingTime)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Start lifecycle-persistent-volume-reclaim-policy test")
	// 	err = globalhelper.LaunchTests(tsparams.TnfPersistentVolumeReclaimPolicy,
	// 		globalhelper.ConvertSpecNameToFileName(CurrentSpecReport().FullText()))
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Verify test case status in Junit and Claim reports")
	// 	err = globalhelper.ValidateIfReportsAreValid(tsparams.TnfPersistentVolumeReclaimPolicy, globalparameters.TestCasePassed)
	// 	Expect(err).ToNot(HaveOccurred())

	// })

	// // 54204
	// It("One deployment, one pod with a volume that uses a reclaim policy of retain [negative]", func() {

	// 	persistentVolume := persistentvolume.RedefineWithPVReclaimPolicy(
	// 		persistentvolume.DefinePersistentVolume(tsparams.TestPVName, tsparams.LifecycleNamespace), corev1.PersistentVolumeReclaimRetain)

	// 	err := tshelper.CreatePersistentVolume(persistentVolume, tsparams.WaitingTime)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	pvc := persistentvolumeclaim.DefinePersistentVolumeClaim(tsparams.TestPVCName, tsparams.LifecycleNamespace)

	// 	err = tshelper.CreateAndWaitUntilPVCIsBound(pvc, tsparams.LifecycleNamespace, tsparams.WaitingTime, persistentVolume.Name)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Define deployment")
	// 	dep := deployment.DefineDeployment(tsparams.TestDeploymentName, tsparams.LifecycleNamespace,
	// 		globalhelper.Configuration.General.TestImage, tsparams.TestTargetLabels)

	// 	dep = deployment.RedefineWithPVC(dep, tsparams.TestVolumeName, tsparams.TestPVCName)

	// 	err = globalhelper.CreateAndWaitUntilDeploymentIsReady(dep, tsparams.WaitingTime)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Start lifecycle-persistent-volume-reclaim-policy test")
	// 	err = globalhelper.LaunchTests(tsparams.TnfPersistentVolumeReclaimPolicy,
	// 		globalhelper.ConvertSpecNameToFileName(CurrentSpecReport().FullText()))
	// 	Expect(err).To(HaveOccurred())

	// 	By("Verify test case status in Junit and Claim reports")
	// 	err = globalhelper.ValidateIfReportsAreValid(tsparams.TnfPersistentVolumeReclaimPolicy, globalparameters.TestCaseFailed)
	// 	Expect(err).ToNot(HaveOccurred())
	// })

	// // 54206
	// It("One pod with a volume that uses a reclaim policy of recycle [negative]", func() {

	// 	persistentVolume := persistentvolume.RedefineWithPVReclaimPolicy(
	// 		persistentvolume.DefinePersistentVolume(tsparams.TestPVName, tsparams.LifecycleNamespace), corev1.PersistentVolumeReclaimRecycle)

	// 	err := tshelper.CreatePersistentVolume(persistentVolume, tsparams.WaitingTime)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	pvc := persistentvolumeclaim.DefinePersistentVolumeClaim(tsparams.TestPVCName, tsparams.LifecycleNamespace)

	// 	err = tshelper.CreateAndWaitUntilPVCIsBound(pvc, tsparams.LifecycleNamespace, tsparams.WaitingTime, persistentVolume.Name)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Define pod")
	// 	put := pod.RedefineWithPVC(pod.DefinePod(tsparams.TestPodName, tsparams.LifecycleNamespace,
	// 		globalhelper.Configuration.General.TestImage), tsparams.TestVolumeName, tsparams.TestPVCName)

	// 	put = pod.RedefinePodWithLabel(put, tsparams.TestTargetLabels)

	// 	err = globalhelper.CreateAndWaitUntilPodIsReady(put, tsparams.WaitingTime)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	By("Start lifecycle-persistent-volume-reclaim-policy test")
	// 	err = globalhelper.LaunchTests(tsparams.TnfPersistentVolumeReclaimPolicy,
	// 		globalhelper.ConvertSpecNameToFileName(CurrentSpecReport().FullText()))
	// 	Expect(err).To(HaveOccurred())

	// 	By("Verify test case status in Junit and Claim reports")
	// 	err = globalhelper.ValidateIfReportsAreValid(tsparams.TnfPersistentVolumeReclaimPolicy, globalparameters.TestCaseFailed)
	// 	Expect(err).ToNot(HaveOccurred())

	// })

	// 54207
	It("Two deployments, one with reclaim policy of delete, other with recycle [negative]", func() {

		By("Define & create first pv")
		persistentVolumea := persistentvolume.RedefineWithPVReclaimPolicy(
			persistentvolume.DefinePersistentVolume(tsparams.TestPVName, tsparams.LifecycleNamespace), corev1.PersistentVolumeReclaimDelete)

		err := tshelper.CreatePersistentVolume(persistentVolumea, tsparams.WaitingTime)
		Expect(err).ToNot(HaveOccurred())

		By("Define & create second pv")
		persistentVolumeb := persistentvolume.RedefineWithPVReclaimPolicy(
			persistentvolume.DefinePersistentVolume("lifecycle-pvb", tsparams.LifecycleNamespace), corev1.PersistentVolumeReclaimRecycle)

		err = tshelper.CreatePersistentVolume(persistentVolumeb, tsparams.WaitingTime)
		Expect(err).ToNot(HaveOccurred())

		By("Define & create first pvc")
		pvca := persistentvolumeclaim.DefinePersistentVolumeClaim(tsparams.TestPVCName, tsparams.LifecycleNamespace)

		err = tshelper.CreateAndWaitUntilPVCIsBound(pvca, tsparams.LifecycleNamespace, tsparams.WaitingTime, persistentVolumea.Name)
		Expect(err).ToNot(HaveOccurred())

		By("Define & create second pvc")
		pvcb := persistentvolumeclaim.DefinePersistentVolumeClaim("lifecycle-pvcb", tsparams.LifecycleNamespace)

		err = tshelper.CreateAndWaitUntilPVCIsBound(pvcb, tsparams.LifecycleNamespace, tsparams.WaitingTime, persistentVolumeb.Name)
		Expect(err).ToNot(HaveOccurred())

		By("Define deployments")
		depa := deployment.DefineDeployment(tsparams.TestDeploymentName, tsparams.LifecycleNamespace,
			globalhelper.Configuration.General.TestImage, tsparams.TestTargetLabels)

		depa = deployment.RedefineWithPVC(depa, tsparams.TestVolumeName, tsparams.TestPVCName)

		depb := deployment.DefineDeployment("lifecycle-dpb", tsparams.LifecycleNamespace, globalhelper.Configuration.General.TestImage,
			tsparams.TestTargetLabels)

		depb = deployment.RedefineWithPVC(depb, tsparams.TestVolumeName, "lifecycle-pvcb")

		err = globalhelper.CreateAndWaitUntilDeploymentIsReady(depa, tsparams.WaitingTime)
		Expect(err).ToNot(HaveOccurred())

		err = globalhelper.CreateAndWaitUntilDeploymentIsReady(depb, tsparams.WaitingTime)
		Expect(err).ToNot(HaveOccurred())

		By("Start lifecycle-persistent-volume-reclaim-policy test")
		err = globalhelper.LaunchTests(tsparams.TnfPersistentVolumeReclaimPolicy,
			globalhelper.ConvertSpecNameToFileName(CurrentSpecReport().FullText()))
		Expect(err).To(HaveOccurred())

		By("Verify test case status in Junit and Claim reports")
		err = globalhelper.ValidateIfReportsAreValid(tsparams.TnfPersistentVolumeReclaimPolicy, globalparameters.TestCaseFailed)
		Expect(err).ToNot(HaveOccurred())

	})
})