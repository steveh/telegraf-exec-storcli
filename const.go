package main

type ListController struct {
	Number      int    `json:"Ctl"`
	Model       string `json:"Model"`
	Adapter     string `json:"AdapterType"`
	VendorID    int    `json:"VendId"`
	DeviceID    int    `json:"DevId"`
	SubVendorID int    `json:"SubVendId"`
	SubDeviceID int    `json:"SubDevId"`
	PCIAddress  string `json:"PCI Address"`
}

type ShowEnclosure struct {
	EID            int    `json:"EID"`
	State          string `json:"State"`
	Slots          int    `json:"Slots"`
	PD             int    `json:"PD"`
	PS             int    `json:"PS"`
	Fans           int    `json:"Fans"`
	TSs            int    `json:"TSs"`
	Alms           int    `json:"Alms"`
	SIM            int    `json:"SIM"`
	ProductId      string `json:"ProdID"`
	VendorSpecific string `json:"VendorSpecific"`
}

type ShowController struct {
	Number                      int
	Model                       string
	Adapter                     string
	Serial                      string
	FirmwareVersion             string
	BiosVersion                 string
	DriverVersion               string
	Status                      string
	MemoryCorrectableErrors     int
	MemoryUncorrectableErrors   int
	ROCTemperatureDegreeCelsius int
	Enclosures                  []ShowEnclosure
}

type ListAllResponse struct {
	Controllers []struct {
		CommandStatus struct {
			CLIVersion      string `json:"CLI Version"`
			OperatingSystem string `json:"Operating system"`
			StatusCode      int    `json:"Status Code"`
			Status          string `json:"Status"`
			Description     string `json:"Description"`
		} `json:"Command Status"`
		ResponseData struct {
			NumberOfControllers int              `json:"Number of Controllers"`
			HostName            string           `json:"Host Name"`
			OperatingSystem     string           `json:"Operating System "`
			StorelibITVersion   string           `json:"StoreLib IT Version"`
			StorelibIR3Version  string           `json:"StoreLib IR3 Version"`
			ITSystemOverview    []ListController `json:"IT System Overview"`
		} `json:"Response Data"`
	} `json:"Controllers"`
}

type ShowAllResponse struct {
	Controllers []struct {
		CommandStatus struct {
			CLIVersion      string `json:"CLI Version"`
			OperatingSystem string `json:"Operating system"`
			Controller      int    `json:"Controller"`
			Status          string `json:"Status"`
			Description     string `json:"Description"`
		} `json:"Command Status"`
		ResponseData struct {
			Basics struct {
				Controller                  int    `json:"Controller"`
				AdapterType                 string `json:"Adapter Type"`
				Model                       string `json:"Model"`
				SerialNumber                string `json:"Serial Number"`
				CurrentSystemDateTime       string `json:"Current System Date/time"`
				ConcurrentCommandsSupported int    `json:"Concurrent commands supported"`
				SASAddress                  string `json:"SAS Address"`
				PCIAddress                  string `json:"PCI Address"`
			} `json:"Basics"`
			Version struct {
				FirmwarePackageBuild string `json:"Firmware Package Build"`
				FirmwareVersion      string `json:"Firmware Version"`
				BiosVersion          string `json:"Bios Version"`
				NVDATAVersion        string `json:"NVDATA Version"`
				PSOCVersion          string `json:"PSOC Version"`
				DriverName           string `json:"Driver Name"`
				DriverVersion        string `json:"Driver Version"`
			} `json:"Version"`
			PciVersion struct {
				VendorID        int    `json:"Vendor Id"`
				DeviceID        int    `json:"Device Id"`
				SubVendorID     int    `json:"SubVendor Id"`
				SubDeviceID     int    `json:"SubDevice Id"`
				HostInterface   string `json:"Host Interface"`
				DeviceInterface string `json:"Device Interface"`
				BusNumber       int    `json:"Bus Number"`
				DeviceNumber    int    `json:"Device Number"`
				FunctionNumber  int    `json:"Function Number"`
				DomainID        int    `json:"Domain ID"`
			} `json:"PCI Version"`
			PendingImagesInFlash struct {
				ImageName string `json:"Image name"`
			} `json:"Pending Images in Flash"`
			Status struct {
				ControllerStatus                                string `json:"Controller Status"`
				MemoryCorrectableErrors                         int    `json:"Memory Correctable Errors"`
				MemoryUncorrectableErrors                       int    `json:"Memory Uncorrectable Errors"`
				BiosWasNotDetectedDuringBoot                    string `json:"Bios was not detected during boot"`
				ControllerHasBootedIntoSafeMode                 string `json:"Controller has booted into safe mode"`
				ControllerHasBootedIntoCertificateProvisionMode string `json:"Controller has booted into certificate provision mode"`
			} `json:"Status"`
			SupportedAdapterOperations struct {
				AlarmControl                         string `json:"Alarm Control"`
				ClusterSupport                       string `json:"Cluster Support"`
				SelfDiagnostic                       string `json:"Self Diagnostic"`
				DenySCSIPassthrough                  string `json:"Deny SCSI Passthrough"`
				DenySMPPassthrough                   string `json:"Deny SMP Passthrough"`
				DenySTPPassthrough                   string `json:"Deny STP Passthrough"`
				SupportMoreThan8PHYs                 string `json:"Support more than 8 Phys"`
				FwAndEventTimeInGMT                  string `json:"FW and Event Time in GMT"`
				SupportEnclosureEnumeration          string `json:"Support Enclosure Enumeration"`
				SupportAllowedOperations             string `json:"Support Allowed Operations"`
				SupportMultipath                     string `json:"Support Multipath"`
				SupportSecurity                      string `json:"Support Security"`
				SupportConfigPageModel               string `json:"Support Config Page Model"`
				SupportTheOCEWithoutAddingDrives     string `json:"Support the OCE without adding drives"`
				SupportEKM                           string `json:"support EKM"`
				SnapshotEnabled                      string `json:"Snapshot Enabled"`
				SupportPFK                           string `json:"Support PFK"`
				SupportPI                            string `json:"Support PI"`
				SupportShieldState                   string `json:"Support Shield State"`
				SupportSetLinkSpeed                  string `json:"Support Set Link Speed"`
				SupportJBOD                          string `json:"Support JBOD"`
				DisableOnlinePFKChange               string `json:"Disable Online PFK Change"`
				RealTimeScheduler                    string `json:"Real Time Scheduler"`
				SupportResetNow                      string `json:"Support Reset Now"`
				SupportEmulatedDrives                string `json:"Support Emulated Drives"`
				SupportSecureBoot                    string `json:"Support Secure Boot"`
				SupportPlatformSecurity              string `json:"Support Platform Security"`
				SupportPackageStampMismatchReporting string `json:"Support Package Stamp Mismatch Reporting"`
			} `json:"Supported Adapter Operations"`
			Hwcfg struct {
				ChipRevision                   string `json:"ChipRevision"`
				BatteryFRU                     string `json:"BatteryFRU"`
				FrontEndPortCount              int    `json:"Front End Port Count"`
				BackendPortCount               int    `json:"Backend Port Count"`
				SerialDebugger                 string `json:"Serial Debugger"`
				NVRAMSize                      string `json:"NVRAM Size"`
				FlashSize                      string `json:"Flash Size"`
				OnBoardMemorySize              string `json:"On Board Memory Size"`
				OnBoardExpander                string `json:"On Board Expander"`
				TemperatureSensorForROC        string `json:"Temperature Sensor for ROC"`
				TemperatureSensorForController string `json:"Temperature Sensor for Controller"`
				CurrentSizeOfCacheCadeGB       int    `json:"Current Size of CacheCade (GB)"`
				CurrentSizeOfFWCacheMB         int    `json:"Current Size of FW Cache (MB)"`
				ROCTemperatureDegreeCelsius    int    `json:"ROC temperature(Degree Celsius)"`
			} `json:"HwCfg"`
			Policies struct {
				PoliciesTable []struct {
					Policy  string `json:"Policy"`
					Current string `json:"Current"`
					Default string `json:"Default"`
				} `json:"Policies Table"`
				FlushTimeDefault             string `json:"Flush Time(Default)"`
				DriveCoercionMode            string `json:"Drive Coercion Mode"`
				AutoRebuild                  string `json:"Auto Rebuild"`
				BatteryWarning               string `json:"Battery Warning"`
				ECCBucketSize                int    `json:"ECC Bucket Size"`
				ECCBucketLeakRateHours       int    `json:"ECC Bucket Leak Rate (hrs)"`
				RestoreHotspareOnInsertion   string `json:"Restore HotSpare on Insertion"`
				ExposeEnclosureDevices       string `json:"Expose Enclosure Devices"`
				MaintainPDFailHistory        string `json:"Maintain PD Fail History"`
				ReorderHostRequests          string `json:"Reorder Host Requests"`
				AutoDetectBackplane          string `json:"Auto detect BackPlane"`
				LoadBalanceMode              string `json:"Load Balance Mode"`
				SecurityKeyAssigned          string `json:"Security Key Assigned"`
				DisableOnlineControllerReset string `json:"Disable Online Controller Reset"`
				UseDriveActivityForLocate    string `json:"Use drive activity for locate"`
			} `json:"Policies"`
			Boot struct {
				MaxDrivesToSpinupAtOneTime                        int    `json:"Max Drives to Spinup at One Time"`
				MaximumNumberOfDirectAttachedDrivesToSpinUpIn1Min int    `json:"Maximum number of direct attached drives to spin up in 1 min"`
				DelayAmongSpinupGroupsSec                         int    `json:"Delay Among Spinup Groups (sec)"`
				AllowBootWithPreservedCache                       string `json:"Allow Boot with Preserved Cache"`
			} `json:"Boot"`
			Defaults struct {
				PHYPolarity                   int    `json:"Phy Polarity"`
				PHYPolaritysplit              int    `json:"Phy PolaritySplit"`
				CachedIO                      string `json:"Cached IO"`
				DefaultSpinDownTimeMins       int    `json:"Default spin down time (mins)"`
				CoercionMode                  string `json:"Coercion Mode"`
				ZCRConfig                     string `json:"ZCR Config"`
				MaxChainedEnclosures          int    `json:"Max Chained Enclosures"`
				DirectPDMapping               string `json:"Direct PD Mapping"`
				RestoreHotSpareOnInsertion    string `json:"Restore Hot Spare on Insertion"`
				ExposeEnclosureDevices        string `json:"Expose Enclosure Devices"`
				MaintainPDFailHistory         string `json:"Maintain PD Fail History"`
				ZeroBasedEnclosureEnumeration string `json:"Zero Based Enclosure Enumeration"`
				DisablePuncturing             string `json:"Disable Puncturing"`
				UnCertifiedHardDiskDrives     string `json:"Un-Certified Hard Disk Drives"`
				SMARTMode                     string `json:"SMART Mode"`
				EnableLEDHeader               string `json:"Enable LED Header"`
				LEDShowDriveActivity          string `json:"LED Show Drive Activity"`
				DirtyLEDShowsDriveActivity    string `json:"Dirty LED Shows Drive Activity"`
				EnableCrashDump               string `json:"EnableCrashDump"`
				DisableOnlineControllerReset  string `json:"Disable Online Controller Reset"`
				TreatSingleSpanR1EAsR10       string `json:"Treat Single span R1E as R10"`
				PowerSavingOption             string `json:"Power Saving option"`
				TTYLogInFlash                 string `json:"TTY Log In Flash"`
				AutoEnhancedImport            string `json:"Auto Enhanced Import"`
				EnableShieldState             string `json:"Enable Shield State"`
				TimeTakenToDetectCME          string `json:"Time taken to detect CME"`
			} `json:"Defaults"`
			Capabilities struct {
				SupportedDrives              string `json:"Supported Drives"`
				EnableJBOD                   string `json:"Enable JBOD"`
				MaxParallelCommands          int    `json:"Max Parallel Commands"`
				MaxSGECount                  int    `json:"Max SGE Count"`
				MaxDataTransferSize          string `json:"Max Data Transfer Size"`
				MaxStripsPerIO               int    `json:"Max Strips PerIO"`
				MaxConfigurableCacheCadeSize int    `json:"Max Configurable CacheCade Size"`
				MinStripSize                 string `json:"Min Strip Size"`
				MaxStripSize                 string `json:"Max Strip Size"`
			} `json:"Capabilities"`
			ScheduledTasks             string `json:"Scheduled Tasks"`
			SecurityProtocolProperties struct {
				SecurityProtocol string `json:"Security Protocol"`
			} `json:"Security Protocol properties"`
			EnclosureInformation []ShowEnclosure `json:"Enclosure Information"`
		} `json:"Response Data"`
	} `json:"Controllers"`
}
