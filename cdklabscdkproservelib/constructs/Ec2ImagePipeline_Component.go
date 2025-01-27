package constructs


// Image Builder Component.
// Experimental.
type Ec2ImagePipeline_Component string

const (
	// Installs the latest version of the Amazon CloudWatch agent.
	//
	// This component installs only the agent. You must take additional steps to configure and use the Amazon CloudWatch agent. For more information, see the documentation at https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/install-CloudWatch-Agent-on-EC2-Instance.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CLOUDWATCH_AGENT_LINUX Ec2ImagePipeline_Component = "AMAZON_CLOUDWATCH_AGENT_LINUX"
	// Installs the latest version of the Amazon CloudWatch agent.
	//
	// This component installs only the agent. You must take additional steps to configure and use the Amazon CloudWatch agent. For more information, see the documentation at https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/install-CloudWatch-Agent-on-EC2-Instance.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CLOUDWATCH_AGENT_WINDOWS Ec2ImagePipeline_Component = "AMAZON_CLOUDWATCH_AGENT_WINDOWS"
	// Installs Amazon Corretto 11 for Debian-based Linux platforms in accordance with the Amazon Corretto 11 User Guide at https://docs.aws.amazon.com/corretto/latest/corretto-11-ug/generic-linux-install.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_11_APT_GENERIC Ec2ImagePipeline_Component = "AMAZON_CORRETTO_11_APT_GENERIC"
	// Installs Amazon Corretto 11 Headless.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_11_HEADLESS Ec2ImagePipeline_Component = "AMAZON_CORRETTO_11_HEADLESS"
	// Installs Amazon Corretto 11 for RPM-based Linux platforms in accordance with the Amazon Corretto 11 User Guide at https://docs.aws.amazon.com/corretto/latest/corretto-11-ug/generic-linux-install.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_11_RPM_GENERIC Ec2ImagePipeline_Component = "AMAZON_CORRETTO_11_RPM_GENERIC"
	// Installs Amazon Corretto 11 for Windows in accordance with the Amazon Corretto 11 User Guide at https://docs.aws.amazon.com/corretto/latest/corretto-11-ug/windows-7-install.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_11_WINDOWS Ec2ImagePipeline_Component = "AMAZON_CORRETTO_11_WINDOWS"
	// Installs Amazon Corretto 11.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_11 Ec2ImagePipeline_Component = "AMAZON_CORRETTO_11"
	// Installs Amazon Corretto 17 Headless.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_17_HEADLESS Ec2ImagePipeline_Component = "AMAZON_CORRETTO_17_HEADLESS"
	// Installs Amazon Corretto 17 JDK in accordance with the Amazon Corretto 17 User Guide at https://docs.aws.amazon.com/corretto/latest/corretto-17-ug/linux-info.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_17_JDK Ec2ImagePipeline_Component = "AMAZON_CORRETTO_17_JDK"
	// Installs Amazon Corretto 17 JRE.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_17_JRE Ec2ImagePipeline_Component = "AMAZON_CORRETTO_17_JRE"
	// Installs Amazon Corretto 17 for Windows in accordance with the Amazon Corretto 17 User Guide at https://docs.aws.amazon.com/corretto/latest/corretto-17-ug/windows-7-install.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_17_WINDOWS Ec2ImagePipeline_Component = "AMAZON_CORRETTO_17_WINDOWS"
	// Installs Amazon Corretto 21 Headless.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_21_HEADLESS Ec2ImagePipeline_Component = "AMAZON_CORRETTO_21_HEADLESS"
	// Installs Amazon Corretto 21 JDK in accordance with the Amazon Corretto 21 User Guide at https://docs.aws.amazon.com/corretto/latest/corretto-21-ug/linux-info.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_21_JDK Ec2ImagePipeline_Component = "AMAZON_CORRETTO_21_JDK"
	// Installs Amazon Corretto 21 JRE.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_21_JRE Ec2ImagePipeline_Component = "AMAZON_CORRETTO_21_JRE"
	// Installs Amazon Corretto 21 for Windows in accordance with the Amazon Corretto 21 User Guide at https://docs.aws.amazon.com/corretto/latest/corretto-21-ug/windows-10-install.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_21_WINDOWS Ec2ImagePipeline_Component = "AMAZON_CORRETTO_21_WINDOWS"
	// Installs Amazon Corretto 8 for Debian-based Linux platforms in accordance with the Amazon Corretto 8 User Guide at https://docs.aws.amazon.com/corretto/latest/corretto-8-ug/generic-linux-install.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_8_APT_GENERIC Ec2ImagePipeline_Component = "AMAZON_CORRETTO_8_APT_GENERIC"
	// Installs Amazon Corretto 8 JDK.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_8_JDK Ec2ImagePipeline_Component = "AMAZON_CORRETTO_8_JDK"
	// Installs Amazon Corretto 8 JRE.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_8_JRE Ec2ImagePipeline_Component = "AMAZON_CORRETTO_8_JRE"
	// Installs Amazon Corretto 8 for RPM-based Linux platforms in accordance with the Amazon Corretto 8 User Guide at https://docs.aws.amazon.com/corretto/latest/corretto-8-ug/generic-linux-install.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_8_RPM_GENERIC Ec2ImagePipeline_Component = "AMAZON_CORRETTO_8_RPM_GENERIC"
	// Installs Amazon Corretto 8 for Windows in accordance with the Amazon Corretto 8 User Guide at https://docs.aws.amazon.com/corretto/latest/corretto-8-ug/windows-7-install.html.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_CORRETTO_8_WINDOWS Ec2ImagePipeline_Component = "AMAZON_CORRETTO_8_WINDOWS"
	// Installs the latest version of Amazon Kinesis Agent for Windows.
	// Experimental.
	Ec2ImagePipeline_Component_AMAZON_KINESIS_AGENT_WINDOWS Ec2ImagePipeline_Component = "AMAZON_KINESIS_AGENT_WINDOWS"
	// Installs the Anaconda distribution and environments for Tensorflow, PyTorch, and MXNet.
	// Experimental.
	Ec2ImagePipeline_Component_ANACONDA_WINDOWS Ec2ImagePipeline_Component = "ANACONDA_WINDOWS"
	// Installs the latest version of Apache Tomcat and the JRE, sets required environment variables, and schedules Tomcat to run on startup.
	// Experimental.
	Ec2ImagePipeline_Component_APACHE_TOMCAT_9_LINUX Ec2ImagePipeline_Component = "APACHE_TOMCAT_9_LINUX"
	// Tests whether the apt package manager is functioning correctly.
	// Experimental.
	Ec2ImagePipeline_Component_APT_REPOSITORY_TEST_LINUX Ec2ImagePipeline_Component = "APT_REPOSITORY_TEST_LINUX"
	// Installs the latest version of the AWS CLI version 2, and creates the symlink /usr/bin/aws that points to the installed application.
	//
	// For more information, see https://docs.aws.amazon.com/cli/latest/userguide/.
	// Experimental.
	Ec2ImagePipeline_Component_AWS_CLI_VERSION_2_LINUX Ec2ImagePipeline_Component = "AWS_CLI_VERSION_2_LINUX"
	// Installs the latest version of the AWS CLI version 2.
	//
	// For more information, review the user guide at https://docs.aws.amazon.com/cli/latest/userguide/.
	// Experimental.
	Ec2ImagePipeline_Component_AWS_CLI_VERSION_2_WINDOWS Ec2ImagePipeline_Component = "AWS_CLI_VERSION_2_WINDOWS"
	// Installs the latest version of the AWS CodeDeploy agent.
	//
	// This component installs only the agent. You must take additional steps to configure and use the AWS CodeDeploy agent. For more information, see the documentation at https://docs.aws.amazon.com/codedeploy/latest/userguide/welcome.html.
	// Experimental.
	Ec2ImagePipeline_Component_AWS_CODEDEPLOY_AGENT_LINUX Ec2ImagePipeline_Component = "AWS_CODEDEPLOY_AGENT_LINUX"
	// Installs the latest version of the AWS CodeDeploy agent.
	//
	// This component installs only the agent. You must take additional steps to configure and use the agent. For more information, see the documentation at https://docs.aws.amazon.com/codedeploy/latest/userguide/codedeploy-agent-operations-install-windows.html.
	// Experimental.
	Ec2ImagePipeline_Component_AWS_CODEDEPLOY_AGENT_WINDOWS Ec2ImagePipeline_Component = "AWS_CODEDEPLOY_AGENT_WINDOWS"
	// Installs the AwsVssComponents Distributor package on a Windows instance.
	//
	// The instance must have an AWS Tools for PowerShell version that includes Systems Manager modules installed. The IAM profile attached to the build instance must have the following permissions - configure the ssm:SendCommand permission with the AWS-ConfigureAWSPackage Systems Manager document on all instances in the Region, and configure the ssm:GetCommandInvocation permission for '*'. For more information, see the documentation at https://docs.aws.amazon.com/imagebuilder/latest/userguide/mgdcomponent-distributor-win.html and https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/application-consistent-snapshots.html.
	// Experimental.
	Ec2ImagePipeline_Component_AWS_VSS_COMPONENTS_WINDOWS Ec2ImagePipeline_Component = "AWS_VSS_COMPONENTS_WINDOWS"
	// Installs Chocolatey for Windows.
	// Experimental.
	Ec2ImagePipeline_Component_CHOCOLATEY Ec2ImagePipeline_Component = "CHOCOLATEY"
	// Validates the Chrony configuration file and ensures that Chrony time sources on Amazon Linux 2 are configured for the Amazon time servers.
	//
	// Uses validation steps outlined here: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/set-time.html.
	// Experimental.
	Ec2ImagePipeline_Component_CHRONY_TIME_CONFIGURATION_TEST Ec2ImagePipeline_Component = "CHRONY_TIME_CONFIGURATION_TEST"
	// Install and configure the latest NICE DCV server on Linux.
	// Experimental.
	Ec2ImagePipeline_Component_DCV_SERVER_LINUX Ec2ImagePipeline_Component = "DCV_SERVER_LINUX"
	// Install and configure the latest NICE DCV server on Windows.
	// Experimental.
	Ec2ImagePipeline_Component_DCV_SERVER_WINDOWS Ec2ImagePipeline_Component = "DCV_SERVER_WINDOWS"
	// Installs a Distributor package on a Windows instance.
	//
	// The instance must have an AWS Tools for PowerShell version that includes Systems Manager modules installed. The IAM profile attached to the build instance must have the following permissions - configure the ssm:SendCommand permission with the AWS-ConfigureAWSPackage Systems Manager document on all instances in the Region, and configure the ssm:GetCommandInvocation permission for '*'. For more information, see the documentation at https://docs.aws.amazon.com/imagebuilder/latest/userguide/mgdcomponent-distributor-win.html.
	// Experimental.
	Ec2ImagePipeline_Component_DISTRIBUTOR_PACKAGE_WINDOWS Ec2ImagePipeline_Component = "DISTRIBUTOR_PACKAGE_WINDOWS"
	// Installs Docker Community Edition from the Docker package repository, and enables the centos user to manage Docker without using sudo.
	//
	// For more information, review the installation guide at https://docs.docker.com/install/linux/docker-ce/centos/.
	// Experimental.
	Ec2ImagePipeline_Component_DOCKER_CE_CENTOS Ec2ImagePipeline_Component = "DOCKER_CE_CENTOS"
	// Install the latest Docker Community Edition from Amazon Linux Extras, and enable the ec2-user user to manage docker without using sudo.
	// Experimental.
	Ec2ImagePipeline_Component_DOCKER_CE_LINUX Ec2ImagePipeline_Component = "DOCKER_CE_LINUX"
	// Installs Docker Community Edition from the Docker package repository, and enables the ubuntu user to manage Docker without using sudo.
	//
	// For more information, review the installation guide at https://docs.docker.com/install/linux/docker-ce/ubuntu/.
	// Experimental.
	Ec2ImagePipeline_Component_DOCKER_CE_UBUNTU Ec2ImagePipeline_Component = "DOCKER_CE_UBUNTU"
	// Installs the latest 8.0 channel release of the Microsoft .NET Desktop Runtime. For more information, see the .NET 8.0 download page at https://dotnet.microsoft.com/download/dotnet/8.0.
	// Experimental.
	Ec2ImagePipeline_Component_DOTNET_DESKTOP_RUNTIME_LTS_WINDOWS Ec2ImagePipeline_Component = "DOTNET_DESKTOP_RUNTIME_LTS_WINDOWS"
	// Installs the latest 8.0 channel release of the Microsoft .NET Hosting Bundle. For more information, see the .NET 8.0 download page at https://dotnet.microsoft.com/download/dotnet/8.0.
	// Experimental.
	Ec2ImagePipeline_Component_DOTNET_HOSTING_BUNDLE_LTS_WINDOWS Ec2ImagePipeline_Component = "DOTNET_HOSTING_BUNDLE_LTS_WINDOWS"
	// Installs the latest 8.0 channel release of the Microsoft .NET Runtime. For more information, see the .NET 8.0 download page at https://dotnet.microsoft.com/download/dotnet/8.0.
	// Experimental.
	Ec2ImagePipeline_Component_DOTNET_RUNTIME_LTS_LINUX Ec2ImagePipeline_Component = "DOTNET_RUNTIME_LTS_LINUX"
	// Installs the latest 8.0 channel release of the Microsoft .NET Runtime. For more information, see the .NET 8.0 download page at https://dotnet.microsoft.com/download/dotnet/8.0.
	// Experimental.
	Ec2ImagePipeline_Component_DOTNET_RUNTIME_LTS_WINDOWS Ec2ImagePipeline_Component = "DOTNET_RUNTIME_LTS_WINDOWS"
	// Installs the latest 8.0 channel release of the Microsoft .NET SDK. For more information, see the .NET 8.0 download page at https://dotnet.microsoft.com/download/dotnet/8.0.
	// Experimental.
	Ec2ImagePipeline_Component_DOTNET_SDK_LTS_LINUX Ec2ImagePipeline_Component = "DOTNET_SDK_LTS_LINUX"
	// Installs the latest 8.0 channel release of the Microsoft .NET SDK. For more information, see the .NET 8.0 download page at https://dotnet.microsoft.com/download/dotnet/8.0.
	// Experimental.
	Ec2ImagePipeline_Component_DOTNET_SDK_LTS_WINDOWS Ec2ImagePipeline_Component = "DOTNET_SDK_LTS_WINDOWS"
	// The EBS volume usage test performs the following actions: 1) It creates an EBS volume and attaches it to the instance.
	//
	// 2) It creates a temporary file on the volume and detaches the volume. 3) It reattaches the volume and validates that the file exists. 4) It detaches and deletes the volume. To perform this test, an IAM policy with the following actions is required: ec2:AttachVolume, ec2:Create Tags, ec2:CreateVolume, ec2:DeleteVolume, ec2:DescribeVolumes, and ec2:DetachVolume.
	// Experimental.
	Ec2ImagePipeline_Component_EBS_VOLUME_USAGE_TEST_LINUX Ec2ImagePipeline_Component = "EBS_VOLUME_USAGE_TEST_LINUX"
	// The EBS volume usage test performs the following actions: 1) It creates an EBS volume and attaches it to the instance.
	//
	// 2) It creates a temporary file on the volume and detaches the volume. 3) It reattaches the volume and validates that the file exists. 4) It detaches and deletes the volume. To perform this test, an IAM policy with the following actions is required: ec2:AttachVolume, ec2:Create Tags, ec2:CreateVolume, ec2:DeleteVolume, ec2:DescribeVolumes, and ec2:DetachVolume.
	// Experimental.
	Ec2ImagePipeline_Component_EBS_VOLUME_USAGE_TEST_WINDOWS Ec2ImagePipeline_Component = "EBS_VOLUME_USAGE_TEST_WINDOWS"
	// Test to ensure all required EC2 network routes exist in the route table.
	// Experimental.
	Ec2ImagePipeline_Component_EC2_NETWORK_ROUTE_TEST_WINDOWS Ec2ImagePipeline_Component = "EC2_NETWORK_ROUTE_TEST_WINDOWS"
	// Installs the latest version of EC2Launch v2.
	//
	// For more information, see the documentation at https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2launch-v2.html.
	// Experimental.
	Ec2ImagePipeline_Component_EC2LAUNCH_V2_WINDOWS Ec2ImagePipeline_Component = "EC2LAUNCH_V2_WINDOWS"
	// Installs Amazon ECS-optimized Windows artifacts.
	//
	// This includes latest Amazon ECS Container Agent and Docker CE version 20.10.21.
	// Experimental.
	Ec2ImagePipeline_Component_ECS_OPTIMIZED_AMI_WINDOWS Ec2ImagePipeline_Component = "ECS_OPTIMIZED_AMI_WINDOWS"
	// Installs Amazon EKS-optimized Windows artifacts for Amazon EKS version 1.30. This includes kubelet version 1.30.2, containerd version 1.7.14, and CSI Proxy version 1.1.2.
	// Experimental.
	Ec2ImagePipeline_Component_EKS_OPTIMIZED_AMI_WINDOWS Ec2ImagePipeline_Component = "EKS_OPTIMIZED_AMI_WINDOWS"
	// The ENI attachment test performs the following actions: 1) It creates an elastic network interface (ENI) and attaches it to the instance.
	//
	// 2) It validates that the attached ENI has an IP address. 3) It detaches and deletes the ENI. To perform this test, an IAM policy with the following actions is required: ec2:AttachNetworkInterface, ec2:CreateNetworkInterface, ec2:CreateTags, ec2:DeleteNetworkInterface, ec2:DescribeNetworkInterfaces, ec2:DescribeNetworkInterfaceAttribute, and ec2:DetachNetworkInterface.
	// Experimental.
	Ec2ImagePipeline_Component_ENI_ATTACHMENT_TEST_LINUX Ec2ImagePipeline_Component = "ENI_ATTACHMENT_TEST_LINUX"
	// The ENI attachment test performs the following actions: 1) It creates an elastic network interface (ENI) and attaches it to the instance.
	//
	// 2) It validates that the attached ENI has an IP address. 3) It detaches and deletes the ENI. To perform this test, an IAM policy with the following actions is required: ec2:AttachNetworkInterface, ec2:CreateNetworkInterface, ec2:CreateTags, ec2:DeleteNetworkInterface, ec2:DescribeNetworkInterfaces, ec2:DescribeNetworkInterfaceAttribute, and ec2:DetachNetworkInterface.
	// Experimental.
	Ec2ImagePipeline_Component_ENI_ATTACHMENT_TEST_WINDOWS Ec2ImagePipeline_Component = "ENI_ATTACHMENT_TEST_WINDOWS"
	// Installs the latest stable release of the Go programming language using the release information from https://go.dev/dl/.
	// Experimental.
	Ec2ImagePipeline_Component_GO_STABLE_LINUX Ec2ImagePipeline_Component = "GO_STABLE_LINUX"
	// Installs the latest stable release of the Go programming language using the release information from https://go.dev/dl/.
	// Experimental.
	Ec2ImagePipeline_Component_GO_STABLE_WINDOWS Ec2ImagePipeline_Component = "GO_STABLE_WINDOWS"
	// Hello world testing document for Linux.
	// Experimental.
	Ec2ImagePipeline_Component_HELLO_WORLD_LINUX Ec2ImagePipeline_Component = "HELLO_WORLD_LINUX"
	// Hello world testing document for Windows.
	// Experimental.
	Ec2ImagePipeline_Component_HELLO_WORLD_WINDOWS Ec2ImagePipeline_Component = "HELLO_WORLD_WINDOWS"
	// Performs a Center for Internet Security (CIS) security assessment for an instance, using Amazon Inspector (Inspector).
	//
	// This component performs the following actions: 1) It installs the Inspector agent. 2) It creates a resource group, assessment target, and assessment template. 3) It runs the assessment and provides a link to the results in the logs and on the Inspector Service console. In order to run successfully, this component requires that the AmazonInspectorFullAccess IAM policy and the ssm:SendCommand and ec2:CreateTags IAM permissions are attached to the instance profile. To find the list of supported Operating Systems and their rules packages, refer to the Inspector documentation https://docs.aws.amazon.com/inspector/v1/userguide/inspector_rule-packages_across_os.html.
	// Experimental.
	Ec2ImagePipeline_Component_INSPECTOR_TEST_LINUX Ec2ImagePipeline_Component = "INSPECTOR_TEST_LINUX"
	// Performs a Center for Internet Security (CIS) security assessment for an instance, using Amazon Inspector (Inspector).
	//
	// This component performs the following actions: 1) It installs the Inspector agent. 2) It creates a resource group, assessment target, and assessment template. 3) It runs the assessment and provides a link to the results in the logs and on the Inspector Service console. In order to run successfully, this component requires that the AmazonInspectorFullAccess IAM policy and the ssm:SendCommand and ec2:CreateTags IAM permissions are attached to the instance profile. To find the list of supported Operating Systems and their rules packages, refer to the Inspector documentation https://docs.aws.amazon.com/inspector/v1/userguide/inspector_rule-packages_across_os.html.
	// Experimental.
	Ec2ImagePipeline_Component_INSPECTOR_TEST_WINDOWS Ec2ImagePipeline_Component = "INSPECTOR_TEST_WINDOWS"
	// Installs a package from the Linux repository.
	// Experimental.
	Ec2ImagePipeline_Component_INSTALL_PACKAGE_FROM_REPOSITORY Ec2ImagePipeline_Component = "INSTALL_PACKAGE_FROM_REPOSITORY"
	// Installs the MariaDB package using apt, yum, or zypper.
	// Experimental.
	Ec2ImagePipeline_Component_MARIADB_LINUX Ec2ImagePipeline_Component = "MARIADB_LINUX"
	// Installs the MATE Desktop Environment, xrdp, TigerVNC server, and enables the xrdp service.
	// Experimental.
	Ec2ImagePipeline_Component_MATE_DE_LINUX Ec2ImagePipeline_Component = "MATE_DE_LINUX"
	// Installs the latest version of the Mono framework.
	//
	// Follows the instructions found at https://www.mono-project.com/.
	// Experimental.
	Ec2ImagePipeline_Component_MONO_LINUX Ec2ImagePipeline_Component = "MONO_LINUX"
	// Installs PHP 8.2.
	// Experimental.
	Ec2ImagePipeline_Component_PHP_8_2_LINUX Ec2ImagePipeline_Component = "PHP_8_2_LINUX"
	// Installs the latest LTS 7.4 release of PowerShell following the instructions at https://docs.microsoft.com/en-us/powershell/scripting/install/installing-powershell-on-linux?view=powershell-7.4.
	// Experimental.
	Ec2ImagePipeline_Component_POWERSHELL_LTS_LINUX Ec2ImagePipeline_Component = "POWERSHELL_LTS_LINUX"
	// Installs the latest LTS 7.4 release of PowerShell using the MSI installer from the GitHub repository located at https://github.com/PowerShell/PowerShell.
	// Experimental.
	Ec2ImagePipeline_Component_POWERSHELL_LTS_WINDOWS Ec2ImagePipeline_Component = "POWERSHELL_LTS_WINDOWS"
	// Installs the latest version of PowerShell using snap.
	// Experimental.
	Ec2ImagePipeline_Component_POWERSHELL_SNAP Ec2ImagePipeline_Component = "POWERSHELL_SNAP"
	// Installs the latest version of PowerShell from the Microsoft RedHat repository.
	// Experimental.
	Ec2ImagePipeline_Component_POWERSHELL_YUM Ec2ImagePipeline_Component = "POWERSHELL_YUM"
	// Installs the latest version of PuTTY from the 64-bit MSI link on the release page: https://the.earth.li/~sgtatham/putty/latest/w64/.
	// Experimental.
	Ec2ImagePipeline_Component_PUTTY Ec2ImagePipeline_Component = "PUTTY"
	// Installs the Python 3 package using apt, yum, or zypper.
	// Experimental.
	Ec2ImagePipeline_Component_PYTHON_3_LINUX Ec2ImagePipeline_Component = "PYTHON_3_LINUX"
	// Installs Python 3.8.2 for Windows.
	// Experimental.
	Ec2ImagePipeline_Component_PYTHON_3_WINDOWS Ec2ImagePipeline_Component = "PYTHON_3_WINDOWS"
	// Reboots the system.
	// Experimental.
	Ec2ImagePipeline_Component_REBOOT_LINUX Ec2ImagePipeline_Component = "REBOOT_LINUX"
	// Tests whether the system can reboot successfully.
	// Experimental.
	Ec2ImagePipeline_Component_REBOOT_TEST_LINUX Ec2ImagePipeline_Component = "REBOOT_TEST_LINUX"
	// Tests whether the system can reboot successfully.
	// Experimental.
	Ec2ImagePipeline_Component_REBOOT_TEST_WINDOWS Ec2ImagePipeline_Component = "REBOOT_TEST_WINDOWS"
	// Reboots the system.
	// Experimental.
	Ec2ImagePipeline_Component_REBOOT_WINDOWS Ec2ImagePipeline_Component = "REBOOT_WINDOWS"
	// Installs SANS SIFT v1.14.0 on Ubuntu, allowing you to leverage a suite of forensics tools. For more information, see https://www.sans.org/tools/sift-workstation/.
	// Experimental.
	Ec2ImagePipeline_Component_SAN_SIFT_LINUX Ec2ImagePipeline_Component = "SAN_SIFT_LINUX"
	// Installs and runs SCAP Compliance Checker (SCC) 5.8 for Red Hat Enterprise Linux (RHEL) 7/8, Ubuntu 18.04/20.04 with all current STIG Q4 2023 benchmarks. SCC supports the AMD64 architecture. Other architectures are not currently supported. For more information, see https://docs.aws.amazon.com/imagebuilder/latest/userguide/toe-stig.html.
	// Experimental.
	Ec2ImagePipeline_Component_SCAP_COMPLIANCE_CHECKER_LINUX Ec2ImagePipeline_Component = "SCAP_COMPLIANCE_CHECKER_LINUX"
	// Installs and runs SCAP Compliance Checker (SCC) 5.10 for Windows with all current STIG Q3 2024 benchmarks. For more information, see https://docs.aws.amazon.com/imagebuilder/latest/userguide/image-builder-stig.html.
	// Experimental.
	Ec2ImagePipeline_Component_SCAP_COMPLIANCE_CHECKER_WINDOWS Ec2ImagePipeline_Component = "SCAP_COMPLIANCE_CHECKER_WINDOWS"
	// Executes a simple boot test.
	// Experimental.
	Ec2ImagePipeline_Component_SIMPLE_BOOT_TEST_LINUX Ec2ImagePipeline_Component = "SIMPLE_BOOT_TEST_LINUX"
	// Executes a simple boot test.
	// Experimental.
	Ec2ImagePipeline_Component_SIMPLE_BOOT_TEST_WINDOWS Ec2ImagePipeline_Component = "SIMPLE_BOOT_TEST_WINDOWS"
	// Applies the high, medium, and low severity STIG settings for Red Hat Enterprise Linux (RHEL) to Amazon Linux 2, Amazon Linux 2023, RHEL 7, CentOS Linux 7, CentOS Linux 8, CentOS Stream 9, RHEL 8, RHEL 9, Ubuntu 18.04, Ubuntu 20.04, and Ubuntu 22.04 instances. For more information, see https://docs.aws.amazon.com/imagebuilder/latest/userguide/toe-stig.html.
	// Experimental.
	Ec2ImagePipeline_Component_STIG_BUILD_LINUX_HIGH Ec2ImagePipeline_Component = "STIG_BUILD_LINUX_HIGH"
	// Applies the low severity STIG settings for Red Hat Enterprise Linux (RHEL) to Amazon Linux 2, Amazon Linux 2023, RHEL 7, CentOS Linux 7, CentOS Linux 8, CentOS Stream 9, RHEL 8, RHEL 9, Ubuntu 18.04, Ubuntu 20.04, and Ubuntu 22.04 instances. For more information, see https://docs.aws.amazon.com/imagebuilder/latest/userguide/toe-stig.html.
	// Experimental.
	Ec2ImagePipeline_Component_STIG_BUILD_LINUX_LOW Ec2ImagePipeline_Component = "STIG_BUILD_LINUX_LOW"
	// Applies the medium and low severity STIG settings for Red Hat Enterprise Linux (RHEL) to Amazon Linux 2, Amazon Linux 2023, RHEL 7, CentOS Linux 7, CentOS Linux 8, CentOS Stream 9, RHEL 8, RHEL 9, Ubuntu 18.04, Ubuntu 20.04, and Ubuntu 22.04 instances. For more information, see https://docs.aws.amazon.com/imagebuilder/latest/userguide/toe-stig.html.
	// Experimental.
	Ec2ImagePipeline_Component_STIG_BUILD_LINUX_MEDIUM Ec2ImagePipeline_Component = "STIG_BUILD_LINUX_MEDIUM"
	// Applies the high, medium, and low severity STIG settings to Windows instances.
	//
	// For more information, see https://docs.aws.amazon.com/imagebuilder/latest/userguide/image-builder-stig.html.
	// Experimental.
	Ec2ImagePipeline_Component_STIG_BUILD_WINDOWS_HIGH Ec2ImagePipeline_Component = "STIG_BUILD_WINDOWS_HIGH"
	// Applies the low severity STIG settings to Windows instances.
	//
	// For more information, see https://docs.aws.amazon.com/imagebuilder/latest/userguide/image-builder-stig.html.
	// Experimental.
	Ec2ImagePipeline_Component_STIG_BUILD_WINDOWS_LOW Ec2ImagePipeline_Component = "STIG_BUILD_WINDOWS_LOW"
	// Applies the medium and low severity STIG settings to Windows instances.
	//
	// For more information, see https://docs.aws.amazon.com/imagebuilder/latest/userguide/image-builder-stig.html.
	// Experimental.
	Ec2ImagePipeline_Component_STIG_BUILD_WINDOWS_MEDIUM Ec2ImagePipeline_Component = "STIG_BUILD_WINDOWS_MEDIUM"
	// Installs the Linux kernel 5.* for Amazon Linux 2 from Amazon Linux Extras.
	// Experimental.
	Ec2ImagePipeline_Component_UPDATE_LINUX_KERNEL_5 Ec2ImagePipeline_Component = "UPDATE_LINUX_KERNEL_5"
	// Installs the latest mainline release of the Linux kernel for CentOS 7 and Red Hat Enterprise Linux 7 and 8 via the 'kernel-ml' package from https://www.elrepo.org.
	// Experimental.
	Ec2ImagePipeline_Component_UPDATE_LINUX_KERNEL_ML Ec2ImagePipeline_Component = "UPDATE_LINUX_KERNEL_ML"
	// Updates Linux by installing all available updates via the UpdateOS action module.
	// Experimental.
	Ec2ImagePipeline_Component_UPDATE_LINUX Ec2ImagePipeline_Component = "UPDATE_LINUX"
	// Updates Windows with the latest security updates.
	// Experimental.
	Ec2ImagePipeline_Component_UPDATE_WINDOWS Ec2ImagePipeline_Component = "UPDATE_WINDOWS"
	// Ensures the `authorized_keys` file contains only the SSH public key returned from the EC2 Instance Metadata Service.
	// Experimental.
	Ec2ImagePipeline_Component_VALIDATE_SINGLE_SSH_PUBLIC_KEY_TEST_LINUX Ec2ImagePipeline_Component = "VALIDATE_SINGLE_SSH_PUBLIC_KEY_TEST_LINUX"
	// Verifies whether the SSH host key was generated after the latest boot.
	// Experimental.
	Ec2ImagePipeline_Component_VALIDATE_SSH_HOST_KEY_GENERATION_LINUX Ec2ImagePipeline_Component = "VALIDATE_SSH_HOST_KEY_GENERATION_LINUX"
	// Ensures the `authorized_keys` file contains the SSH public key returned from the EC2 Instance Metadata Service.
	// Experimental.
	Ec2ImagePipeline_Component_VALIDATE_SSH_PUBLIC_KEY_LINUX Ec2ImagePipeline_Component = "VALIDATE_SSH_PUBLIC_KEY_LINUX"
	// Verifies the Windows license status in the Common Information Model.
	// Experimental.
	Ec2ImagePipeline_Component_WINDOWS_ACTIVATION_TEST Ec2ImagePipeline_Component = "WINDOWS_ACTIVATION_TEST"
	// Checks the EC2 logs for the statement `Windows is Ready to use` and for the password generation message on Windows Server 2016 and later SKUs.
	//
	// This component does not support instances launched without an EC2 key pair.
	// Experimental.
	Ec2ImagePipeline_Component_WINDOWS_IS_READY_WITH_PASSWORD_GENERATION_TEST Ec2ImagePipeline_Component = "WINDOWS_IS_READY_WITH_PASSWORD_GENERATION_TEST"
	// Installs the Internet Information Services (IIS) web server and management tools.
	//
	// The installation is performed by enabling the Windows features built into the Windows operating system.
	// Experimental.
	Ec2ImagePipeline_Component_WINDOWS_SERVER_IIS Ec2ImagePipeline_Component = "WINDOWS_SERVER_IIS"
	// Tests whether yum repository works successfully.
	// Experimental.
	Ec2ImagePipeline_Component_YUM_REPOSITORY_TEST_LINUX Ec2ImagePipeline_Component = "YUM_REPOSITORY_TEST_LINUX"
)

