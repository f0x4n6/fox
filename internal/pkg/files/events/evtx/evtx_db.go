package evtx

var db = map[string]struct {
	msg string
	src string
}{
	"1100": {
		msg: "The event logging service has shut down",
		src: "Windows",
	},
	"1101": {
		msg: "Audit events have been dropped by the transport.",
		src: "Windows",
	},
	"1102": {
		msg: "The audit log was cleared",
		src: "Windows",
	},
	"1104": {
		msg: "The security Log is now full",
		src: "Windows",
	},
	"1105": {
		msg: "Event log automatic backup",
		src: "Windows",
	},
	"1108": {
		msg: "The event logging service encountered an error",
		src: "Windows",
	},
	"4608": {
		msg: "Windows is starting up",
		src: "Windows",
	},
	"4609": {
		msg: "Windows is shutting down",
		src: "Windows",
	},
	"4610": {
		msg: "An authentication package has been loaded by the Local Security Authority",
		src: "Windows",
	},
	"4611": {
		msg: "A trusted logon process has been registered with the Local Security Authority",
		src: "Windows",
	},
	"4612": {
		msg: "Internal resources allocated for the queuing of audit messages have been exhausted, leading to the loss of some audits.",
		src: "Windows",
	},
	"4614": {
		msg: "A notification package has been loaded by the Security Account Manager.",
		src: "Windows",
	},
	"4615": {
		msg: "Invalid use of LPC port",
		src: "Windows",
	},
	"4616": {
		msg: "The system time was changed.",
		src: "Windows",
	},
	"4618": {
		msg: "A monitored security event pattern has occurred",
		src: "Windows",
	},
	"4621": {
		msg: "Administrator recovered system from CrashOnAuditFail",
		src: "Windows",
	},
	"4622": {
		msg: "A security package has been loaded by the Local Security Authority.",
		src: "Windows",
	},
	"4624": {
		msg: "An account was successfully logged on",
		src: "Windows",
	},
	"4625": {
		msg: "An account failed to log on",
		src: "Windows",
	},
	"4626": {
		msg: "User/Device claims information",
		src: "Windows",
	},
	"4627": {
		msg: "Group membership information.",
		src: "Windows",
	},
	"4634": {
		msg: "An account was logged off",
		src: "Windows",
	},
	"4646": {
		msg: "IKE DoS-prevention mode started",
		src: "Windows",
	},
	"4647": {
		msg: "User initiated logoff",
		src: "Windows",
	},
	"4648": {
		msg: "A logon was attempted using explicit credentials",
		src: "Windows",
	},
	"4649": {
		msg: "A replay attack was detected",
		src: "Windows",
	},
	"4650": {
		msg: "An IPsec Main Mode security association was established",
		src: "Windows",
	},
	"4651": {
		msg: "An IPsec Main Mode security association was established",
		src: "Windows",
	},
	"4652": {
		msg: "An IPsec Main Mode negotiation failed",
		src: "Windows",
	},
	"4653": {
		msg: "An IPsec Main Mode negotiation failed",
		src: "Windows",
	},
	"4654": {
		msg: "An IPsec Quick Mode negotiation failed",
		src: "Windows",
	},
	"4655": {
		msg: "An IPsec Main Mode security association ended",
		src: "Windows",
	},
	"4656": {
		msg: "A handle to an object was requested",
		src: "Windows",
	},
	"4657": {
		msg: "A registry value was modified",
		src: "Windows",
	},
	"4658": {
		msg: "The handle to an object was closed",
		src: "Windows",
	},
	"4659": {
		msg: "A handle to an object was requested with intent to delete",
		src: "Windows",
	},
	"4660": {
		msg: "An object was deleted",
		src: "Windows",
	},
	"4661": {
		msg: "A handle to an object was requested",
		src: "Windows",
	},
	"4662": {
		msg: "An operation was performed on an object",
		src: "Windows",
	},
	"4663": {
		msg: "An attempt was made to access an object",
		src: "Windows",
	},
	"4664": {
		msg: "An attempt was made to create a hard link",
		src: "Windows",
	},
	"4665": {
		msg: "An attempt was made to create an application client context.",
		src: "Windows",
	},
	"4666": {
		msg: "An application attempted an operation",
		src: "Windows",
	},
	"4667": {
		msg: "An application client context was deleted",
		src: "Windows",
	},
	"4668": {
		msg: "An application was initialized",
		src: "Windows",
	},
	"4670": {
		msg: "Permissions on an object were changed",
		src: "Windows",
	},
	"4671": {
		msg: "An application attempted to access a blocked ordinal through the TBS",
		src: "Windows",
	},
	"4672": {
		msg: "Special privileges assigned to new logon",
		src: "Windows",
	},
	"4673": {
		msg: "A privileged service was called",
		src: "Windows",
	},
	"4674": {
		msg: "An operation was attempted on a privileged object",
		src: "Windows",
	},
	"4675": {
		msg: "SIDs were filtered",
		src: "Windows",
	},
	"4688": {
		msg: "A new process has been created",
		src: "Windows",
	},
	"4689": {
		msg: "A process has exited",
		src: "Windows",
	},
	"4690": {
		msg: "An attempt was made to duplicate a handle to an object",
		src: "Windows",
	},
	"4691": {
		msg: "Indirect access to an object was requested",
		src: "Windows",
	},
	"4692": {
		msg: "Backup of data protection master key was attempted",
		src: "Windows",
	},
	"4693": {
		msg: "Recovery of data protection master key was attempted",
		src: "Windows",
	},
	"4694": {
		msg: "Protection of auditable protected data was attempted",
		src: "Windows",
	},
	"4695": {
		msg: "Unprotection of auditable protected data was attempted",
		src: "Windows",
	},
	"4696": {
		msg: "A primary token was assigned to process",
		src: "Windows",
	},
	"4697": {
		msg: "A service was installed in the system",
		src: "Windows",
	},
	"4698": {
		msg: "A scheduled task was created",
		src: "Windows",
	},
	"4699": {
		msg: "A scheduled task was deleted",
		src: "Windows",
	},
	"4700": {
		msg: "A scheduled task was enabled",
		src: "Windows",
	},
	"4701": {
		msg: "A scheduled task was disabled",
		src: "Windows",
	},
	"4702": {
		msg: "A scheduled task was updated",
		src: "Windows",
	},
	"4703": {
		msg: "A token right was adjusted",
		src: "Windows",
	},
	"4704": {
		msg: "A user right was assigned",
		src: "Windows",
	},
	"4705": {
		msg: "A user right was removed",
		src: "Windows",
	},
	"4706": {
		msg: "A new trust was created to a domain",
		src: "Windows",
	},
	"4707": {
		msg: "A trust to a domain was removed",
		src: "Windows",
	},
	"4709": {
		msg: "IPsec Services was started",
		src: "Windows",
	},
	"4710": {
		msg: "IPsec Services was disabled",
		src: "Windows",
	},
	"4711": {
		msg: "PAStore Engine (1%)",
		src: "Windows",
	},
	"4712": {
		msg: "IPsec Services encountered a potentially serious failure",
		src: "Windows",
	},
	"4713": {
		msg: "Kerberos policy was changed",
		src: "Windows",
	},
	"4714": {
		msg: "Encrypted data recovery policy was changed",
		src: "Windows",
	},
	"4715": {
		msg: "The audit policy (SACL) on an object was changed",
		src: "Windows",
	},
	"4716": {
		msg: "Trusted domain information was modified",
		src: "Windows",
	},
	"4717": {
		msg: "System security access was granted to an account",
		src: "Windows",
	},
	"4718": {
		msg: "System security access was removed from an account",
		src: "Windows",
	},
	"4719": {
		msg: "System audit policy was changed",
		src: "Windows",
	},
	"4720": {
		msg: "A user account was created",
		src: "Windows",
	},
	"4722": {
		msg: "A user account was enabled",
		src: "Windows",
	},
	"4723": {
		msg: "An attempt was made to change an account's password",
		src: "Windows",
	},
	"4724": {
		msg: "An attempt was made to reset an accounts password",
		src: "Windows",
	},
	"4725": {
		msg: "A user account was disabled",
		src: "Windows",
	},
	"4726": {
		msg: "A user account was deleted",
		src: "Windows",
	},
	"4727": {
		msg: "A security-enabled global group was created",
		src: "Windows",
	},
	"4728": {
		msg: "A member was added to a security-enabled global group",
		src: "Windows",
	},
	"4729": {
		msg: "A member was removed from a security-enabled global group",
		src: "Windows",
	},
	"4730": {
		msg: "A security-enabled global group was deleted",
		src: "Windows",
	},
	"4731": {
		msg: "A security-enabled local group was created",
		src: "Windows",
	},
	"4732": {
		msg: "A member was added to a security-enabled local group",
		src: "Windows",
	},
	"4733": {
		msg: "A member was removed from a security-enabled local group",
		src: "Windows",
	},
	"4734": {
		msg: "A security-enabled local group was deleted",
		src: "Windows",
	},
	"4735": {
		msg: "A security-enabled local group was changed",
		src: "Windows",
	},
	"4737": {
		msg: "A security-enabled global group was changed",
		src: "Windows",
	},
	"4738": {
		msg: "A user account was changed",
		src: "Windows",
	},
	"4739": {
		msg: "Domain Policy was changed",
		src: "Windows",
	},
	"4740": {
		msg: "A user account was locked out",
		src: "Windows",
	},
	"4741": {
		msg: "A computer account was created",
		src: "Windows",
	},
	"4742": {
		msg: "A computer account was changed",
		src: "Windows",
	},
	"4743": {
		msg: "A computer account was deleted",
		src: "Windows",
	},
	"4744": {
		msg: "A security-disabled local group was created",
		src: "Windows",
	},
	"4745": {
		msg: "A security-disabled local group was changed",
		src: "Windows",
	},
	"4746": {
		msg: "A member was added to a security-disabled local group",
		src: "Windows",
	},
	"4747": {
		msg: "A member was removed from a security-disabled local group",
		src: "Windows",
	},
	"4748": {
		msg: "A security-disabled local group was deleted",
		src: "Windows",
	},
	"4749": {
		msg: "A security-disabled global group was created",
		src: "Windows",
	},
	"4750": {
		msg: "A security-disabled global group was changed",
		src: "Windows",
	},
	"4751": {
		msg: "A member was added to a security-disabled global group",
		src: "Windows",
	},
	"4752": {
		msg: "A member was removed from a security-disabled global group",
		src: "Windows",
	},
	"4753": {
		msg: "A security-disabled global group was deleted",
		src: "Windows",
	},
	"4754": {
		msg: "A security-enabled universal group was created",
		src: "Windows",
	},
	"4755": {
		msg: "A security-enabled universal group was changed",
		src: "Windows",
	},
	"4756": {
		msg: "A member was added to a security-enabled universal group",
		src: "Windows",
	},
	"4757": {
		msg: "A member was removed from a security-enabled universal group",
		src: "Windows",
	},
	"4758": {
		msg: "A security-enabled universal group was deleted",
		src: "Windows",
	},
	"4759": {
		msg: "A security-disabled universal group was created",
		src: "Windows",
	},
	"4760": {
		msg: "A security-disabled universal group was changed",
		src: "Windows",
	},
	"4761": {
		msg: "A member was added to a security-disabled universal group",
		src: "Windows",
	},
	"4762": {
		msg: "A member was removed from a security-disabled universal group",
		src: "Windows",
	},
	"4763": {
		msg: "A security-disabled universal group was deleted",
		src: "Windows",
	},
	"4764": {
		msg: "A groups type was changed",
		src: "Windows",
	},
	"4765": {
		msg: "SID History was added to an account",
		src: "Windows",
	},
	"4766": {
		msg: "An attempt to add SID History to an account failed",
		src: "Windows",
	},
	"4767": {
		msg: "A user account was unlocked",
		src: "Windows",
	},
	"4768": {
		msg: "A Kerberos authentication ticket (TGT) was requested",
		src: "Windows",
	},
	"4769": {
		msg: "A Kerberos service ticket was requested",
		src: "Windows",
	},
	"4770": {
		msg: "A Kerberos service ticket was renewed",
		src: "Windows",
	},
	"4771": {
		msg: "Kerberos pre-authentication failed",
		src: "Windows",
	},
	"4772": {
		msg: "A Kerberos authentication ticket request failed",
		src: "Windows",
	},
	"4773": {
		msg: "A Kerberos service ticket request failed",
		src: "Windows",
	},
	"4774": {
		msg: "An account was mapped for logon",
		src: "Windows",
	},
	"4775": {
		msg: "An account could not be mapped for logon",
		src: "Windows",
	},
	"4776": {
		msg: "The domain controller attempted to validate the credentials for an account",
		src: "Windows",
	},
	"4777": {
		msg: "The domain controller failed to validate the credentials for an account",
		src: "Windows",
	},
	"4778": {
		msg: "A session was reconnected to a Window Station",
		src: "Windows",
	},
	"4779": {
		msg: "A session was disconnected from a Window Station",
		src: "Windows",
	},
	"4780": {
		msg: "The ACL was set on accounts which are members of administrators groups",
		src: "Windows",
	},
	"4781": {
		msg: "The Data of an account was changed",
		src: "Windows",
	},
	"4782": {
		msg: "The password hash an account was accessed",
		src: "Windows",
	},
	"4783": {
		msg: "A basic application group was created",
		src: "Windows",
	},
	"4784": {
		msg: "A basic application group was changed",
		src: "Windows",
	},
	"4785": {
		msg: "A member was added to a basic application group",
		src: "Windows",
	},
	"4786": {
		msg: "A member was removed from a basic application group",
		src: "Windows",
	},
	"4787": {
		msg: "A non-member was added to a basic application group",
		src: "Windows",
	},
	"4788": {
		msg: "A non-member was removed from a basic application group..",
		src: "Windows",
	},
	"4789": {
		msg: "A basic application group was deleted",
		src: "Windows",
	},
	"4790": {
		msg: "An LDAP query group was created",
		src: "Windows",
	},
	"4791": {
		msg: "A basic application group was changed",
		src: "Windows",
	},
	"4792": {
		msg: "An LDAP query group was deleted",
		src: "Windows",
	},
	"4793": {
		msg: "The Password Policy Checking API was called",
		src: "Windows",
	},
	"4794": {
		msg: "An attempt was made to set the Directory Services Restore Mode administrator password",
		src: "Windows",
	},
	"4797": {
		msg: "An attempt was made to query the existence of a blank password for an account",
		src: "Windows",
	},
	"4798": {
		msg: "A user's local group membership was enumerated.",
		src: "Windows",
	},
	"4799": {
		msg: "A security-enabled local group membership was enumerated",
		src: "Windows",
	},
	"4800": {
		msg: "The workstation was locked",
		src: "Windows",
	},
	"4801": {
		msg: "The workstation was unlocked",
		src: "Windows",
	},
	"4802": {
		msg: "The screen saver was invoked",
		src: "Windows",
	},
	"4803": {
		msg: "The screen saver was dismissed",
		src: "Windows",
	},
	"4816": {
		msg: "RPC detected an integrity violation while decrypting an incoming message",
		src: "Windows",
	},
	"4817": {
		msg: "Auditing settings on object were changed.",
		src: "Windows",
	},
	"4818": {
		msg: "Proposed Central Access Policy does not grant the same access permissions as the current Central Access Policy",
		src: "Windows",
	},
	"4819": {
		msg: "Central Access Policies on the machine have been changed",
		src: "Windows",
	},
	"4820": {
		msg: "A Kerberos Ticket-granting-ticket (TGT) was denied because the device does not meet the access control restrictions",
		src: "Windows",
	},
	"4821": {
		msg: "A Kerberos service ticket was denied because the user, device, or both does not meet the access control restrictions",
		src: "Windows",
	},
	"4822": {
		msg: "NTLM authentication failed because the account was a member of the Protected User group",
		src: "Windows",
	},
	"4823": {
		msg: "NTLM authentication failed because access control restrictions are required",
		src: "Windows",
	},
	"4824": {
		msg: "Kerberos preauthentication by using DES or RC4 failed because the account was a member of the Protected User group",
		src: "Windows",
	},
	"4825": {
		msg: "A user was denied the access to Remote Desktop. By default, users are allowed to connect only if they are members of the Remote Desktop Users group or Administrators group",
		src: "Windows",
	},
	"4826": {
		msg: "Boot Configuration Data loaded",
		src: "Windows",
	},
	"4830": {
		msg: "SID History was removed from an account",
		src: "Windows",
	},
	"4864": {
		msg: "A Dataspace collision was detected",
		src: "Windows",
	},
	"4865": {
		msg: "A trusted forest information entry was added",
		src: "Windows",
	},
	"4866": {
		msg: "A trusted forest information entry was removed",
		src: "Windows",
	},
	"4867": {
		msg: "A trusted forest information entry was modified",
		src: "Windows",
	},
	"4868": {
		msg: "The certificate manager denied a pending certificate request",
		src: "Windows",
	},
	"4869": {
		msg: "Certificate Services received a resubmitted certificate request",
		src: "Windows",
	},
	"4870": {
		msg: "Certificate Services revoked a certificate",
		src: "Windows",
	},
	"4871": {
		msg: "Certificate Services received a request to publish the certificate revocation list (CRL)",
		src: "Windows",
	},
	"4872": {
		msg: "Certificate Services published the certificate revocation list (CRL)",
		src: "Windows",
	},
	"4873": {
		msg: "A certificate request extension changed",
		src: "Windows",
	},
	"4874": {
		msg: "One or more certificate request attributes changed.",
		src: "Windows",
	},
	"4875": {
		msg: "Certificate Services received a request to shut down",
		src: "Windows",
	},
	"4876": {
		msg: "Certificate Services backup started",
		src: "Windows",
	},
	"4877": {
		msg: "Certificate Services backup completed",
		src: "Windows",
	},
	"4878": {
		msg: "Certificate Services restore started",
		src: "Windows",
	},
	"4879": {
		msg: "Certificate Services restore completed",
		src: "Windows",
	},
	"4880": {
		msg: "Certificate Services started",
		src: "Windows",
	},
	"4881": {
		msg: "Certificate Services stopped",
		src: "Windows",
	},
	"4882": {
		msg: "The security permissions for Certificate Services changed",
		src: "Windows",
	},
	"4883": {
		msg: "Certificate Services retrieved an archived key",
		src: "Windows",
	},
	"4884": {
		msg: "Certificate Services imported a certificate into its database",
		src: "Windows",
	},
	"4885": {
		msg: "The audit filter for Certificate Services changed",
		src: "Windows",
	},
	"4886": {
		msg: "Certificate Services received a certificate request",
		src: "Windows",
	},
	"4887": {
		msg: "Certificate Services approved a certificate request and issued a certificate",
		src: "Windows",
	},
	"4888": {
		msg: "Certificate Services denied a certificate request",
		src: "Windows",
	},
	"4889": {
		msg: "Certificate Services set the status of a certificate request to pending",
		src: "Windows",
	},
	"4890": {
		msg: "The certificate manager settings for Certificate Services changed.",
		src: "Windows",
	},
	"4891": {
		msg: "A configuration entry changed in Certificate Services",
		src: "Windows",
	},
	"4892": {
		msg: "A property of Certificate Services changed",
		src: "Windows",
	},
	"4893": {
		msg: "Certificate Services archived a key",
		src: "Windows",
	},
	"4894": {
		msg: "Certificate Services imported and archived a key",
		src: "Windows",
	},
	"4895": {
		msg: "Certificate Services published the CA certificate to Active Directory Domain Services",
		src: "Windows",
	},
	"4896": {
		msg: "One or more rows have been deleted from the certificate database",
		src: "Windows",
	},
	"4897": {
		msg: "Role separation enabled",
		src: "Windows",
	},
	"4898": {
		msg: "Certificate Services loaded a template",
		src: "Windows",
	},
	"4899": {
		msg: "A Certificate Services template was updated",
		src: "Windows",
	},
	"4900": {
		msg: "Certificate Services template security was updated",
		src: "Windows",
	},
	"4902": {
		msg: "The Per-user audit policy table was created",
		src: "Windows",
	},
	"4904": {
		msg: "An attempt was made to register a security event Orig",
		src: "Windows",
	},
	"4905": {
		msg: "An attempt was made to unregister a security event Orig",
		src: "Windows",
	},
	"4906": {
		msg: "The CrashOnAuditFail value has changed",
		src: "Windows",
	},
	"4907": {
		msg: "Auditing settings on object were changed",
		src: "Windows",
	},
	"4908": {
		msg: "Special Groups Logon table modified",
		src: "Windows",
	},
	"4909": {
		msg: "The local policy settings for the TBS were changed",
		src: "Windows",
	},
	"4910": {
		msg: "The group policy settings for the TBS were changed",
		src: "Windows",
	},
	"4911": {
		msg: "ReOrig attributes of the object were changed",
		src: "Windows",
	},
	"4912": {
		msg: "Per User Audit Policy was changed",
		src: "Windows",
	},
	"4913": {
		msg: "Central Access Policy on the object was changed",
		src: "Windows",
	},
	"4928": {
		msg: "An Active Directory replica Orig naming context was established",
		src: "Windows",
	},
	"4929": {
		msg: "An Active Directory replica Orig naming context was removed",
		src: "Windows",
	},
	"4930": {
		msg: "An Active Directory replica Orig naming context was modified",
		src: "Windows",
	},
	"4931": {
		msg: "An Active Directory replica destination naming context was modified",
		src: "Windows",
	},
	"4932": {
		msg: "Synchronization of a replica of an Active Directory naming context has begun",
		src: "Windows",
	},
	"4933": {
		msg: "Synchronization of a replica of an Active Directory naming context has ended",
		src: "Windows",
	},
	"4934": {
		msg: "Attributes of an Active Directory object were replicated",
		src: "Windows",
	},
	"4935": {
		msg: "Replication failure begins",
		src: "Windows",
	},
	"4936": {
		msg: "Replication failure ends",
		src: "Windows",
	},
	"4937": {
		msg: "A lingering object was removed from a replica",
		src: "Windows",
	},
	"4944": {
		msg: "The following policy was active when the Windows Firewall started",
		src: "Windows",
	},
	"4945": {
		msg: "A rule was listed when the Windows Firewall started",
		src: "Windows",
	},
	"4946": {
		msg: "A change has been made to Windows Firewall exception list. A rule was added",
		src: "Windows",
	},
	"4947": {
		msg: "A change has been made to Windows Firewall exception list. A rule was modified",
		src: "Windows",
	},
	"4948": {
		msg: "A change has been made to Windows Firewall exception list. A rule was deleted",
		src: "Windows",
	},
	"4949": {
		msg: "Windows Firewall settings were restored to the default values",
		src: "Windows",
	},
	"4950": {
		msg: "A Windows Firewall setting has changed",
		src: "Windows",
	},
	"4951": {
		msg: "A rule has been ignored because its major version number was not recognized by Windows Firewall",
		src: "Windows",
	},
	"4952": {
		msg: "Parts of a rule have been ignored because its minor version number was not recognized by Windows Firewall",
		src: "Windows",
	},
	"4953": {
		msg: "A rule has been ignored by Windows Firewall because it could not parse the rule",
		src: "Windows",
	},
	"4954": {
		msg: "Windows Firewall Group Policy settings has changed. The new settings have been applied",
		src: "Windows",
	},
	"4956": {
		msg: "Windows Firewall has changed the active profile",
		src: "Windows",
	},
	"4957": {
		msg: "Windows Firewall did not apply the following rule",
		src: "Windows",
	},
	"4958": {
		msg: "Windows Firewall did not apply the following rule because the rule referred to items not configured on this computer",
		src: "Windows",
	},
	"4960": {
		msg: "IPsec dropped an inbound packet that failed an integrity check",
		src: "Windows",
	},
	"4961": {
		msg: "IPsec dropped an inbound packet that failed a replay check",
		src: "Windows",
	},
	"4962": {
		msg: "IPsec dropped an inbound packet that failed a replay check",
		src: "Windows",
	},
	"4963": {
		msg: "IPsec dropped an inbound clear text packet that should have been secured",
		src: "Windows",
	},
	"4964": {
		msg: "Special groups have been assigned to a new logon",
		src: "Windows",
	},
	"4965": {
		msg: "IPsec received a packet from a remote computer with an incorrect Security Parameter Index (SPI).",
		src: "Windows",
	},
	"4976": {
		msg: "During Main Mode negotiation, IPsec received an invalid negotiation packet.",
		src: "Windows",
	},
	"4977": {
		msg: "During Quick Mode negotiation, IPsec received an invalid negotiation packet.",
		src: "Windows",
	},
	"4978": {
		msg: "During Extended Mode negotiation, IPsec received an invalid negotiation packet.",
		src: "Windows",
	},
	"4979": {
		msg: "IPsec Main Mode and Extended Mode security associations were established.",
		src: "Windows",
	},
	"4980": {
		msg: "IPsec Main Mode and Extended Mode security associations were established",
		src: "Windows",
	},
	"4981": {
		msg: "IPsec Main Mode and Extended Mode security associations were established",
		src: "Windows",
	},
	"4982": {
		msg: "IPsec Main Mode and Extended Mode security associations were established",
		src: "Windows",
	},
	"4983": {
		msg: "An IPsec Extended Mode negotiation failed",
		src: "Windows",
	},
	"4984": {
		msg: "An IPsec Extended Mode negotiation failed",
		src: "Windows",
	},
	"4985": {
		msg: "The state of a transaction has changed",
		src: "Windows",
	},
	"5024": {
		msg: "The Windows Firewall Service has started successfully",
		src: "Windows",
	},
	"5025": {
		msg: "The Windows Firewall Service has been stopped",
		src: "Windows",
	},
	"5027": {
		msg: "The Windows Firewall Service was unable to retrieve the security policy from the local storage",
		src: "Windows",
	},
	"5028": {
		msg: "The Windows Firewall Service was unable to parse the new security policy.",
		src: "Windows",
	},
	"5029": {
		msg: "The Windows Firewall Service failed to initialize the driver",
		src: "Windows",
	},
	"5030": {
		msg: "The Windows Firewall Service failed to start",
		src: "Windows",
	},
	"5031": {
		msg: "The Windows Firewall Service blocked an application from accepting incoming connections on the network.",
		src: "Windows",
	},
	"5032": {
		msg: "Windows Firewall was unable to notify the user that it blocked an application from accepting incoming connections on the network",
		src: "Windows",
	},
	"5033": {
		msg: "The Windows Firewall Driver has started successfully",
		src: "Windows",
	},
	"5034": {
		msg: "The Windows Firewall Driver has been stopped",
		src: "Windows",
	},
	"5035": {
		msg: "The Windows Firewall Driver failed to start",
		src: "Windows",
	},
	"5037": {
		msg: "The Windows Firewall Driver detected critical runtime error. Terminating",
		src: "Windows",
	},
	"5038": {
		msg: "Code integrity determined that the image hash of a file is not valid",
		src: "Windows",
	},
	"5039": {
		msg: "A registry key was virtualized.",
		src: "Windows",
	},
	"5040": {
		msg: "A change has been made to IPsec settings. An Authentication Set was added.",
		src: "Windows",
	},
	"5041": {
		msg: "A change has been made to IPsec settings. An Authentication Set was modified",
		src: "Windows",
	},
	"5042": {
		msg: "A change has been made to IPsec settings. An Authentication Set was deleted",
		src: "Windows",
	},
	"5043": {
		msg: "A change has been made to IPsec settings. A Connection Security Rule was added",
		src: "Windows",
	},
	"5044": {
		msg: "A change has been made to IPsec settings. A Connection Security Rule was modified",
		src: "Windows",
	},
	"5045": {
		msg: "A change has been made to IPsec settings. A Connection Security Rule was deleted",
		src: "Windows",
	},
	"5046": {
		msg: "A change has been made to IPsec settings. A Crypto Set was added",
		src: "Windows",
	},
	"5047": {
		msg: "A change has been made to IPsec settings. A Crypto Set was modified",
		src: "Windows",
	},
	"5048": {
		msg: "A change has been made to IPsec settings. A Crypto Set was deleted",
		src: "Windows",
	},
	"5049": {
		msg: "An IPsec Security Association was deleted",
		src: "Windows",
	},
	"5050": {
		msg: "An attempt to programmatically disable the Windows Firewall using a call to INetFwProfile.FirewallEnabled(FALSE",
		src: "Windows",
	},
	"5051": {
		msg: "A file was virtualized",
		src: "Windows",
	},
	"5056": {
		msg: "A cryptographic self test was performed",
		src: "Windows",
	},
	"5057": {
		msg: "A cryptographic primitive operation failed",
		src: "Windows",
	},
	"5058": {
		msg: "Key file operation",
		src: "Windows",
	},
	"5059": {
		msg: "Key migration operation",
		src: "Windows",
	},
	"5060": {
		msg: "Verification operation failed",
		src: "Windows",
	},
	"5061": {
		msg: "Cryptographic operation",
		src: "Windows",
	},
	"5062": {
		msg: "A kernel-mode cryptographic self test was performed",
		src: "Windows",
	},
	"5063": {
		msg: "A cryptographic provider operation was attempted",
		src: "Windows",
	},
	"5064": {
		msg: "A cryptographic context operation was attempted",
		src: "Windows",
	},
	"5065": {
		msg: "A cryptographic context modification was attempted",
		src: "Windows",
	},
	"5066": {
		msg: "A cryptographic function operation was attempted",
		src: "Windows",
	},
	"5067": {
		msg: "A cryptographic function modification was attempted",
		src: "Windows",
	},
	"5068": {
		msg: "A cryptographic function provider operation was attempted",
		src: "Windows",
	},
	"5069": {
		msg: "A cryptographic function property operation was attempted",
		src: "Windows",
	},
	"5070": {
		msg: "A cryptographic function property operation was attempted",
		src: "Windows",
	},
	"5071": {
		msg: "Key access denied by Microsoft key distribution service",
		src: "Windows",
	},
	"5120": {
		msg: "OCSP Responder Service Started",
		src: "Windows",
	},
	"5121": {
		msg: "OCSP Responder Service Stopped",
		src: "Windows",
	},
	"5122": {
		msg: "A Configuration entry changed in the OCSP Responder Service",
		src: "Windows",
	},
	"5123": {
		msg: "A configuration entry changed in the OCSP Responder Service",
		src: "Windows",
	},
	"5124": {
		msg: "A security setting was updated on OCSP Responder Service",
		src: "Windows",
	},
	"5125": {
		msg: "A request was submitted to OCSP Responder Service",
		src: "Windows",
	},
	"5126": {
		msg: "Signing Certificate was automatically updated by the OCSP Responder Service",
		src: "Windows",
	},
	"5127": {
		msg: "The OCSP Revocation Provider successfully updated the revocation information",
		src: "Windows",
	},
	"5136": {
		msg: "A directory service object was modified",
		src: "Windows",
	},
	"5137": {
		msg: "A directory service object was created",
		src: "Windows",
	},
	"5138": {
		msg: "A directory service object was undeleted",
		src: "Windows",
	},
	"5139": {
		msg: "A directory service object was moved",
		src: "Windows",
	},
	"5140": {
		msg: "A network share object was accessed",
		src: "Windows",
	},
	"5141": {
		msg: "A directory service object was deleted",
		src: "Windows",
	},
	"5142": {
		msg: "A network share object was added.",
		src: "Windows",
	},
	"5143": {
		msg: "A network share object was modified",
		src: "Windows",
	},
	"5144": {
		msg: "A network share object was deleted.",
		src: "Windows",
	},
	"5145": {
		msg: "A network share object was checked to see whether client can be granted desired access",
		src: "Windows",
	},
	"5146": {
		msg: "The Windows Filtering Platform has blocked a packet",
		src: "Windows",
	},
	"5147": {
		msg: "A more restrictive Windows Filtering Platform filter has blocked a packet",
		src: "Windows",
	},
	"5148": {
		msg: "The Windows Filtering Platform has detected a DoS attack and entered a defensive mode; packets associated with this attack will be discarded.",
		src: "Windows",
	},
	"5149": {
		msg: "The DoS attack has subsided and normal processing is being resumed.",
		src: "Windows",
	},
	"5150": {
		msg: "The Windows Filtering Platform has blocked a packet.",
		src: "Windows",
	},
	"5151": {
		msg: "A more restrictive Windows Filtering Platform filter has blocked a packet.",
		src: "Windows",
	},
	"5152": {
		msg: "The Windows Filtering Platform blocked a packet",
		src: "Windows",
	},
	"5153": {
		msg: "A more restrictive Windows Filtering Platform filter has blocked a packet",
		src: "Windows",
	},
	"5154": {
		msg: "The Windows Filtering Platform has permitted an application or service to listen on a port for incoming connections",
		src: "Windows",
	},
	"5155": {
		msg: "The Windows Filtering Platform has blocked an application or service from listening on a port for incoming connections",
		src: "Windows",
	},
	"5156": {
		msg: "The Windows Filtering Platform has allowed a connection",
		src: "Windows",
	},
	"5157": {
		msg: "The Windows Filtering Platform has blocked a connection",
		src: "Windows",
	},
	"5158": {
		msg: "The Windows Filtering Platform has permitted a bind to a local port",
		src: "Windows",
	},
	"5159": {
		msg: "The Windows Filtering Platform has blocked a bind to a local port",
		src: "Windows",
	},
	"5168": {
		msg: "Spn check for SMB/SMB2 fails.",
		src: "Windows",
	},
	"5169": {
		msg: "A directory service object was modified",
		src: "Windows",
	},
	"5170": {
		msg: "A directory service object was modified during a background cleanup task",
		src: "Windows",
	},
	"5376": {
		msg: "Credential Manager credentials were backed up",
		src: "Windows",
	},
	"5377": {
		msg: "Credential Manager credentials were restored from a backup",
		src: "Windows",
	},
	"5378": {
		msg: "The requested credentials delegation was disallowed by policy",
		src: "Windows",
	},
	"5440": {
		msg: "The following callout was present when the Windows Filtering Platform Base Filtering Engine started",
		src: "Windows",
	},
	"5441": {
		msg: "The following filter was present when the Windows Filtering Platform Base Filtering Engine started",
		src: "Windows",
	},
	"5442": {
		msg: "The following provider was present when the Windows Filtering Platform Base Filtering Engine started",
		src: "Windows",
	},
	"5443": {
		msg: "The following provider context was present when the Windows Filtering Platform Base Filtering Engine started",
		src: "Windows",
	},
	"5444": {
		msg: "The following sub-layer was present when the Windows Filtering Platform Base Filtering Engine started",
		src: "Windows",
	},
	"5446": {
		msg: "A Windows Filtering Platform callout has been changed",
		src: "Windows",
	},
	"5447": {
		msg: "A Windows Filtering Platform filter has been changed",
		src: "Windows",
	},
	"5448": {
		msg: "A Windows Filtering Platform provider has been changed",
		src: "Windows",
	},
	"5449": {
		msg: "A Windows Filtering Platform provider context has been changed",
		src: "Windows",
	},
	"5450": {
		msg: "A Windows Filtering Platform sub-layer has been changed",
		src: "Windows",
	},
	"5451": {
		msg: "An IPsec Quick Mode security association was established",
		src: "Windows",
	},
	"5452": {
		msg: "An IPsec Quick Mode security association ended",
		src: "Windows",
	},
	"5453": {
		msg: "An IPsec negotiation with a remote computer failed because the IKE and AuthIP IPsec Keying Modules (IKEEXT) service is not started",
		src: "Windows",
	},
	"5456": {
		msg: "PAStore Engine applied Active Directory storage IPsec policy on the computer",
		src: "Windows",
	},
	"5457": {
		msg: "PAStore Engine failed to apply Active Directory storage IPsec policy on the computer",
		src: "Windows",
	},
	"5458": {
		msg: "PAStore Engine applied locally cached copy of Active Directory storage IPsec policy on the computer",
		src: "Windows",
	},
	"5459": {
		msg: "PAStore Engine failed to apply locally cached copy of Active Directory storage IPsec policy on the computer",
		src: "Windows",
	},
	"5460": {
		msg: "PAStore Engine applied local registry storage IPsec policy on the computer",
		src: "Windows",
	},
	"5461": {
		msg: "PAStore Engine failed to apply local registry storage IPsec policy on the computer",
		src: "Windows",
	},
	"5462": {
		msg: "PAStore Engine failed to apply some rules of the active IPsec policy on the computer",
		src: "Windows",
	},
	"5463": {
		msg: "PAStore Engine polled for changes to the active IPsec policy and detected no changes",
		src: "Windows",
	},
	"5464": {
		msg: "PAStore Engine polled for changes to the active IPsec policy, detected changes, and applied them to IPsec Services",
		src: "Windows",
	},
	"5465": {
		msg: "PAStore Engine received a control for forced reloading of IPsec policy and processed the control successfully",
		src: "Windows",
	},
	"5466": {
		msg: "PAStore Engine polled for changes to the Active Directory IPsec policy, determined that Active Directory cannot be reached, and will use the cached copy of the Active Directory IPsec policy instead",
		src: "Windows",
	},
	"5467": {
		msg: "PAStore Engine polled for changes to the Active Directory IPsec policy, determined that Active Directory can be reached, and found no changes to the policy",
		src: "Windows",
	},
	"5468": {
		msg: "PAStore Engine polled for changes to the Active Directory IPsec policy, determined that Active Directory can be reached, found changes to the policy, and applied those changes",
		src: "Windows",
	},
	"5471": {
		msg: "PAStore Engine loaded local storage IPsec policy on the computer",
		src: "Windows",
	},
	"5472": {
		msg: "PAStore Engine failed to load local storage IPsec policy on the computer",
		src: "Windows",
	},
	"5473": {
		msg: "PAStore Engine loaded directory storage IPsec policy on the computer",
		src: "Windows",
	},
	"5474": {
		msg: "PAStore Engine failed to load directory storage IPsec policy on the computer",
		src: "Windows",
	},
	"5477": {
		msg: "PAStore Engine failed to add quick mode filter",
		src: "Windows",
	},
	"5478": {
		msg: "IPsec Services has started successfully",
		src: "Windows",
	},
	"5479": {
		msg: "IPsec Services has been shut down successfully",
		src: "Windows",
	},
	"5480": {
		msg: "IPsec Services failed to get the complete list of network interfaces on the computer",
		src: "Windows",
	},
	"5483": {
		msg: "IPsec Services failed to initialize RPC server. IPsec Services could not be started",
		src: "Windows",
	},
	"5484": {
		msg: "IPsec Services has experienced a critical failure and has been shut down",
		src: "Windows",
	},
	"5485": {
		msg: "IPsec Services failed to process some IPsec filters on a plug-and-play event for network interfaces",
		src: "Windows",
	},
	"5632": {
		msg: "A request was made to authenticate to a wireless network",
		src: "Windows",
	},
	"5633": {
		msg: "A request was made to authenticate to a wired network",
		src: "Windows",
	},
	"5712": {
		msg: "A Remote Procedure Call (RPC) was attempted",
		src: "Windows",
	},
	"5888": {
		msg: "An object in the COM+ Catalog was modified",
		src: "Windows",
	},
	"5889": {
		msg: "An object was deleted from the COM+ Catalog",
		src: "Windows",
	},
	"5890": {
		msg: "An object was added to the COM+ Catalog",
		src: "Windows",
	},
	"6144": {
		msg: "Security policy in the group policy objects has been applied successfully",
		src: "Windows",
	},
	"6145": {
		msg: "One or more errors occured while processing security policy in the group policy objects",
		src: "Windows",
	},
	"6272": {
		msg: "Network Policy Server granted access to a user",
		src: "Windows",
	},
	"6273": {
		msg: "Network Policy Server denied access to a user",
		src: "Windows",
	},
	"6274": {
		msg: "Network Policy Server discarded the request for a user",
		src: "Windows",
	},
	"6275": {
		msg: "Network Policy Server discarded the accounting request for a user",
		src: "Windows",
	},
	"6276": {
		msg: "Network Policy Server quarantined a user",
		src: "Windows",
	},
	"6277": {
		msg: "Network Policy Server granted access to a user but put it on probation because the host did not meet the defined health policy",
		src: "Windows",
	},
	"6278": {
		msg: "Network Policy Server granted full access to a user because the host met the defined health policy",
		src: "Windows",
	},
	"6279": {
		msg: "Network Policy Server locked the user account due to repeated failed authentication attempts",
		src: "Windows",
	},
	"6280": {
		msg: "Network Policy Server unlocked the user account",
		src: "Windows",
	},
	"6281": {
		msg: "Code Integrity determined that the page hashes of an image file are not valid...",
		src: "Windows",
	},
	"6400": {
		msg: "BranchCache: Received an incorrectly formatted response while discovering availability of content.",
		src: "Windows",
	},
	"6401": {
		msg: "BranchCache: Received invalid data from a peer. Data discarded.",
		src: "Windows",
	},
	"6402": {
		msg: "BranchCache: The message to the hosted cache offering it data is incorrectly formatted.",
		src: "Windows",
	},
	"6403": {
		msg: "BranchCache: The hosted cache sent an incorrectly formatted response to the client's message to offer it data.",
		src: "Windows",
	},
	"6404": {
		msg: "BranchCache: Hosted cache could not be authenticated using the provisioned SSL certificate.",
		src: "Windows",
	},
	"6405": {
		msg: "BranchCache: %2 instance(s) of event id %1 occurred.",
		src: "Windows",
	},
	"6406": {
		msg: "%1 registered to Windows Firewall to control filtering for the following:",
		src: "Windows",
	},
	"6407": {
		msg: "%1",
		src: "Windows",
	},
	"6408": {
		msg: "Registered product %1 failed and Windows Firewall is now controlling the filtering for %2.",
		src: "Windows",
	},
	"6409": {
		msg: "BranchCache: A service connection point object could not be parsed",
		src: "Windows",
	},
	"6410": {
		msg: "Code integrity determined that a file does not meet the security requirements to load into a process. This could be due to the use of shared sections or other issues",
		src: "Windows",
	},
	"6416": {
		msg: "A new external device was recognized by the system.",
		src: "Windows",
	},
	"6417": {
		msg: "The FIPS mode crypto selftests succeeded",
		src: "Windows",
	},
	"6418": {
		msg: "The FIPS mode crypto selftests failed",
		src: "Windows",
	},
	"6419": {
		msg: "A request was made to disable a device",
		src: "Windows",
	},
	"6420": {
		msg: "A device was disabled",
		src: "Windows",
	},
	"6421": {
		msg: "A request was made to enable a device",
		src: "Windows",
	},
	"6422": {
		msg: "A device was enabled",
		src: "Windows",
	},
	"6423": {
		msg: "The installation of this device is forbidden by system policy",
		src: "Windows",
	},
	"6424": {
		msg: "The installation of this device was allowed, after having previously been forbidden by policy",
		src: "Windows",
	},
	"8191": {
		msg: "Highest System-Defined Audit Message Value",
		src: "Windows",
	},
	"11": {
		msg: "Site collection audit policy changed",
		src: "SharePoint",
	},
	"12": {
		msg: "Audit policy changed",
		src: "SharePoint",
	},
	"13": {
		msg: "Document checked in",
		src: "SharePoint",
	},
	"14": {
		msg: "Document checked out",
		src: "SharePoint",
	},
	"15": {
		msg: "Child object deleted",
		src: "SharePoint",
	},
	"16": {
		msg: "Child object moved",
		src: "SharePoint",
	},
	"17": {
		msg: "Object copied",
		src: "SharePoint",
	},
	"18": {
		msg: "Custom event",
		src: "SharePoint",
	},
	"19": {
		msg: "Object deleted",
		src: "SharePoint",
	},
	"20": {
		msg: "SharePoint audit logs deleted",
		src: "SharePoint",
	},
	"21": {
		msg: "Object moved",
		src: "SharePoint",
	},
	"22": {
		msg: "Object profile changed",
		src: "SharePoint",
	},
	"23": {
		msg: "SharePoint object structure changed",
		src: "SharePoint",
	},
	"24": {
		msg: "Search performed",
		src: "SharePoint",
	},
	"25": {
		msg: "SharePoint group created",
		src: "SharePoint",
	},
	"26": {
		msg: "SharePoint group deleted",
		src: "SharePoint",
	},
	"27": {
		msg: "SharePoint group member added",
		src: "SharePoint",
	},
	"28": {
		msg: "SharePoint group member removed",
		src: "SharePoint",
	},
	"29": {
		msg: "Unique permissions created",
		src: "SharePoint",
	},
	"30": {
		msg: "Unique permissions removed",
		src: "SharePoint",
	},
	"31": {
		msg: "Permissions updated",
		src: "SharePoint",
	},
	"32": {
		msg: "Permissions removed",
		src: "SharePoint",
	},
	"33": {
		msg: "Unique permission levels created",
		src: "SharePoint",
	},
	"34": {
		msg: "Permission level created",
		src: "SharePoint",
	},
	"35": {
		msg: "Permission level deleted",
		src: "SharePoint",
	},
	"36": {
		msg: "Permission level modified",
		src: "SharePoint",
	},
	"37": {
		msg: "SharePoint site collection administrator added",
		src: "SharePoint",
	},
	"38": {
		msg: "SharePoint site collection administrator removed",
		src: "SharePoint",
	},
	"39": {
		msg: "Object restored",
		src: "SharePoint",
	},
	"40": {
		msg: "Site collection updated",
		src: "SharePoint",
	},
	"41": {
		msg: "Web updated",
		src: "SharePoint",
	},
	"42": {
		msg: "Document library updated",
		src: "SharePoint",
	},
	"43": {
		msg: "Document updated",
		src: "SharePoint",
	},
	"44": {
		msg: "List updated",
		src: "SharePoint",
	},
	"45": {
		msg: "List item updated",
		src: "SharePoint",
	},
	"46": {
		msg: "Folder updated",
		src: "SharePoint",
	},
	"47": {
		msg: "Document viewed",
		src: "SharePoint",
	},
	"48": {
		msg: "Document library viewed",
		src: "SharePoint",
	},
	"49": {
		msg: "List viewed",
		src: "SharePoint",
	},
	"50": {
		msg: "Object viewed",
		src: "SharePoint",
	},
	"51": {
		msg: "Workflow accessed",
		src: "SharePoint",
	},
	"52": {
		msg: "Information management policy created",
		src: "SharePoint",
	},
	"53": {
		msg: "Information management policy changed",
		src: "SharePoint",
	},
	"54": {
		msg: "Site collection information management policy created",
		src: "SharePoint",
	},
	"55": {
		msg: "Site collection information management policy changed",
		src: "SharePoint",
	},
	"56": {
		msg: "Export of objects started",
		src: "SharePoint",
	},
	"57": {
		msg: "Export of objects completed",
		src: "SharePoint",
	},
	"58": {
		msg: "Import of objects started",
		src: "SharePoint",
	},
	"59": {
		msg: "Import of objects completed",
		src: "SharePoint",
	},
	"60": {
		msg: "Possible tampering warning",
		src: "SharePoint",
	},
	"61": {
		msg: "Retention policy processed",
		src: "SharePoint",
	},
	"62": {
		msg: "Document fragment updated",
		src: "SharePoint",
	},
	"63": {
		msg: "Content type imported",
		src: "SharePoint",
	},
	"64": {
		msg: "Information management policy deleted",
		src: "SharePoint",
	},
	"65": {
		msg: "Item declared as a record",
		src: "SharePoint",
	},
	"66": {
		msg: "Item undeclared as a record",
		src: "SharePoint",
	},
	"24000": {
		msg: "SQL audit event",
		src: "SQL Server",
	},
	"24001": {
		msg: "Login succeeded (action_id LGIS)",
		src: "SQL Server",
	},
	"24002": {
		msg: "Logout succeeded (action_id LGO)",
		src: "SQL Server",
	},
	"24003": {
		msg: "Login failed (action_id LGIF)",
		src: "SQL Server",
	},
	"24004": {
		msg: "Change own password succeeded (action_id PWCS; class_type LX)",
		src: "SQL Server",
	},
	"24005": {
		msg: "Change own password failed (action_id PWCS; class_type LX)",
		src: "SQL Server",
	},
	"24006": {
		msg: "Change password succeeded (action_id PWC class_type LX)",
		src: "SQL Server",
	},
	"24007": {
		msg: "Change password failed (action_id PWC class_type LX)",
		src: "SQL Server",
	},
	"24008": {
		msg: "Reset own password succeeded (action_id PWRS; class_type LX)",
		src: "SQL Server",
	},
	"24009": {
		msg: "Reset own password failed (action_id PWRS; class_type LX)",
		src: "SQL Server",
	},
	"24010": {
		msg: "Reset password succeeded (action_id PWR; class_type LX)",
		src: "SQL Server",
	},
	"24011": {
		msg: "Reset password failed (action_id PWR; class_type LX)",
		src: "SQL Server",
	},
	"24012": {
		msg: "Must change password (action_id PWMC)",
		src: "SQL Server",
	},
	"24013": {
		msg: "Account unlocked (action_id PWU)",
		src: "SQL Server",
	},
	"24014": {
		msg: "Change application role password succeeded (action_id PWC; class_type AR)",
		src: "SQL Server",
	},
	"24015": {
		msg: "Change application role password failed (action_id PWC class_type AR)",
		src: "SQL Server",
	},
	"24016": {
		msg: "Add member to server role succeeded (action_id APRL class_type SG)",
		src: "SQL Server",
	},
	"24017": {
		msg: "Add member to server role failed (action_id APRL class_type SG)",
		src: "SQL Server",
	},
	"24018": {
		msg: "Remove member from server role succeeded (action_id DPRL class_type SG)",
		src: "SQL Server",
	},
	"24019": {
		msg: "Remove member from server role failed (action_id DPRL class_type SG)",
		src: "SQL Server",
	},
	"24020": {
		msg: "Add member to database role succeeded (action_id APRL class_type RL)",
		src: "SQL Server",
	},
	"24021": {
		msg: "Add member to database role failed (action_id APRL class_type RL)",
		src: "SQL Server",
	},
	"24022": {
		msg: "Remove member from database role succeeded (action_id DPRL class_type RL)",
		src: "SQL Server",
	},
	"24023": {
		msg: "Remove member from database role failed (action_id DPRL class_type RL)",
		src: "SQL Server",
	},
	"24024": {
		msg: "Issued database backup command (action_id BA class_type DB)",
		src: "SQL Server",
	},
	"24025": {
		msg: "Issued transaction log backup command (action_id BAL)",
		src: "SQL Server",
	},
	"24026": {
		msg: "Issued database restore command (action_id RS class_type DB)",
		src: "SQL Server",
	},
	"24027": {
		msg: "Issued transaction log restore command (action_id RS class_type DB)",
		src: "SQL Server",
	},
	"24028": {
		msg: "Issued database console command (action_id DBCC)",
		src: "SQL Server",
	},
	"24029": {
		msg: "Issued a bulk administration command (action_id ADBO)",
		src: "SQL Server",
	},
	"24030": {
		msg: "Issued an alter connection command (action_id ALCN)",
		src: "SQL Server",
	},
	"24031": {
		msg: "Issued an alter resources command (action_id ALRS)",
		src: "SQL Server",
	},
	"24032": {
		msg: "Issued an alter server state command (action_id ALSS)",
		src: "SQL Server",
	},
	"24033": {
		msg: "Issued an alter server settings command (action_id ALST)",
		src: "SQL Server",
	},
	"24034": {
		msg: "Issued a view server state command (action_id VSST)",
		src: "SQL Server",
	},
	"24035": {
		msg: "Issued an external access assembly command (action_id XA)",
		src: "SQL Server",
	},
	"24036": {
		msg: "Issued an unsafe assembly command (action_id XU)",
		src: "SQL Server",
	},
	"24037": {
		msg: "Issued an alter reOrig governor command (action_id ALRS class_type RG)",
		src: "SQL Server",
	},
	"24038": {
		msg: "Issued a database authenticate command (action_id AUTH)",
		src: "SQL Server",
	},
	"24039": {
		msg: "Issued a database checkpoint command (action_id CP)",
		src: "SQL Server",
	},
	"24040": {
		msg: "Issued a database show plan command (action_id SPLN)",
		src: "SQL Server",
	},
	"24041": {
		msg: "Issued a subscribe to query information command (action_id SUQN)",
		src: "SQL Server",
	},
	"24042": {
		msg: "Issued a view database state command (action_id VDST)",
		src: "SQL Server",
	},
	"24043": {
		msg: "Issued a change server audit command (action_id AL class_type A)",
		src: "SQL Server",
	},
	"24044": {
		msg: "Issued a change server audit specification command (action_id AL class_type SA)",
		src: "SQL Server",
	},
	"24045": {
		msg: "Issued a change database audit specification command (action_id AL class_type DA)",
		src: "SQL Server",
	},
	"24046": {
		msg: "Issued a create server audit command (action_id CR class_type A)",
		src: "SQL Server",
	},
	"24047": {
		msg: "Issued a create server audit specification command (action_id CR class_type SA)",
		src: "SQL Server",
	},
	"24048": {
		msg: "Issued a create database audit specification command (action_id CR class_type DA)",
		src: "SQL Server",
	},
	"24049": {
		msg: "Issued a delete server audit command (action_id DR class_type A)",
		src: "SQL Server",
	},
	"24050": {
		msg: "Issued a delete server audit specification command (action_id DR class_type SA)",
		src: "SQL Server",
	},
	"24051": {
		msg: "Issued a delete database audit specification command (action_id DR class_type DA)",
		src: "SQL Server",
	},
	"24052": {
		msg: "Audit failure (action_id AUSF)",
		src: "SQL Server",
	},
	"24053": {
		msg: "Audit session changed (action_id AUSC)",
		src: "SQL Server",
	},
	"24054": {
		msg: "Started SQL server (action_id SVSR)",
		src: "SQL Server",
	},
	"24055": {
		msg: "Paused SQL server (action_id SVPD)",
		src: "SQL Server",
	},
	"24056": {
		msg: "Resumed SQL server (action_id SVCN)",
		src: "SQL Server",
	},
	"24057": {
		msg: "Stopped SQL server (action_id SVSD)",
		src: "SQL Server",
	},
	"24058": {
		msg: "Issued a create server object command (action_id CR; class_type AG, EP, SD, SE, T)",
		src: "SQL Server",
	},
	"24059": {
		msg: "Issued a change server object command (action_id AL; class_type AG, EP, SD, SE, T)",
		src: "SQL Server",
	},
	"24060": {
		msg: "Issued a delete server object command (action_id DR; class_type AG, EP, SD, SE, T)",
		src: "SQL Server",
	},
	"24061": {
		msg: "Issued a create server setting command (action_id CR class_type SR)",
		src: "SQL Server",
	},
	"24062": {
		msg: "Issued a change server setting command (action_id AL class_type SR)",
		src: "SQL Server",
	},
	"24063": {
		msg: "Issued a delete server setting command (action_id DR class_type SR)",
		src: "SQL Server",
	},
	"24064": {
		msg: "Issued a create server cryptographic provider command (action_id CR class_type CP)",
		src: "SQL Server",
	},
	"24065": {
		msg: "Issued a delete server cryptographic provider command (action_id DR class_type CP)",
		src: "SQL Server",
	},
	"24066": {
		msg: "Issued a change server cryptographic provider command (action_id AL class_type CP)",
		src: "SQL Server",
	},
	"24067": {
		msg: "Issued a create server credential command (action_id CR class_type CD)",
		src: "SQL Server",
	},
	"24068": {
		msg: "Issued a delete server credential command (action_id DR class_type CD)",
		src: "SQL Server",
	},
	"24069": {
		msg: "Issued a change server credential command (action_id AL class_type CD)",
		src: "SQL Server",
	},
	"24070": {
		msg: "Issued a change server master key command (action_id AL class_type MK)",
		src: "SQL Server",
	},
	"24071": {
		msg: "Issued a back up server master key command (action_id BA class_type MK)",
		src: "SQL Server",
	},
	"24072": {
		msg: "Issued a restore server master key command (action_id RS class_type MK)",
		src: "SQL Server",
	},
	"24073": {
		msg: "Issued a map server credential to login command (action_id CMLG)",
		src: "SQL Server",
	},
	"24074": {
		msg: "Issued a remove map between server credential and login command (action_id NMLG)",
		src: "SQL Server",
	},
	"24075": {
		msg: "Issued a create server principal command (action_id CR class_type LX, SL)",
		src: "SQL Server",
	},
	"24076": {
		msg: "Issued a delete server principal command (action_id DR class_type LX, SL)",
		src: "SQL Server",
	},
	"24077": {
		msg: "Issued a change server principal credentials command (action_id CCLG)",
		src: "SQL Server",
	},
	"24078": {
		msg: "Issued a disable server principal command (action_id LGDA)",
		src: "SQL Server",
	},
	"24079": {
		msg: "Issued a change server principal default database command (action_id LGDB)",
		src: "SQL Server",
	},
	"24080": {
		msg: "Issued an enable server principal command (action_id LGEA)",
		src: "SQL Server",
	},
	"24081": {
		msg: "Issued a change server principal default language command (action_id LGLG)",
		src: "SQL Server",
	},
	"24082": {
		msg: "Issued a change server principal password expiration command (action_id PWEX)",
		src: "SQL Server",
	},
	"24083": {
		msg: "Issued a change server principal password policy command (action_id PWPL)",
		src: "SQL Server",
	},
	"24084": {
		msg: "Issued a change server principal Data command (action_id LGNM)",
		src: "SQL Server",
	},
	"24085": {
		msg: "Issued a create database command (action_id CR class_type DB)",
		src: "SQL Server",
	},
	"24086": {
		msg: "Issued a change database command (action_id AL class_type DB)",
		src: "SQL Server",
	},
	"24087": {
		msg: "Issued a delete database command (action_id DR class_type DB)",
		src: "SQL Server",
	},
	"24088": {
		msg: "Issued a create certificate command (action_id CR class_type CR)",
		src: "SQL Server",
	},
	"24089": {
		msg: "Issued a change certificate command (action_id AL class_type CR)",
		src: "SQL Server",
	},
	"24090": {
		msg: "Issued a delete certificate command (action_id DR class_type CR)",
		src: "SQL Server",
	},
	"24091": {
		msg: "Issued a back up certificate command (action_id BA class_type CR)",
		src: "SQL Server",
	},
	"24092": {
		msg: "Issued an access certificate command (action_id AS class_type CR)",
		src: "SQL Server",
	},
	"24093": {
		msg: "Issued a create asymmetric key command (action_id CR class_type AK)",
		src: "SQL Server",
	},
	"24094": {
		msg: "Issued a change asymmetric key command (action_id AL class_type AK)",
		src: "SQL Server",
	},
	"24095": {
		msg: "Issued a delete asymmetric key command (action_id DR class_type AK)",
		src: "SQL Server",
	},
	"24096": {
		msg: "Issued an access asymmetric key command (action_id AS class_type AK)",
		src: "SQL Server",
	},
	"24097": {
		msg: "Issued a create database master key command (action_id CR class_type MK)",
		src: "SQL Server",
	},
	"24098": {
		msg: "Issued a change database master key command (action_id AL class_type MK)",
		src: "SQL Server",
	},
	"24099": {
		msg: "Issued a delete database master key command (action_id DR class_type MK)",
		src: "SQL Server",
	},
	"24100": {
		msg: "Issued a back up database master key command (action_id BA class_type MK)",
		src: "SQL Server",
	},
	"24101": {
		msg: "Issued a restore database master key command (action_id RS class_type MK)",
		src: "SQL Server",
	},
	"24102": {
		msg: "Issued an open database master key command (action_id OP class_type MK)",
		src: "SQL Server",
	},
	"24103": {
		msg: "Issued a create database symmetric key command (action_id CR class_type SK)",
		src: "SQL Server",
	},
	"24104": {
		msg: "Issued a change database symmetric key command (action_id AL class_type SK)",
		src: "SQL Server",
	},
	"24105": {
		msg: "Issued a delete database symmetric key command (action_id DR class_type SK)",
		src: "SQL Server",
	},
	"24106": {
		msg: "Issued a back up database symmetric key command (action_id BA class_type SK)",
		src: "SQL Server",
	},
	"24107": {
		msg: "Issued an open database symmetric key command (action_id OP class_type SK)",
		src: "SQL Server",
	},
	"24108": {
		msg: "Issued a create database object command (action_id CR)",
		src: "SQL Server",
	},
	"24109": {
		msg: "Issued a change database object command (action_id AL)",
		src: "SQL Server",
	},
	"24110": {
		msg: "Issued a delete database object command (action_id DR)",
		src: "SQL Server",
	},
	"24111": {
		msg: "Issued an access database object command (action_id AS)",
		src: "SQL Server",
	},
	"24112": {
		msg: "Issued a create assembly command (action_id CR class_type AS)",
		src: "SQL Server",
	},
	"24113": {
		msg: "Issued a change assembly command (action_id AL class_type AS)",
		src: "SQL Server",
	},
	"24114": {
		msg: "Issued a delete assembly command (action_id DR class_type AS)",
		src: "SQL Server",
	},
	"24115": {
		msg: "Issued a create schema command (action_id CR class_type SC)",
		src: "SQL Server",
	},
	"24116": {
		msg: "Issued a change schema command (action_id AL class_type SC)",
		src: "SQL Server",
	},
	"24117": {
		msg: "Issued a delete schema command (action_id DR class_type SC)",
		src: "SQL Server",
	},
	"24118": {
		msg: "Issued a create database encryption key command (action_id CR class_type DK)",
		src: "SQL Server",
	},
	"24119": {
		msg: "Issued a change database encryption key command (action_id AL class_type DK)",
		src: "SQL Server",
	},
	"24120": {
		msg: "Issued a delete database encryption key command (action_id DR class_type DK)",
		src: "SQL Server",
	},
	"24121": {
		msg: "Issued a create database user command (action_id CR; class_type US)",
		src: "SQL Server",
	},
	"24122": {
		msg: "Issued a change database user command (action_id AL; class_type US)",
		src: "SQL Server",
	},
	"24123": {
		msg: "Issued a delete database user command (action_id DR; class_type US)",
		src: "SQL Server",
	},
	"24124": {
		msg: "Issued a create database role command (action_id CR class_type RL)",
		src: "SQL Server",
	},
	"24125": {
		msg: "Issued a change database role command (action_id AL class_type RL)",
		src: "SQL Server",
	},
	"24126": {
		msg: "Issued a delete database role command (action_id DR class_type RL)",
		src: "SQL Server",
	},
	"24127": {
		msg: "Issued a create application role command (action_id CR class_type AR)",
		src: "SQL Server",
	},
	"24128": {
		msg: "Issued a change application role command (action_id AL class_type AR)",
		src: "SQL Server",
	},
	"24129": {
		msg: "Issued a delete application role command (action_id DR class_type AR)",
		src: "SQL Server",
	},
	"24130": {
		msg: "Issued a change database user login command (action_id USAF)",
		src: "SQL Server",
	},
	"24131": {
		msg: "Issued an auto-change database user login command (action_id USLG)",
		src: "SQL Server",
	},
	"24132": {
		msg: "Issued a create schema object command (action_id CR class_type D)",
		src: "SQL Server",
	},
	"24133": {
		msg: "Issued a change schema object command (action_id AL class_type D)",
		src: "SQL Server",
	},
	"24134": {
		msg: "Issued a delete schema object command (action_id DR class_type D)",
		src: "SQL Server",
	},
	"24135": {
		msg: "Issued a transfer schema object command (action_id TRO class_type D)",
		src: "SQL Server",
	},
	"24136": {
		msg: "Issued a create schema type command (action_id CR class_type TY)",
		src: "SQL Server",
	},
	"24137": {
		msg: "Issued a change schema type command (action_id AL class_type TY)",
		src: "SQL Server",
	},
	"24138": {
		msg: "Issued a delete schema type command (action_id DR class_type TY)",
		src: "SQL Server",
	},
	"24139": {
		msg: "Issued a transfer schema type command (action_id TRO class_type TY)",
		src: "SQL Server",
	},
	"24140": {
		msg: "Issued a create XML schema collection command (action_id CR class_type SX)",
		src: "SQL Server",
	},
	"24141": {
		msg: "Issued a change XML schema collection command (action_id AL class_type SX)",
		src: "SQL Server",
	},
	"24142": {
		msg: "Issued a delete XML schema collection command (action_id DR class_type SX)",
		src: "SQL Server",
	},
	"24143": {
		msg: "Issued a transfer XML schema collection command (action_id TRO class_type SX)",
		src: "SQL Server",
	},
	"24144": {
		msg: "Issued an impersonate within server scope command (action_id IMP; class_type LX)",
		src: "SQL Server",
	},
	"24145": {
		msg: "Issued an impersonate within database scope command (action_id IMP; class_type US)",
		src: "SQL Server",
	},
	"24146": {
		msg: "Issued a change server object owner command (action_id TO class_type SG)",
		src: "SQL Server",
	},
	"24147": {
		msg: "Issued a change database owner command (action_id TO class_type DB)",
		src: "SQL Server",
	},
	"24148": {
		msg: "Issued a change schema owner command (action_id TO class_type SC)",
		src: "SQL Server",
	},
	"24150": {
		msg: "Issued a change role owner command (action_id TO class_type RL)",
		src: "SQL Server",
	},
	"24151": {
		msg: "Issued a change database object owner command (action_id TO)",
		src: "SQL Server",
	},
	"24152": {
		msg: "Issued a change symmetric key owner command (action_id TO class_type SK)",
		src: "SQL Server",
	},
	"24153": {
		msg: "Issued a change certificate owner command (action_id TO class_type CR)",
		src: "SQL Server",
	},
	"24154": {
		msg: "Issued a change asymmetric key owner command (action_id TO class_type AK)",
		src: "SQL Server",
	},
	"24155": {
		msg: "Issued a change schema object owner command (action_id TO class_type OB)",
		src: "SQL Server",
	},
	"24156": {
		msg: "Issued a change schema type owner command (action_id TO class_type TY)",
		src: "SQL Server",
	},
	"24157": {
		msg: "Issued a change XML schema collection owner command (action_id TO class_type SX)",
		src: "SQL Server",
	},
	"24158": {
		msg: "Grant server permissions succeeded (action_id G class_type SR)",
		src: "SQL Server",
	},
	"24159": {
		msg: "Grant server permissions failed (action_id G class_type SR)",
		src: "SQL Server",
	},
	"24160": {
		msg: "Grant server permissions with grant succeeded (action_id GWG class_type SR)",
		src: "SQL Server",
	},
	"24161": {
		msg: "Grant server permissions with grant failed (action_id GWG class_type SR)",
		src: "SQL Server",
	},
	"24162": {
		msg: "Deny server permissions succeeded (action_id D class_type SR)",
		src: "SQL Server",
	},
	"24163": {
		msg: "Deny server permissions failed (action_id D class_type SR)",
		src: "SQL Server",
	},
	"24164": {
		msg: "Deny server permissions with cascade succeeded (action_id DWC class_type SR)",
		src: "SQL Server",
	},
	"24165": {
		msg: "Deny server permissions with cascade failed (action_id DWC class_type SR)",
		src: "SQL Server",
	},
	"24166": {
		msg: "Revoke server permissions succeeded (action_id R class_type SR)",
		src: "SQL Server",
	},
	"24167": {
		msg: "Revoke server permissions failed (action_id R class_type SR)",
		src: "SQL Server",
	},
	"24168": {
		msg: "Revoke server permissions with grant succeeded (action_id RWG class_type SR)",
		src: "SQL Server",
	},
	"24169": {
		msg: "Revoke server permissions with grant failed (action_id RWG class_type SR)",
		src: "SQL Server",
	},
	"24170": {
		msg: "Revoke server permissions with cascade succeeded (action_id RWC class_type SR)",
		src: "SQL Server",
	},
	"24171": {
		msg: "Revoke server permissions with cascade failed (action_id RWC class_type SR)",
		src: "SQL Server",
	},
	"24172": {
		msg: "Issued grant server object permissions command (action_id G; class_type LX)",
		src: "SQL Server",
	},
	"24173": {
		msg: "Issued grant server object permissions with grant command (action_id GWG; class_type LX)",
		src: "SQL Server",
	},
	"24174": {
		msg: "Issued deny server object permissions command (action_id D; class_type LX)",
		src: "SQL Server",
	},
	"24175": {
		msg: "Issued deny server object permissions with cascade command (action_id DWC; class_type LX)",
		src: "SQL Server",
	},
	"24176": {
		msg: "Issued revoke server object permissions command (action_id R; class_type LX)",
		src: "SQL Server",
	},
	"24177": {
		msg: "Issued revoke server object permissions with grant command (action_id; RWG class_type LX)",
		src: "SQL Server",
	},
	"24178": {
		msg: "Issued revoke server object permissions with cascade command (action_id RWC; class_type LX)",
		src: "SQL Server",
	},
	"24179": {
		msg: "Grant database permissions succeeded (action_id G class_type DB)",
		src: "SQL Server",
	},
	"24180": {
		msg: "Grant database permissions failed (action_id G class_type DB)",
		src: "SQL Server",
	},
	"24181": {
		msg: "Grant database permissions with grant succeeded (action_id GWG class_type DB)",
		src: "SQL Server",
	},
	"24182": {
		msg: "Grant database permissions with grant failed (action_id GWG class_type DB)",
		src: "SQL Server",
	},
	"24183": {
		msg: "Deny database permissions succeeded (action_id D class_type DB)",
		src: "SQL Server",
	},
	"24184": {
		msg: "Deny database permissions failed (action_id D class_type DB)",
		src: "SQL Server",
	},
	"24185": {
		msg: "Deny database permissions with cascade succeeded (action_id DWC class_type DB)",
		src: "SQL Server",
	},
	"24186": {
		msg: "Deny database permissions with cascade failed (action_id DWC class_type DB)",
		src: "SQL Server",
	},
	"24187": {
		msg: "Revoke database permissions succeeded (action_id R class_type DB)",
		src: "SQL Server",
	},
	"24188": {
		msg: "Revoke database permissions failed (action_id R class_type DB)",
		src: "SQL Server",
	},
	"24189": {
		msg: "Revoke database permissions with grant succeeded (action_id RWG class_type DB)",
		src: "SQL Server",
	},
	"24190": {
		msg: "Revoke database permissions with grant failed (action_id RWG class_type DB)",
		src: "SQL Server",
	},
	"24191": {
		msg: "Revoke database permissions with cascade succeeded (action_id RWC class_type DB)",
		src: "SQL Server",
	},
	"24192": {
		msg: "Revoke database permissions with cascade failed (action_id RWC class_type DB)",
		src: "SQL Server",
	},
	"24193": {
		msg: "Issued grant database object permissions command (action_id G class_type US)",
		src: "SQL Server",
	},
	"24194": {
		msg: "Issued grant database object permissions with grant command (action_id GWG; class_type US)",
		src: "SQL Server",
	},
	"24195": {
		msg: "Issued deny database object permissions command (action_id D; class_type US)",
		src: "SQL Server",
	},
	"24196": {
		msg: "Issued deny database object permissions with cascade command (action_id DWC; class_type US)",
		src: "SQL Server",
	},
	"24197": {
		msg: "Issued revoke database object permissions command (action_id R; class_type US)",
		src: "SQL Server",
	},
	"24198": {
		msg: "Issued revoke database object permissions with grant command (action_id RWG; class_type US)",
		src: "SQL Server",
	},
	"24199": {
		msg: "Issued revoke database object permissions with cascade command (action_id RWC; class_type US)",
		src: "SQL Server",
	},
	"24200": {
		msg: "Issued grant schema permissions command (action_id G class_type SC)",
		src: "SQL Server",
	},
	"24201": {
		msg: "Issued grant schema permissions with grant command (action_id GWG class_type SC)",
		src: "SQL Server",
	},
	"24202": {
		msg: "Issued deny schema permissions command (action_id D class_type SC)",
		src: "SQL Server",
	},
	"24203": {
		msg: "Issued deny schema permissions with cascade command (action_id DWC class_type SC)",
		src: "SQL Server",
	},
	"24204": {
		msg: "Issued revoke schema permissions command (action_id R class_type SC)",
		src: "SQL Server",
	},
	"24205": {
		msg: "Issued revoke schema permissions with grant command (action_id RWG class_type SC)",
		src: "SQL Server",
	},
	"24206": {
		msg: "Issued revoke schema permissions with cascade command (action_id RWC class_type SC)",
		src: "SQL Server",
	},
	"24207": {
		msg: "Issued grant assembly permissions command (action_id G class_type AS)",
		src: "SQL Server",
	},
	"24208": {
		msg: "Issued grant assembly permissions with grant command (action_id GWG class_type AS)",
		src: "SQL Server",
	},
	"24209": {
		msg: "Issued deny assembly permissions command (action_id D class_type AS)",
		src: "SQL Server",
	},
	"24210": {
		msg: "Issued deny assembly permissions with cascade command (action_id DWC class_type AS)",
		src: "SQL Server",
	},
	"24211": {
		msg: "Issued revoke assembly permissions command (action_id R class_type AS)",
		src: "SQL Server",
	},
	"24212": {
		msg: "Issued revoke assembly permissions with grant command (action_id RWG class_type AS)",
		src: "SQL Server",
	},
	"24213": {
		msg: "Issued revoke assembly permissions with cascade command (action_id RWC class_type AS)",
		src: "SQL Server",
	},
	"24214": {
		msg: "Issued grant database role permissions command (action_id G class_type RL)",
		src: "SQL Server",
	},
	"24215": {
		msg: "Issued grant database role permissions with grant command (action_id GWG class_type RL)",
		src: "SQL Server",
	},
	"24216": {
		msg: "Issued deny database role permissions command (action_id D class_type RL)",
		src: "SQL Server",
	},
	"24217": {
		msg: "Issued deny database role permissions with cascade command (action_id DWC class_type RL)",
		src: "SQL Server",
	},
	"24218": {
		msg: "Issued revoke database role permissions command (action_id R class_type RL)",
		src: "SQL Server",
	},
	"24219": {
		msg: "Issued revoke database role permissions with grant command (action_id RWG class_type RL)",
		src: "SQL Server",
	},
	"24220": {
		msg: "Issued revoke database role permissions with cascade command (action_id RWC class_type RL)",
		src: "SQL Server",
	},
	"24221": {
		msg: "Issued grant application role permissions command (action_id G class_type AR)",
		src: "SQL Server",
	},
	"24222": {
		msg: "Issued grant application role permissions with grant command (action_id GWG class_type AR)",
		src: "SQL Server",
	},
	"24223": {
		msg: "Issued deny application role permissions command (action_id D class_type AR)",
		src: "SQL Server",
	},
	"24224": {
		msg: "Issued deny application role permissions with cascade command (action_id DWC class_type AR)",
		src: "SQL Server",
	},
	"24225": {
		msg: "Issued revoke application role permissions command (action_id R class_type AR)",
		src: "SQL Server",
	},
	"24226": {
		msg: "Issued revoke application role permissions with grant command (action_id RWG class_type AR)",
		src: "SQL Server",
	},
	"24227": {
		msg: "Issued revoke application role permissions with cascade command (action_id RWC class_type AR)",
		src: "SQL Server",
	},
	"24228": {
		msg: "Issued grant symmetric key permissions command (action_id G class_type SK)",
		src: "SQL Server",
	},
	"24229": {
		msg: "Issued grant symmetric key permissions with grant command (action_id GWG class_type SK)",
		src: "SQL Server",
	},
	"24230": {
		msg: "Issued deny symmetric key permissions command (action_id D class_type SK)",
		src: "SQL Server",
	},
	"24231": {
		msg: "Issued deny symmetric key permissions with cascade command (action_id DWC class_type SK)",
		src: "SQL Server",
	},
	"24232": {
		msg: "Issued revoke symmetric key permissions command (action_id R class_type SK)",
		src: "SQL Server",
	},
	"24233": {
		msg: "Issued revoke symmetric key permissions with grant command (action_id RWG class_type SK)",
		src: "SQL Server",
	},
	"24234": {
		msg: "Issued revoke symmetric key permissions with cascade command (action_id RWC class_type SK)",
		src: "SQL Server",
	},
	"24235": {
		msg: "Issued grant certificate permissions command (action_id G class_type CR)",
		src: "SQL Server",
	},
	"24236": {
		msg: "Issued grant certificate permissions with grant command (action_id GWG class_type CR)",
		src: "SQL Server",
	},
	"24237": {
		msg: "Issued deny certificate permissions command (action_id D class_type CR)",
		src: "SQL Server",
	},
	"24238": {
		msg: "Issued deny certificate permissions with cascade command (action_id DWC class_type CR)",
		src: "SQL Server",
	},
	"24239": {
		msg: "Issued revoke certificate permissions command (action_id R class_type CR)",
		src: "SQL Server",
	},
	"24240": {
		msg: "Issued revoke certificate permissions with grant command (action_id RWG class_type CR)",
		src: "SQL Server",
	},
	"24241": {
		msg: "Issued revoke certificate permissions with cascade command (action_id RWC class_type CR)",
		src: "SQL Server",
	},
	"24242": {
		msg: "Issued grant asymmetric key permissions command (action_id G class_type AK)",
		src: "SQL Server",
	},
	"24243": {
		msg: "Issued grant asymmetric key permissions with grant command (action_id GWG class_type AK)",
		src: "SQL Server",
	},
	"24244": {
		msg: "Issued deny asymmetric key permissions command (action_id D class_type AK)",
		src: "SQL Server",
	},
	"24245": {
		msg: "Issued deny asymmetric key permissions with cascade command (action_id DWC class_type AK)",
		src: "SQL Server",
	},
	"24246": {
		msg: "Issued revoke asymmetric key permissions command (action_id R class_type AK)",
		src: "SQL Server",
	},
	"24247": {
		msg: "Issued revoke asymmetric key permissions with grant command (action_id RWG class_type AK)",
		src: "SQL Server",
	},
	"24248": {
		msg: "Issued revoke asymmetric key permissions with cascade command (action_id RWC class_type AK)",
		src: "SQL Server",
	},
	"24249": {
		msg: "Issued grant schema object permissions command (action_id G class_type OB)",
		src: "SQL Server",
	},
	"24250": {
		msg: "Issued grant schema object permissions with grant command (action_id GWG class_type OB)",
		src: "SQL Server",
	},
	"24251": {
		msg: "Issued deny schema object permissions command (action_id D class_type OB)",
		src: "SQL Server",
	},
	"24252": {
		msg: "Issued deny schema object permissions with cascade command (action_id DWC class_type OB)",
		src: "SQL Server",
	},
	"24253": {
		msg: "Issued revoke schema object permissions command (action_id R class_type OB)",
		src: "SQL Server",
	},
	"24254": {
		msg: "Issued revoke schema object permissions with grant command (action_id RWG class_type OB)",
		src: "SQL Server",
	},
	"24255": {
		msg: "Issued revoke schema object permissions with cascade command (action_id RWC class_type OB)",
		src: "SQL Server",
	},
	"24256": {
		msg: "Issued grant schema type permissions command (action_id G class_type TY)",
		src: "SQL Server",
	},
	"24257": {
		msg: "Issued grant schema type permissions with grant command (action_id GWG class_type TY)",
		src: "SQL Server",
	},
	"24258": {
		msg: "Issued deny schema type permissions command (action_id D class_type TY)",
		src: "SQL Server",
	},
	"24259": {
		msg: "Issued deny schema type permissions with cascade command (action_id DWC class_type TY)",
		src: "SQL Server",
	},
	"24260": {
		msg: "Issued revoke schema type permissions command (action_id R class_type TY)",
		src: "SQL Server",
	},
	"24261": {
		msg: "Issued revoke schema type permissions with grant command (action_id RWG class_type TY)",
		src: "SQL Server",
	},
	"24262": {
		msg: "Issued revoke schema type permissions with cascade command (action_id RWC class_type TY)",
		src: "SQL Server",
	},
	"24263": {
		msg: "Issued grant XML schema collection permissions command (action_id G class_type SX)",
		src: "SQL Server",
	},
	"24264": {
		msg: "Issued grant XML schema collection permissions with grant command (action_id GWG class_type SX)",
		src: "SQL Server",
	},
	"24265": {
		msg: "Issued deny XML schema collection permissions command (action_id D class_type SX)",
		src: "SQL Server",
	},
	"24266": {
		msg: "Issued deny XML schema collection permissions with cascade command (action_id DWC class_type SX)",
		src: "SQL Server",
	},
	"24267": {
		msg: "Issued revoke XML schema collection permissions command (action_id R class_type SX)",
		src: "SQL Server",
	},
	"24268": {
		msg: "Issued revoke XML schema collection permissions with grant command (action_id RWG class_type SX)",
		src: "SQL Server",
	},
	"24269": {
		msg: "Issued revoke XML schema collection permissions with cascade command (action_id RWC class_type SX)",
		src: "SQL Server",
	},
	"24270": {
		msg: "Issued reference database object permissions command (action_id RF)",
		src: "SQL Server",
	},
	"24271": {
		msg: "Issued send service request command (action_id SN)",
		src: "SQL Server",
	},
	"24272": {
		msg: "Issued check permissions with schema command (action_id VWCT)",
		src: "SQL Server",
	},
	"24273": {
		msg: "Issued use service broker transport security command (action_id LGB)",
		src: "SQL Server",
	},
	"24274": {
		msg: "Issued use database mirroring transport security command (action_id LGM)",
		src: "SQL Server",
	},
	"24275": {
		msg: "Issued alter trace command (action_id ALTR)",
		src: "SQL Server",
	},
	"24276": {
		msg: "Issued start trace command (action_id TASA)",
		src: "SQL Server",
	},
	"24277": {
		msg: "Issued stop trace command (action_id TASP)",
		src: "SQL Server",
	},
	"24278": {
		msg: "Issued enable trace C2 audit mode command (action_id C2ON)",
		src: "SQL Server",
	},
	"24279": {
		msg: "Issued disable trace C2 audit mode command (action_id C2OF)",
		src: "SQL Server",
	},
	"24280": {
		msg: "Issued server full-text command (action_id FT)",
		src: "SQL Server",
	},
	"24281": {
		msg: "Issued select command (action_id SL)",
		src: "SQL Server",
	},
	"24282": {
		msg: "Issued update command (action_id UP)",
		src: "SQL Server",
	},
	"24283": {
		msg: "Issued insert command (action_id IN)",
		src: "SQL Server",
	},
	"24284": {
		msg: "Issued delete command (action_id DL)",
		src: "SQL Server",
	},
	"24285": {
		msg: "Issued execute command (action_id EX)",
		src: "SQL Server",
	},
	"24286": {
		msg: "Issued receive command (action_id RC)",
		src: "SQL Server",
	},
	"24287": {
		msg: "Issued check references command (action_id RF)",
		src: "SQL Server",
	},
	"24288": {
		msg: "Issued a create user-defined server role command (action_id CR class_type SG)",
		src: "SQL Server",
	},
	"24289": {
		msg: "Issued a change user-defined server role command (action_id AL class_type SG)",
		src: "SQL Server",
	},
	"24290": {
		msg: "Issued a delete user-defined server role command (action_id DR class_type SG)",
		src: "SQL Server",
	},
	"24291": {
		msg: "Issued grant user-defined server role permissions command (action_id G class_type SG)",
		src: "SQL Server",
	},
	"24292": {
		msg: "Issued grant user-defined server role permissions with grant command (action_id GWG class_type SG)",
		src: "SQL Server",
	},
	"24293": {
		msg: "Issued deny user-defined server role permissions command (action_id D class_type SG)",
		src: "SQL Server",
	},
	"24294": {
		msg: "Issued deny user-defined server role permissions with cascade command (action_id DWC class_type SG)",
		src: "SQL Server",
	},
	"24295": {
		msg: "Issued revoke user-defined server role permissions command (action_id R class_type SG)",
		src: "SQL Server",
	},
	"24296": {
		msg: "Issued revoke user-defined server role permissions with grant command (action_id RWG class_type SG)",
		src: "SQL Server",
	},
	"24297": {
		msg: "Issued revoke user-defined server role permissions with cascade command (action_id RWC class_type SG)",
		src: "SQL Server",
	},
	"24298": {
		msg: "Database login succeeded (action_id DBAS)",
		src: "SQL Server",
	},
	"24299": {
		msg: "Database login failed (action_id DBAF)",
		src: "SQL Server",
	},
	"24300": {
		msg: "Database logout successful (action_id DAGL)",
		src: "SQL Server",
	},
	"24301": {
		msg: "Change password succeeded (action_id PWC; class_type US)",
		src: "SQL Server",
	},
	"24302": {
		msg: "Change password failed (action_id PWC; class_type US)",
		src: "SQL Server",
	},
	"24303": {
		msg: "Change own password succeeded (action_id PWCS; class_type US)",
		src: "SQL Server",
	},
	"24304": {
		msg: "Change own password failed (action_id PWCS; class_type US)",
		src: "SQL Server",
	},
	"24305": {
		msg: "Reset own password succeeded (action_id PWRS; class_type US)",
		src: "SQL Server",
	},
	"24306": {
		msg: "Reset own password failed (action_id PWRS; class_type US)",
		src: "SQL Server",
	},
	"24307": {
		msg: "Reset password succeeded (action_id PWR; class_type US)",
		src: "SQL Server",
	},
	"24308": {
		msg: "Reset password failed (action_id PWR; class_type US)",
		src: "SQL Server",
	},
	"24309": {
		msg: "Copy password (action_id USTC)",
		src: "SQL Server",
	},
	"24310": {
		msg: "User-defined SQL audit event (action_id UDAU)",
		src: "SQL Server",
	},
	"24311": {
		msg: "Issued a change database audit command (action_id AL class_type DU)",
		src: "SQL Server",
	},
	"24312": {
		msg: "Issued a create database audit command (action_id CR class_type DU)",
		src: "SQL Server",
	},
	"24313": {
		msg: "Issued a delete database audit command (action_id DR class_type DU)",
		src: "SQL Server",
	},
	"24314": {
		msg: "Issued a begin transaction command (action_id TXBG)",
		src: "SQL Server",
	},
	"24315": {
		msg: "Issued a commit transaction command (action_id TXCM)",
		src: "SQL Server",
	},
	"24316": {
		msg: "Issued a rollback transaction command (action_id TXRB)",
		src: "SQL Server",
	},
	"24317": {
		msg: "Issued a create column master key command (action_id CR; class_type CM)",
		src: "SQL Server",
	},
	"24318": {
		msg: "Issued a delete column master key command (action_id DR; class_type CM)",
		src: "SQL Server",
	},
	"24319": {
		msg: "A column master key was viewed (action_id VW; class_type CM)",
		src: "SQL Server",
	},
	"24320": {
		msg: "Issued a create column encryption key command (action_id CR; class_type CK)",
		src: "SQL Server",
	},
	"24321": {
		msg: "Issued a change column encryption key command (action_id AL; class_type CK)",
		src: "SQL Server",
	},
	"24322": {
		msg: "Issued a delete column encryption key command (action_id DR; class_type CK)",
		src: "SQL Server",
	},
	"24323": {
		msg: "A column encryption key was viewed (action_id VW; class_type CK)",
		src: "SQL Server",
	},
	"24324": {
		msg: "Issued a create database credential command (action_id CR; class_type DC)",
		src: "SQL Server",
	},
	"24325": {
		msg: "Issued a change database credential command (action_id AL; class_type DC)",
		src: "SQL Server",
	},
	"24326": {
		msg: "Issued a delete database credential command (action_id DR; class_type DC)",
		src: "SQL Server",
	},
	"24327": {
		msg: "Issued a change database scoped configuration command (action_id AL; class_type DS)",
		src: "SQL Server",
	},
	"24328": {
		msg: "Issued a create external data Orig command (action_id CR; class_type ED)",
		src: "SQL Server",
	},
	"24329": {
		msg: "Issued a change external data Orig command (action_id AL; class_type ED)",
		src: "SQL Server",
	},
	"24330": {
		msg: "Issued a delete external data Orig command (action_id DR; class_type ED)",
		src: "SQL Server",
	},
	"24331": {
		msg: "Issued a create external file format command (action_id CR; class_type EF)",
		src: "SQL Server",
	},
	"24332": {
		msg: "Issued a delete external file format command (action_id DR; class_type EF)",
		src: "SQL Server",
	},
	"24333": {
		msg: "Issued a create external reOrig pool command (action_id CR; class_type ER)",
		src: "SQL Server",
	},
	"24334": {
		msg: "Issued a change external reOrig pool command (action_id AL; class_type ER)",
		src: "SQL Server",
	},
	"24335": {
		msg: "Issued a delete external reOrig pool command (action_id DR; class_type ER)",
		src: "SQL Server",
	},
	"24337": {
		msg: "Global transaction login (action_id LGG)",
		src: "SQL Server",
	},
	"24338": {
		msg: "Grant permissions on a database scoped credential succeeded (action_id G; class_type DC)",
		src: "SQL Server",
	},
	"24339": {
		msg: "Grant permissions on a database scoped credential failed (action_id G; class_type DC)",
		src: "SQL Server",
	},
	"24340": {
		msg: "Grant permissions on a database scoped credential with grant succeeded (action_id GWG; class_type DC)",
		src: "SQL Server",
	},
	"24341": {
		msg: "Grant permissions on a database scoped credential with grant failed (action_id GWG; class_type DC)",
		src: "SQL Server",
	},
	"24342": {
		msg: "Deny permissions on a database scoped credential succeeded (action_id D; class_type DC)",
		src: "SQL Server",
	},
	"24343": {
		msg: "Deny permissions on a database scoped credential failed (action_id D; class_type DC)",
		src: "SQL Server",
	},
	"24344": {
		msg: "Deny permissions on a database scoped credential with cascade succeeded (action_id DWC; class_type DC)",
		src: "SQL Server",
	},
	"24345": {
		msg: "Deny permissions on a database scoped credential with cascade failed (action_id DWC; class_type DC)",
		src: "SQL Server",
	},
	"24346": {
		msg: "Revoke permissions on a database scoped credential succeeded (action_id R; class_type DC)",
		src: "SQL Server",
	},
	"24347": {
		msg: "Revoke permissions on a database scoped credential failed (action_id R; class_type DC)",
		src: "SQL Server",
	},
	"24348": {
		msg: "Revoke permissions with cascade on a database scoped credential succeeded (action_id RWC; class_type DC)",
		src: "SQL Server",
	},
	"24349": {
		msg: "Issued a change assembly owner command (action_id TO class_type AS)",
		src: "SQL Server",
	},
	"24350": {
		msg: "Revoke permissions with cascade on a database scoped credential failed (action_id RWC; class_type DC)",
		src: "SQL Server",
	},
	"24351": {
		msg: "Revoke permissions with grant on a database scoped credential succeeded (action_id RWG; class_type DC)",
		src: "SQL Server",
	},
	"24352": {
		msg: "Revoke permissions with grant on a database scoped credential failed (action_id RWG; class_type DC)",
		src: "SQL Server",
	},
	"24353": {
		msg: "Issued a change database scoped credential owner command (action_id TO; class_type DC)",
		src: "SQL Server",
	},
	"24354": {
		msg: "Issued a create external library command (action_id CR; class_type EL)",
		src: "SQL Server",
	},
	"24355": {
		msg: "Issued a change external library command (action_id AL; class_type EL)",
		src: "SQL Server",
	},
	"24356": {
		msg: "Issued a drop external library command (action_id DR; class_type EL)",
		src: "SQL Server",
	},
	"24357": {
		msg: "Grant permissions on an external library succeeded (action_id G; class_type EL)",
		src: "SQL Server",
	},
	"24358": {
		msg: "Grant permissions on an external library failed (action_id G; class_type EL)",
		src: "SQL Server",
	},
	"24359": {
		msg: "Grant permissions on an external library with grant succeeded (action_id GWG; class_type EL)",
		src: "SQL Server",
	},
	"24360": {
		msg: "Grant permissions on an external library with grant failed (action_id GWG; class_type EL)",
		src: "SQL Server",
	},
	"24361": {
		msg: "Deny permissions on an external library succeeded (action_id D; class_type EL)",
		src: "SQL Server",
	},
	"24362": {
		msg: "Deny permissions on an external library failed (action_id D; class_type EL)",
		src: "SQL Server",
	},
	"24363": {
		msg: "Deny permissions on an external library with cascade succeeded (action_id DWC; class_type EL)",
		src: "SQL Server",
	},
	"24364": {
		msg: "Deny permissions on an external library with cascade failed (action_id DWC; class_type EL)",
		src: "SQL Server",
	},
	"24365": {
		msg: "Revoke permissions on an external library succeeded (action_id R; class_type EL)",
		src: "SQL Server",
	},
	"24366": {
		msg: "Revoke permissions on an external library failed (action_id R; class_type EL)",
		src: "SQL Server",
	},
	"24367": {
		msg: "Revoke permissions with cascade on an external library succeeded (action_id RWC; class_type EL)",
		src: "SQL Server",
	},
	"24368": {
		msg: "Revoke permissions with cascade on an external library failed (action_id RWC; class_type EL)",
		src: "SQL Server",
	},
	"24369": {
		msg: "Revoke permissions with grant on an external library succeeded (action_id RWG; class_type EL)",
		src: "SQL Server",
	},
	"24370": {
		msg: "Revoke permissions with grant on an external library failed (action_id RWG; class_type EL)",
		src: "SQL Server",
	},
	"24371": {
		msg: "Issued a create database scoped reOrig governor command (action_id CR; class_type DR)",
		src: "SQL Server",
	},
	"24372": {
		msg: "Issued a change database scoped reOrig governor command (action_id AL; class_type DR)",
		src: "SQL Server",
	},
	"24373": {
		msg: "Issued a drop database scoped reOrig governor command (action_id DR; class_type DR)",
		src: "SQL Server",
	},
	"24374": {
		msg: "Issued a database bulk administration command (action_id DABO; class_type DB)",
		src: "SQL Server",
	},
	"24375": {
		msg: "Command to change permission failed (action_id D, DWC, G, GWG, R, RWC, RWG; class_type DC, EL)",
		src: "SQL Server",
	},
	"25000": {
		msg: "Undocumented Exchange mailbox operation",
		src: "Exchange",
	},
	"25001": {
		msg: "Operation Copy - Copy item to another Exchange mailbox folder",
		src: "Exchange",
	},
	"25002": {
		msg: "Operation Create - Create item in Exchange mailbox",
		src: "Exchange",
	},
	"25003": {
		msg: "Operation FolderBind - Access Exchange mailbox folder",
		src: "Exchange",
	},
	"25004": {
		msg: "Operation HardDelete - Delete Exchange mailbox item permanently from Recoverable Items folder",
		src: "Exchange",
	},
	"25005": {
		msg: "Operation MessageBind - Access Exchange mailbox item",
		src: "Exchange",
	},
	"25006": {
		msg: "Operation Move - Move item to another Exchange mailbox folder",
		src: "Exchange",
	},
	"25007": {
		msg: "Operation MoveToDeletedItems - Move Exchange mailbox item to Deleted Items folder",
		src: "Exchange",
	},
	"25008": {
		msg: "Operation SendAs - Send message using Send As Exchange mailbox permissions",
		src: "Exchange",
	},
	"25009": {
		msg: "Operation SendOnBehalf - Send message using Send on Behalf Exchange mailbox permissions",
		src: "Exchange",
	},
	"25010": {
		msg: "Operation SoftDelete - Delete Exchange mailbox item from Deleted Items folder",
		src: "Exchange",
	},
	"25011": {
		msg: "Operation Update - Update Exchange mailbox item's properties",
		src: "Exchange",
	},
	"25012": {
		msg: "Information Event - Mailbox audit policy applied",
		src: "Exchange",
	},
	"25100": {
		msg: "Undocumented Exchange admin operation",
		src: "Exchange",
	},
	"25101": {
		msg: "Add-ADPermission Exchange cmdlet issued",
		src: "Exchange",
	},
	"25102": {
		msg: "Add-AvailabilityAddressSpace Exchange cmdlet issued",
		src: "Exchange",
	},
	"25103": {
		msg: "Add-ContentFilterPhrase Exchange cmdlet issued",
		src: "Exchange",
	},
	"25104": {
		msg: "Add-DatabaseAvailabilityGroupServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25105": {
		msg: "Add-DistributionGroupMember Exchange cmdlet issued",
		src: "Exchange",
	},
	"25106": {
		msg: "Add-FederatedDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25107": {
		msg: "Add-IPAllowListEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25108": {
		msg: "Add-IPAllowListProvider Exchange cmdlet issued",
		src: "Exchange",
	},
	"25109": {
		msg: "Add-IPBlockListEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25110": {
		msg: "Add-IPBlockListProvider Exchange cmdlet issued",
		src: "Exchange",
	},
	"25111": {
		msg: "Add-MailboxDatabaseCopy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25112": {
		msg: "Add-MailboxFolderPermission Exchange cmdlet issued",
		src: "Exchange",
	},
	"25113": {
		msg: "Add-MailboxPermission Exchange cmdlet issued",
		src: "Exchange",
	},
	"25114": {
		msg: "Add-ManagementRoleEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25115": {
		msg: "Add-PublicFolderAdministrativePermission Exchange cmdlet issued",
		src: "Exchange",
	},
	"25116": {
		msg: "Add-PublicFolderClientPermission Exchange cmdlet issued",
		src: "Exchange",
	},
	"25117": {
		msg: "Add-RoleGroupMember Exchange cmdlet issued",
		src: "Exchange",
	},
	"25118": {
		msg: "Clean-MailboxDatabase Exchange cmdlet issued",
		src: "Exchange",
	},
	"25119": {
		msg: "Clear-ActiveSyncDevice Exchange cmdlet issued",
		src: "Exchange",
	},
	"25120": {
		msg: "Clear-TextMessagingAccount Exchange cmdlet issued",
		src: "Exchange",
	},
	"25121": {
		msg: "Compare-TextMessagingVerificationCode Exchange cmdlet issued",
		src: "Exchange",
	},
	"25122": {
		msg: "Connect-Mailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25123": {
		msg: "Disable-AddressListPaging Exchange cmdlet issued",
		src: "Exchange",
	},
	"25124": {
		msg: "Disable-CmdletExtensionAgent Exchange cmdlet issued",
		src: "Exchange",
	},
	"25125": {
		msg: "Disable-DistributionGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25126": {
		msg: "Disable-InboxRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25127": {
		msg: "Disable-JournalRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25128": {
		msg: "Disable-Mailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25129": {
		msg: "Disable-MailContact Exchange cmdlet issued",
		src: "Exchange",
	},
	"25130": {
		msg: "Disable-MailPublicFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25131": {
		msg: "Disable-MailUser Exchange cmdlet issued",
		src: "Exchange",
	},
	"25132": {
		msg: "Disable-OutlookAnywhere Exchange cmdlet issued",
		src: "Exchange",
	},
	"25133": {
		msg: "Disable-OutlookProtectionRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25134": {
		msg: "Disable-RemoteMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25135": {
		msg: "Disable-ServiceEmailChannel Exchange cmdlet issued",
		src: "Exchange",
	},
	"25136": {
		msg: "Disable-TransportAgent Exchange cmdlet issued",
		src: "Exchange",
	},
	"25137": {
		msg: "Disable-TransportRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25138": {
		msg: "Disable-UMAutoAttendant Exchange cmdlet issued",
		src: "Exchange",
	},
	"25139": {
		msg: "Disable-UMIPGateway Exchange cmdlet issued",
		src: "Exchange",
	},
	"25140": {
		msg: "Disable-UMMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25141": {
		msg: "Disable-UMServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25142": {
		msg: "Dismount-Database Exchange cmdlet issued",
		src: "Exchange",
	},
	"25143": {
		msg: "Enable-AddressListPaging Exchange cmdlet issued",
		src: "Exchange",
	},
	"25144": {
		msg: "Enable-AntispamUpdates Exchange cmdlet issued",
		src: "Exchange",
	},
	"25145": {
		msg: "Enable-CmdletExtensionAgent Exchange cmdlet issued",
		src: "Exchange",
	},
	"25146": {
		msg: "Enable-DistributionGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25147": {
		msg: "Enable-ExchangeCertificate Exchange cmdlet issued",
		src: "Exchange",
	},
	"25148": {
		msg: "Enable-InboxRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25149": {
		msg: "Enable-JournalRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25150": {
		msg: "Enable-Mailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25151": {
		msg: "Enable-MailContact Exchange cmdlet issued",
		src: "Exchange",
	},
	"25152": {
		msg: "Enable-MailPublicFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25153": {
		msg: "Enable-MailUser Exchange cmdlet issued",
		src: "Exchange",
	},
	"25154": {
		msg: "Enable-OutlookAnywhere Exchange cmdlet issued",
		src: "Exchange",
	},
	"25155": {
		msg: "Enable-OutlookProtectionRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25156": {
		msg: "Enable-RemoteMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25157": {
		msg: "Enable-ServiceEmailChannel Exchange cmdlet issued",
		src: "Exchange",
	},
	"25158": {
		msg: "Enable-TransportAgent Exchange cmdlet issued",
		src: "Exchange",
	},
	"25159": {
		msg: "Enable-TransportRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25160": {
		msg: "Enable-UMAutoAttendant Exchange cmdlet issued",
		src: "Exchange",
	},
	"25161": {
		msg: "Enable-UMIPGateway Exchange cmdlet issued",
		src: "Exchange",
	},
	"25162": {
		msg: "Enable-UMMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25163": {
		msg: "Enable-UMServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25164": {
		msg: "Export-ActiveSyncLog Exchange cmdlet issued",
		src: "Exchange",
	},
	"25165": {
		msg: "Export-AutoDiscoverConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25166": {
		msg: "Export-ExchangeCertificate Exchange cmdlet issued",
		src: "Exchange",
	},
	"25167": {
		msg: "Export-JournalRuleCollection Exchange cmdlet issued",
		src: "Exchange",
	},
	"25168": {
		msg: "Export-MailboxDiagnosticLogs Exchange cmdlet issued",
		src: "Exchange",
	},
	"25169": {
		msg: "Export-Message Exchange cmdlet issued",
		src: "Exchange",
	},
	"25170": {
		msg: "Export-RecipientDataProperty Exchange cmdlet issued",
		src: "Exchange",
	},
	"25171": {
		msg: "Export-TransportRuleCollection Exchange cmdlet issued",
		src: "Exchange",
	},
	"25172": {
		msg: "Export-UMCallDataRecord Exchange cmdlet issued",
		src: "Exchange",
	},
	"25173": {
		msg: "Export-UMPrompt Exchange cmdlet issued",
		src: "Exchange",
	},
	"25174": {
		msg: "Import-ExchangeCertificate Exchange cmdlet issued",
		src: "Exchange",
	},
	"25175": {
		msg: "Import-JournalRuleCollection Exchange cmdlet issued",
		src: "Exchange",
	},
	"25176": {
		msg: "Import-RecipientDataProperty Exchange cmdlet issued",
		src: "Exchange",
	},
	"25177": {
		msg: "Import-TransportRuleCollection Exchange cmdlet issued",
		src: "Exchange",
	},
	"25178": {
		msg: "Import-UMPrompt Exchange cmdlet issued",
		src: "Exchange",
	},
	"25179": {
		msg: "Install-TransportAgent Exchange cmdlet issued",
		src: "Exchange",
	},
	"25180": {
		msg: "Mount-Database Exchange cmdlet issued",
		src: "Exchange",
	},
	"25181": {
		msg: "Move-ActiveMailboxDatabase Exchange cmdlet issued",
		src: "Exchange",
	},
	"25182": {
		msg: "Move-AddressList Exchange cmdlet issued",
		src: "Exchange",
	},
	"25183": {
		msg: "Move-DatabasePath Exchange cmdlet issued",
		src: "Exchange",
	},
	"25184": {
		msg: "Move-OfflineAddressBook Exchange cmdlet issued",
		src: "Exchange",
	},
	"25185": {
		msg: "New-AcceptedDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25186": {
		msg: "New-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25187": {
		msg: "New-ActiveSyncMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25188": {
		msg: "New-ActiveSyncVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25189": {
		msg: "New-AddressList Exchange cmdlet issued",
		src: "Exchange",
	},
	"25190": {
		msg: "New-AdminAuditLogSearch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25191": {
		msg: "New-AutodiscoverVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25192": {
		msg: "New-AvailabilityReportOutage Exchange cmdlet issued",
		src: "Exchange",
	},
	"25193": {
		msg: "New-ClientAccessArray Exchange cmdlet issued",
		src: "Exchange",
	},
	"25194": {
		msg: "New-DatabaseAvailabilityGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25195": {
		msg: "New-DatabaseAvailabilityGroupNetwork Exchange cmdlet issued",
		src: "Exchange",
	},
	"25196": {
		msg: "New-DeliveryAgentConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25197": {
		msg: "New-DistributionGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25198": {
		msg: "New-DynamicDistributionGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25199": {
		msg: "New-EcpVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25200": {
		msg: "New-EdgeSubscription Exchange cmdlet issued",
		src: "Exchange",
	},
	"25201": {
		msg: "New-EdgeSyncServiceConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25202": {
		msg: "New-EmailAddressPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25203": {
		msg: "New-ExchangeCertificate Exchange cmdlet issued",
		src: "Exchange",
	},
	"25204": {
		msg: "New-FederationTrust Exchange cmdlet issued",
		src: "Exchange",
	},
	"25205": {
		msg: "New-ForeignConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25206": {
		msg: "New-GlobalAddressList Exchange cmdlet issued",
		src: "Exchange",
	},
	"25207": {
		msg: "New-InboxRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25208": {
		msg: "New-JournalRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25209": {
		msg: "New-Mailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25210": {
		msg: "New-MailboxAuditLogSearch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25211": {
		msg: "New-MailboxDatabase Exchange cmdlet issued",
		src: "Exchange",
	},
	"25212": {
		msg: "New-MailboxFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25213": {
		msg: "New-MailboxRepairRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25214": {
		msg: "New-MailboxRestoreRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25215": {
		msg: "New-MailContact Exchange cmdlet issued",
		src: "Exchange",
	},
	"25216": {
		msg: "New-MailMessage Exchange cmdlet issued",
		src: "Exchange",
	},
	"25217": {
		msg: "New-MailUser Exchange cmdlet issued",
		src: "Exchange",
	},
	"25218": {
		msg: "New-ManagedContentSettings Exchange cmdlet issued",
		src: "Exchange",
	},
	"25219": {
		msg: "New-ManagedFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25220": {
		msg: "New-ManagedFolderMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25221": {
		msg: "New-ManagementRole Exchange cmdlet issued",
		src: "Exchange",
	},
	"25222": {
		msg: "New-ManagementRoleAssignment Exchange cmdlet issued",
		src: "Exchange",
	},
	"25223": {
		msg: "New-ManagementScope Exchange cmdlet issued",
		src: "Exchange",
	},
	"25224": {
		msg: "New-MessageClassification Exchange cmdlet issued",
		src: "Exchange",
	},
	"25225": {
		msg: "New-MoveRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25226": {
		msg: "New-OabVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25227": {
		msg: "New-OfflineAddressBook Exchange cmdlet issued",
		src: "Exchange",
	},
	"25228": {
		msg: "New-OrganizationRelationship Exchange cmdlet issued",
		src: "Exchange",
	},
	"25229": {
		msg: "New-OutlookProtectionRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25230": {
		msg: "New-OutlookProvider Exchange cmdlet issued",
		src: "Exchange",
	},
	"25231": {
		msg: "New-OwaMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25232": {
		msg: "New-OwaVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25233": {
		msg: "New-PublicFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25234": {
		msg: "New-PublicFolderDatabase Exchange cmdlet issued",
		src: "Exchange",
	},
	"25235": {
		msg: "New-PublicFolderDatabaseRepairRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25236": {
		msg: "New-ReceiveConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25237": {
		msg: "New-RemoteDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25238": {
		msg: "New-RemoteMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25239": {
		msg: "New-RetentionPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25240": {
		msg: "New-RetentionPolicyTag Exchange cmdlet issued",
		src: "Exchange",
	},
	"25241": {
		msg: "New-RoleAssignmentPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25242": {
		msg: "New-RoleGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25243": {
		msg: "New-RoutingGroupConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25244": {
		msg: "New-RpcClientAccess Exchange cmdlet issued",
		src: "Exchange",
	},
	"25245": {
		msg: "New-SendConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25246": {
		msg: "New-SharingPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25247": {
		msg: "New-SystemMessage Exchange cmdlet issued",
		src: "Exchange",
	},
	"25248": {
		msg: "New-ThrottlingPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25249": {
		msg: "New-TransportRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25250": {
		msg: "New-UMAutoAttendant Exchange cmdlet issued",
		src: "Exchange",
	},
	"25251": {
		msg: "New-UMDialPlan Exchange cmdlet issued",
		src: "Exchange",
	},
	"25252": {
		msg: "New-UMHuntGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25253": {
		msg: "New-UMIPGateway Exchange cmdlet issued",
		src: "Exchange",
	},
	"25254": {
		msg: "New-UMMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25255": {
		msg: "New-WebServicesVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25256": {
		msg: "New-X400AuthoritativeDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25257": {
		msg: "Remove-AcceptedDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25258": {
		msg: "Remove-ActiveSyncDevice Exchange cmdlet issued",
		src: "Exchange",
	},
	"25259": {
		msg: "Remove-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25260": {
		msg: "Remove-ActiveSyncDeviceClass Exchange cmdlet issued",
		src: "Exchange",
	},
	"25261": {
		msg: "Remove-ActiveSyncMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25262": {
		msg: "Remove-ActiveSyncVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25263": {
		msg: "Remove-AddressList Exchange cmdlet issued",
		src: "Exchange",
	},
	"25264": {
		msg: "Remove-ADPermission Exchange cmdlet issued",
		src: "Exchange",
	},
	"25265": {
		msg: "Remove-AutodiscoverVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25266": {
		msg: "Remove-AvailabilityAddressSpace Exchange cmdlet issued",
		src: "Exchange",
	},
	"25267": {
		msg: "Remove-AvailabilityReportOutage Exchange cmdlet issued",
		src: "Exchange",
	},
	"25268": {
		msg: "Remove-ClientAccessArray Exchange cmdlet issued",
		src: "Exchange",
	},
	"25269": {
		msg: "Remove-ContentFilterPhrase Exchange cmdlet issued",
		src: "Exchange",
	},
	"25270": {
		msg: "Remove-DatabaseAvailabilityGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25271": {
		msg: "Remove-DatabaseAvailabilityGroupNetwork Exchange cmdlet issued",
		src: "Exchange",
	},
	"25272": {
		msg: "Remove-DatabaseAvailabilityGroupServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25273": {
		msg: "Remove-DeliveryAgentConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25274": {
		msg: "Remove-DistributionGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25275": {
		msg: "Remove-DistributionGroupMember Exchange cmdlet issued",
		src: "Exchange",
	},
	"25276": {
		msg: "Remove-DynamicDistributionGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25277": {
		msg: "Remove-EcpVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25278": {
		msg: "Remove-EdgeSubscription Exchange cmdlet issued",
		src: "Exchange",
	},
	"25279": {
		msg: "Remove-EmailAddressPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25280": {
		msg: "Remove-ExchangeCertificate Exchange cmdlet issued",
		src: "Exchange",
	},
	"25281": {
		msg: "Remove-FederatedDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25282": {
		msg: "Remove-FederationTrust Exchange cmdlet issued",
		src: "Exchange",
	},
	"25283": {
		msg: "Remove-ForeignConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25284": {
		msg: "Remove-GlobalAddressList Exchange cmdlet issued",
		src: "Exchange",
	},
	"25285": {
		msg: "Remove-InboxRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25286": {
		msg: "Remove-IPAllowListEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25287": {
		msg: "Remove-IPAllowListProvider Exchange cmdlet issued",
		src: "Exchange",
	},
	"25288": {
		msg: "Remove-IPBlockListEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25289": {
		msg: "Remove-IPBlockListProvider Exchange cmdlet issued",
		src: "Exchange",
	},
	"25290": {
		msg: "Remove-JournalRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25291": {
		msg: "Remove-Mailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25292": {
		msg: "Remove-MailboxDatabase Exchange cmdlet issued",
		src: "Exchange",
	},
	"25293": {
		msg: "Remove-MailboxDatabaseCopy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25294": {
		msg: "Remove-MailboxFolderPermission Exchange cmdlet issued",
		src: "Exchange",
	},
	"25295": {
		msg: "Remove-MailboxPermission Exchange cmdlet issued",
		src: "Exchange",
	},
	"25296": {
		msg: "Remove-MailboxRestoreRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25297": {
		msg: "Remove-MailContact Exchange cmdlet issued",
		src: "Exchange",
	},
	"25298": {
		msg: "Remove-MailUser Exchange cmdlet issued",
		src: "Exchange",
	},
	"25299": {
		msg: "Remove-ManagedContentSettings Exchange cmdlet issued",
		src: "Exchange",
	},
	"25300": {
		msg: "Remove-ManagedFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25301": {
		msg: "Remove-ManagedFolderMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25302": {
		msg: "Remove-ManagementRole Exchange cmdlet issued",
		src: "Exchange",
	},
	"25303": {
		msg: "Remove-ManagementRoleAssignment Exchange cmdlet issued",
		src: "Exchange",
	},
	"25304": {
		msg: "Remove-ManagementRoleEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25305": {
		msg: "Remove-ManagementScope Exchange cmdlet issued",
		src: "Exchange",
	},
	"25306": {
		msg: "Remove-Message Exchange cmdlet issued",
		src: "Exchange",
	},
	"25307": {
		msg: "Remove-MessageClassification Exchange cmdlet issued",
		src: "Exchange",
	},
	"25308": {
		msg: "Remove-MoveRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25309": {
		msg: "Remove-OabVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25310": {
		msg: "Remove-OfflineAddressBook Exchange cmdlet issued",
		src: "Exchange",
	},
	"25311": {
		msg: "Remove-OrganizationRelationship Exchange cmdlet issued",
		src: "Exchange",
	},
	"25312": {
		msg: "Remove-OutlookProtectionRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25313": {
		msg: "Remove-OutlookProvider Exchange cmdlet issued",
		src: "Exchange",
	},
	"25314": {
		msg: "Remove-OwaMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25315": {
		msg: "Remove-OwaVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25316": {
		msg: "Remove-PublicFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25317": {
		msg: "Remove-PublicFolderAdministrativePermission Exchange cmdlet issued",
		src: "Exchange",
	},
	"25318": {
		msg: "Remove-PublicFolderClientPermission Exchange cmdlet issued",
		src: "Exchange",
	},
	"25319": {
		msg: "Remove-PublicFolderDatabase Exchange cmdlet issued",
		src: "Exchange",
	},
	"25320": {
		msg: "Remove-ReceiveConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25321": {
		msg: "Remove-RemoteDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25322": {
		msg: "Remove-RemoteMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25323": {
		msg: "Remove-RetentionPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25324": {
		msg: "Remove-RetentionPolicyTag Exchange cmdlet issued",
		src: "Exchange",
	},
	"25325": {
		msg: "Remove-RoleAssignmentPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25326": {
		msg: "Remove-RoleGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25327": {
		msg: "Remove-RoleGroupMember Exchange cmdlet issued",
		src: "Exchange",
	},
	"25328": {
		msg: "Remove-RoutingGroupConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25329": {
		msg: "Remove-RpcClientAccess Exchange cmdlet issued",
		src: "Exchange",
	},
	"25330": {
		msg: "Remove-SendConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25331": {
		msg: "Remove-SharingPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25332": {
		msg: "Remove-StoreMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25333": {
		msg: "Remove-SystemMessage Exchange cmdlet issued",
		src: "Exchange",
	},
	"25334": {
		msg: "Remove-ThrottlingPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25335": {
		msg: "Remove-TransportRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25336": {
		msg: "Remove-UMAutoAttendant Exchange cmdlet issued",
		src: "Exchange",
	},
	"25337": {
		msg: "Remove-UMDialPlan Exchange cmdlet issued",
		src: "Exchange",
	},
	"25338": {
		msg: "Remove-UMHuntGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25339": {
		msg: "Remove-UMIPGateway Exchange cmdlet issued",
		src: "Exchange",
	},
	"25340": {
		msg: "Remove-UMMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25341": {
		msg: "Remove-WebServicesVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25342": {
		msg: "Remove-X400AuthoritativeDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25343": {
		msg: "Restore-DatabaseAvailabilityGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25344": {
		msg: "Restore-DetailsTemplate Exchange cmdlet issued",
		src: "Exchange",
	},
	"25345": {
		msg: "Restore-Mailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25346": {
		msg: "Resume-MailboxDatabaseCopy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25347": {
		msg: "Resume-MailboxExportRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25348": {
		msg: "Resume-MailboxRestoreRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25349": {
		msg: "Resume-Message Exchange cmdlet issued",
		src: "Exchange",
	},
	"25350": {
		msg: "Resume-MoveRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25351": {
		msg: "Resume-PublicFolderReplication Exchange cmdlet issued",
		src: "Exchange",
	},
	"25352": {
		msg: "Resume-Queue Exchange cmdlet issued",
		src: "Exchange",
	},
	"25353": {
		msg: "Retry-Queue Exchange cmdlet issued",
		src: "Exchange",
	},
	"25354": {
		msg: "Send-TextMessagingVerificationCode Exchange cmdlet issued",
		src: "Exchange",
	},
	"25355": {
		msg: "Set-AcceptedDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25356": {
		msg: "Set-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25357": {
		msg: "Set-ActiveSyncMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25358": {
		msg: "Set-ActiveSyncOrganizationSettings Exchange cmdlet issued",
		src: "Exchange",
	},
	"25359": {
		msg: "Set-ActiveSyncVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25360": {
		msg: "Set-AddressList Exchange cmdlet issued",
		src: "Exchange",
	},
	"25361": {
		msg: "Set-AdminAuditLogConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25362": {
		msg: "Set-ADServerSettings Exchange cmdlet issued",
		src: "Exchange",
	},
	"25363": {
		msg: "Set-ADSite Exchange cmdlet issued",
		src: "Exchange",
	},
	"25364": {
		msg: "Set-AdSiteLink Exchange cmdlet issued",
		src: "Exchange",
	},
	"25365": {
		msg: "Set-AutodiscoverVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25366": {
		msg: "Set-AvailabilityConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25367": {
		msg: "Set-AvailabilityReportOutage Exchange cmdlet issued",
		src: "Exchange",
	},
	"25368": {
		msg: "Set-CalendarNotification Exchange cmdlet issued",
		src: "Exchange",
	},
	"25369": {
		msg: "Set-CalendarProcessing Exchange cmdlet issued",
		src: "Exchange",
	},
	"25370": {
		msg: "Set-CASMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25371": {
		msg: "Set-ClientAccessArray Exchange cmdlet issued",
		src: "Exchange",
	},
	"25372": {
		msg: "Set-ClientAccessServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25373": {
		msg: "Set-CmdletExtensionAgent Exchange cmdlet issued",
		src: "Exchange",
	},
	"25374": {
		msg: "Set-Contact Exchange cmdlet issued",
		src: "Exchange",
	},
	"25375": {
		msg: "Set-ContentFilterConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25376": {
		msg: "Set-DatabaseAvailabilityGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25377": {
		msg: "Set-DatabaseAvailabilityGroupNetwork Exchange cmdlet issued",
		src: "Exchange",
	},
	"25378": {
		msg: "Set-DeliveryAgentConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25379": {
		msg: "Set-DetailsTemplate Exchange cmdlet issued",
		src: "Exchange",
	},
	"25380": {
		msg: "Set-DistributionGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25381": {
		msg: "Set-DynamicDistributionGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25382": {
		msg: "Set-EcpVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25383": {
		msg: "Set-EdgeSyncServiceConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25384": {
		msg: "Set-EmailAddressPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25385": {
		msg: "Set-EventLogLevel Exchange cmdlet issued",
		src: "Exchange",
	},
	"25386": {
		msg: "Set-ExchangeAssistanceConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25387": {
		msg: "Set-ExchangeServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25388": {
		msg: "Set-FederatedOrganizationIdentifier Exchange cmdlet issued",
		src: "Exchange",
	},
	"25389": {
		msg: "Set-FederationTrust Exchange cmdlet issued",
		src: "Exchange",
	},
	"25390": {
		msg: "Set-ForeignConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25391": {
		msg: "Set-GlobalAddressList Exchange cmdlet issued",
		src: "Exchange",
	},
	"25392": {
		msg: "Set-Group Exchange cmdlet issued",
		src: "Exchange",
	},
	"25393": {
		msg: "Set-ImapSettings Exchange cmdlet issued",
		src: "Exchange",
	},
	"25394": {
		msg: "Set-InboxRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25395": {
		msg: "Set-IPAllowListConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25396": {
		msg: "Set-IPAllowListProvider Exchange cmdlet issued",
		src: "Exchange",
	},
	"25397": {
		msg: "Set-IPAllowListProvidersConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25398": {
		msg: "Set-IPBlockListConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25399": {
		msg: "Set-IPBlockListProvider Exchange cmdlet issued",
		src: "Exchange",
	},
	"25400": {
		msg: "Set-IPBlockListProvidersConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25401": {
		msg: "Set-IRMConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25402": {
		msg: "Set-JournalRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25403": {
		msg: "Set-Mailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25404": {
		msg: "Set-MailboxAuditBypassAssociation Exchange cmdlet issued",
		src: "Exchange",
	},
	"25405": {
		msg: "Set-MailboxAutoReplyConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25406": {
		msg: "Set-MailboxCalendarConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25407": {
		msg: "Set-MailboxCalendarFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25408": {
		msg: "Set-MailboxDatabase Exchange cmdlet issued",
		src: "Exchange",
	},
	"25409": {
		msg: "Set-MailboxDatabaseCopy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25410": {
		msg: "Set-MailboxFolderPermission Exchange cmdlet issued",
		src: "Exchange",
	},
	"25411": {
		msg: "Set-MailboxJunkEmailConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25412": {
		msg: "Set-MailboxMessageConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25413": {
		msg: "Set-MailboxRegionalConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25414": {
		msg: "Set-MailboxRestoreRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25415": {
		msg: "Set-MailboxServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25416": {
		msg: "Set-MailboxSpellingConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25417": {
		msg: "Set-MailContact Exchange cmdlet issued",
		src: "Exchange",
	},
	"25418": {
		msg: "Set-MailPublicFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25419": {
		msg: "Set-MailUser Exchange cmdlet issued",
		src: "Exchange",
	},
	"25420": {
		msg: "Set-ManagedContentSettings Exchange cmdlet issued",
		src: "Exchange",
	},
	"25421": {
		msg: "Set-ManagedFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25422": {
		msg: "Set-ManagedFolderMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25423": {
		msg: "Set-ManagementRoleAssignment Exchange cmdlet issued",
		src: "Exchange",
	},
	"25424": {
		msg: "Set-ManagementRoleEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25425": {
		msg: "Set-ManagementScope Exchange cmdlet issued",
		src: "Exchange",
	},
	"25426": {
		msg: "Set-MessageClassification Exchange cmdlet issued",
		src: "Exchange",
	},
	"25427": {
		msg: "Set-MoveRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25428": {
		msg: "Set-OabVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25429": {
		msg: "Set-OfflineAddressBook Exchange cmdlet issued",
		src: "Exchange",
	},
	"25430": {
		msg: "Set-OrganizationConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25431": {
		msg: "Set-OrganizationRelationship Exchange cmdlet issued",
		src: "Exchange",
	},
	"25432": {
		msg: "Set-OutlookAnywhere Exchange cmdlet issued",
		src: "Exchange",
	},
	"25433": {
		msg: "Set-OutlookProtectionRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25434": {
		msg: "Set-OutlookProvider Exchange cmdlet issued",
		src: "Exchange",
	},
	"25435": {
		msg: "Set-OwaMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25436": {
		msg: "Set-OwaVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25437": {
		msg: "Set-PopSettings Exchange cmdlet issued",
		src: "Exchange",
	},
	"25438": {
		msg: "Set-PowerShellVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25439": {
		msg: "Set-PublicFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25440": {
		msg: "Set-PublicFolderDatabase Exchange cmdlet issued",
		src: "Exchange",
	},
	"25441": {
		msg: "Set-ReceiveConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25442": {
		msg: "Set-RecipientFilterConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25443": {
		msg: "Set-RemoteDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25444": {
		msg: "Set-RemoteMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25445": {
		msg: "Set-ReOrigConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25446": {
		msg: "Set-RetentionPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25447": {
		msg: "Set-RetentionPolicyTag Exchange cmdlet issued",
		src: "Exchange",
	},
	"25448": {
		msg: "Set-RoleAssignmentPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25449": {
		msg: "Set-RoleGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25450": {
		msg: "Set-RoutingGroupConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25451": {
		msg: "Set-RpcClientAccess Exchange cmdlet issued",
		src: "Exchange",
	},
	"25452": {
		msg: "Set-SendConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25453": {
		msg: "Set-SenderFilterConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25454": {
		msg: "Set-SenderIdConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25455": {
		msg: "Set-SenderReputationConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25456": {
		msg: "Set-SharingPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25457": {
		msg: "Set-SystemMessage Exchange cmdlet issued",
		src: "Exchange",
	},
	"25458": {
		msg: "Set-TextMessagingAccount Exchange cmdlet issued",
		src: "Exchange",
	},
	"25459": {
		msg: "Set-ThrottlingPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25460": {
		msg: "Set-ThrottlingPolicyAssociation Exchange cmdlet issued",
		src: "Exchange",
	},
	"25461": {
		msg: "Set-TransportAgent Exchange cmdlet issued",
		src: "Exchange",
	},
	"25462": {
		msg: "Set-TransportConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25463": {
		msg: "Set-TransportRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25464": {
		msg: "Set-TransportServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25465": {
		msg: "Set-UMAutoAttendant Exchange cmdlet issued",
		src: "Exchange",
	},
	"25466": {
		msg: "Set-UMDialPlan Exchange cmdlet issued",
		src: "Exchange",
	},
	"25467": {
		msg: "Set-UMIPGateway Exchange cmdlet issued",
		src: "Exchange",
	},
	"25468": {
		msg: "Set-UMMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25469": {
		msg: "Set-UMMailboxPIN Exchange cmdlet issued",
		src: "Exchange",
	},
	"25470": {
		msg: "Set-UMMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25471": {
		msg: "Set-UmServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25472": {
		msg: "Set-User Exchange cmdlet issued",
		src: "Exchange",
	},
	"25473": {
		msg: "Set-WebServicesVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25474": {
		msg: "Set-X400AuthoritativeDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25475": {
		msg: "Start-DatabaseAvailabilityGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25476": {
		msg: "Start-EdgeSynchronization Exchange cmdlet issued",
		src: "Exchange",
	},
	"25477": {
		msg: "Start-ManagedFolderAssistant Exchange cmdlet issued",
		src: "Exchange",
	},
	"25478": {
		msg: "Start-RetentionAutoTagLearning Exchange cmdlet issued",
		src: "Exchange",
	},
	"25479": {
		msg: "Stop-DatabaseAvailabilityGroup Exchange cmdlet issued",
		src: "Exchange",
	},
	"25480": {
		msg: "Stop-ManagedFolderAssistant Exchange cmdlet issued",
		src: "Exchange",
	},
	"25481": {
		msg: "Suspend-MailboxDatabaseCopy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25482": {
		msg: "Suspend-MailboxRestoreRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25483": {
		msg: "Suspend-Message Exchange cmdlet issued",
		src: "Exchange",
	},
	"25484": {
		msg: "Suspend-MoveRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25485": {
		msg: "Suspend-PublicFolderReplication Exchange cmdlet issued",
		src: "Exchange",
	},
	"25486": {
		msg: "Suspend-Queue Exchange cmdlet issued",
		src: "Exchange",
	},
	"25487": {
		msg: "Test-ActiveSyncConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25488": {
		msg: "Test-AssistantHealth Exchange cmdlet issued",
		src: "Exchange",
	},
	"25489": {
		msg: "Test-CalendarConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25490": {
		msg: "Test-EcpConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25491": {
		msg: "Test-EdgeSynchronization Exchange cmdlet issued",
		src: "Exchange",
	},
	"25492": {
		msg: "Test-ExchangeSearch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25493": {
		msg: "Test-FederationTrust Exchange cmdlet issued",
		src: "Exchange",
	},
	"25494": {
		msg: "Test-FederationTrustCertificate Exchange cmdlet issued",
		src: "Exchange",
	},
	"25495": {
		msg: "Test-ImapConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25496": {
		msg: "Test-IPAllowListProvider Exchange cmdlet issued",
		src: "Exchange",
	},
	"25497": {
		msg: "Test-IPBlockListProvider Exchange cmdlet issued",
		src: "Exchange",
	},
	"25498": {
		msg: "Test-IRMConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25499": {
		msg: "Test-Mailflow Exchange cmdlet issued",
		src: "Exchange",
	},
	"25500": {
		msg: "Test-MAPIConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25501": {
		msg: "Test-MRSHealth Exchange cmdlet issued",
		src: "Exchange",
	},
	"25502": {
		msg: "Test-OrganizationRelationship Exchange cmdlet issued",
		src: "Exchange",
	},
	"25503": {
		msg: "Test-OutlookConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25504": {
		msg: "Test-OutlookWebServices Exchange cmdlet issued",
		src: "Exchange",
	},
	"25505": {
		msg: "Test-OwaConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25506": {
		msg: "Test-PopConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25507": {
		msg: "Test-PowerShellConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25508": {
		msg: "Test-ReplicationHealth Exchange cmdlet issued",
		src: "Exchange",
	},
	"25509": {
		msg: "Test-SenderId Exchange cmdlet issued",
		src: "Exchange",
	},
	"25510": {
		msg: "Test-ServiceHealth Exchange cmdlet issued",
		src: "Exchange",
	},
	"25511": {
		msg: "Test-SmtpConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25512": {
		msg: "Test-SystemHealth Exchange cmdlet issued",
		src: "Exchange",
	},
	"25513": {
		msg: "Test-UMConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25514": {
		msg: "Test-WebServicesConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25515": {
		msg: "Uninstall-TransportAgent Exchange cmdlet issued",
		src: "Exchange",
	},
	"25516": {
		msg: "Update-AddressList Exchange cmdlet issued",
		src: "Exchange",
	},
	"25517": {
		msg: "Update-DistributionGroupMember Exchange cmdlet issued",
		src: "Exchange",
	},
	"25518": {
		msg: "Update-EmailAddressPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25519": {
		msg: "Update-FileDistributionService Exchange cmdlet issued",
		src: "Exchange",
	},
	"25520": {
		msg: "Update-GlobalAddressList Exchange cmdlet issued",
		src: "Exchange",
	},
	"25521": {
		msg: "Update-MailboxDatabaseCopy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25522": {
		msg: "Update-OfflineAddressBook Exchange cmdlet issued",
		src: "Exchange",
	},
	"25523": {
		msg: "Update-PublicFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25524": {
		msg: "Update-PublicFolderHierarchy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25525": {
		msg: "Update-Recipient Exchange cmdlet issued",
		src: "Exchange",
	},
	"25526": {
		msg: "Update-RoleGroupMember Exchange cmdlet issued",
		src: "Exchange",
	},
	"25527": {
		msg: "Update-SafeList Exchange cmdlet issued",
		src: "Exchange",
	},
	"25528": {
		msg: "Write-AdminAuditLog Exchange cmdlet issued",
		src: "Exchange",
	},
	"25529": {
		msg: "Add-GlobalMonitoringOverride Exchange cmdlet issued",
		src: "Exchange",
	},
	"25530": {
		msg: "Add-ResubmitRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25531": {
		msg: "Add-ServerMonitoringOverride Exchange cmdlet issued",
		src: "Exchange",
	},
	"25532": {
		msg: "Clear-MobileDevice Exchange cmdlet issued",
		src: "Exchange",
	},
	"25533": {
		msg: "Complete-MigrationBatch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25534": {
		msg: "Disable-App Exchange cmdlet issued",
		src: "Exchange",
	},
	"25535": {
		msg: "Disable-MailboxQuarantine Exchange cmdlet issued",
		src: "Exchange",
	},
	"25536": {
		msg: "Disable-UMCallAnsweringRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25537": {
		msg: "Disable-UMService Exchange cmdlet issued",
		src: "Exchange",
	},
	"25538": {
		msg: "Dump-ProvisioningCache Exchange cmdlet issued",
		src: "Exchange",
	},
	"25539": {
		msg: "Enable-App Exchange cmdlet issued",
		src: "Exchange",
	},
	"25540": {
		msg: "Enable-MailboxQuarantine Exchange cmdlet issued",
		src: "Exchange",
	},
	"25541": {
		msg: "Enable-UMCallAnsweringRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25542": {
		msg: "Enable-UMService Exchange cmdlet issued",
		src: "Exchange",
	},
	"25543": {
		msg: "Export-DlpPolicyCollection Exchange cmdlet issued",
		src: "Exchange",
	},
	"25544": {
		msg: "Export-MigrationReport Exchange cmdlet issued",
		src: "Exchange",
	},
	"25545": {
		msg: "Import-DlpPolicyCollection Exchange cmdlet issued",
		src: "Exchange",
	},
	"25546": {
		msg: "Import-DlpPolicyTemplate Exchange cmdlet issued",
		src: "Exchange",
	},
	"25547": {
		msg: "Invoke-MonitoringProbe Exchange cmdlet issued",
		src: "Exchange",
	},
	"25548": {
		msg: "New-AddressBookPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25549": {
		msg: "New-App Exchange cmdlet issued",
		src: "Exchange",
	},
	"25550": {
		msg: "New-AuthServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25551": {
		msg: "New-ClassificationRuleCollection Exchange cmdlet issued",
		src: "Exchange",
	},
	"25552": {
		msg: "New-DlpPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25553": {
		msg: "New-HybridConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25554": {
		msg: "New-MailboxExportRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25555": {
		msg: "New-MailboxImportRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25556": {
		msg: "New-MailboxSearch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25557": {
		msg: "New-MalwareFilterPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25558": {
		msg: "New-MigrationBatch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25559": {
		msg: "New-MigrationEndpoint Exchange cmdlet issued",
		src: "Exchange",
	},
	"25560": {
		msg: "New-MobileDeviceMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25561": {
		msg: "New-OnPremisesOrganization Exchange cmdlet issued",
		src: "Exchange",
	},
	"25562": {
		msg: "New-PartnerApplication Exchange cmdlet issued",
		src: "Exchange",
	},
	"25563": {
		msg: "New-PolicyTipConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25564": {
		msg: "New-PowerShellVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25565": {
		msg: "New-PublicFolderMigrationRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25566": {
		msg: "New-ReOrigPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25567": {
		msg: "New-SiteMailboxProvisioningPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25568": {
		msg: "New-SyncMailPublicFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25569": {
		msg: "New-UMCallAnsweringRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25570": {
		msg: "New-WorkloadManagementPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25571": {
		msg: "New-WorkloadPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25572": {
		msg: "Redirect-Message Exchange cmdlet issued",
		src: "Exchange",
	},
	"25573": {
		msg: "Remove-AddressBookPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25574": {
		msg: "Remove-App Exchange cmdlet issued",
		src: "Exchange",
	},
	"25575": {
		msg: "Remove-AuthServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25576": {
		msg: "Remove-ClassificationRuleCollection Exchange cmdlet issued",
		src: "Exchange",
	},
	"25577": {
		msg: "Remove-DlpPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25578": {
		msg: "Remove-DlpPolicyTemplate Exchange cmdlet issued",
		src: "Exchange",
	},
	"25579": {
		msg: "Remove-GlobalMonitoringOverride Exchange cmdlet issued",
		src: "Exchange",
	},
	"25580": {
		msg: "Remove-HybridConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25581": {
		msg: "Remove-LinkedUser Exchange cmdlet issued",
		src: "Exchange",
	},
	"25582": {
		msg: "Remove-MailboxExportRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25583": {
		msg: "Remove-MailboxImportRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25584": {
		msg: "Remove-MailboxSearch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25585": {
		msg: "Remove-MalwareFilterPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25586": {
		msg: "Remove-MalwareFilterRecoveryItem Exchange cmdlet issued",
		src: "Exchange",
	},
	"25587": {
		msg: "Remove-MigrationBatch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25588": {
		msg: "Remove-MigrationEndpoint Exchange cmdlet issued",
		src: "Exchange",
	},
	"25589": {
		msg: "Remove-MigrationUser Exchange cmdlet issued",
		src: "Exchange",
	},
	"25590": {
		msg: "Remove-MobileDevice Exchange cmdlet issued",
		src: "Exchange",
	},
	"25591": {
		msg: "Remove-MobileDeviceMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25592": {
		msg: "Remove-OnPremisesOrganization Exchange cmdlet issued",
		src: "Exchange",
	},
	"25593": {
		msg: "Remove-PartnerApplication Exchange cmdlet issued",
		src: "Exchange",
	},
	"25594": {
		msg: "Remove-PolicyTipConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25595": {
		msg: "Remove-PowerShellVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25596": {
		msg: "Remove-PublicFolderMigrationRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25597": {
		msg: "Remove-ReOrigPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25598": {
		msg: "Remove-ResubmitRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25599": {
		msg: "Remove-SiteMailboxProvisioningPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25600": {
		msg: "Remove-UMCallAnsweringRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25601": {
		msg: "Remove-UserPhoto Exchange cmdlet issued",
		src: "Exchange",
	},
	"25602": {
		msg: "Remove-WorkloadManagementPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25603": {
		msg: "Remove-WorkloadPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25604": {
		msg: "Reset-ProvisioningCache Exchange cmdlet issued",
		src: "Exchange",
	},
	"25605": {
		msg: "Resume-MailboxImportRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25606": {
		msg: "Resume-MalwareFilterRecoveryItem Exchange cmdlet issued",
		src: "Exchange",
	},
	"25607": {
		msg: "Resume-PublicFolderMigrationRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25608": {
		msg: "Set-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25609": {
		msg: "Set-AddressBookPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25610": {
		msg: "Set-App Exchange cmdlet issued",
		src: "Exchange",
	},
	"25611": {
		msg: "Set-AuthConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25612": {
		msg: "Set-AuthServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25613": {
		msg: "Set-ClassificationRuleCollection Exchange cmdlet issued",
		src: "Exchange",
	},
	"25614": {
		msg: "Set-DlpPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25615": {
		msg: "Set-FrontendTransportService Exchange cmdlet issued",
		src: "Exchange",
	},
	"25616": {
		msg: "Set-HybridConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25617": {
		msg: "Set-HybridMailflow Exchange cmdlet issued",
		src: "Exchange",
	},
	"25618": {
		msg: "Set-MailboxExportRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25619": {
		msg: "Set-MailboxImportRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25620": {
		msg: "Set-MailboxSearch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25621": {
		msg: "Set-MailboxTransportService Exchange cmdlet issued",
		src: "Exchange",
	},
	"25622": {
		msg: "Set-MalwareFilteringServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25623": {
		msg: "Set-MalwareFilterPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25624": {
		msg: "Set-MigrationBatch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25625": {
		msg: "Set-MigrationConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25626": {
		msg: "Set-MigrationEndpoint Exchange cmdlet issued",
		src: "Exchange",
	},
	"25627": {
		msg: "Set-MobileDeviceMailboxPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25628": {
		msg: "Set-Notification Exchange cmdlet issued",
		src: "Exchange",
	},
	"25629": {
		msg: "Set-OnPremisesOrganization Exchange cmdlet issued",
		src: "Exchange",
	},
	"25630": {
		msg: "Set-PartnerApplication Exchange cmdlet issued",
		src: "Exchange",
	},
	"25631": {
		msg: "Set-PendingFederatedDomain Exchange cmdlet issued",
		src: "Exchange",
	},
	"25632": {
		msg: "Set-PolicyTipConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25633": {
		msg: "Set-PublicFolderMigrationRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25634": {
		msg: "Set-ReOrigPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25635": {
		msg: "Set-ResubmitRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25636": {
		msg: "Set-RMSTemplate Exchange cmdlet issued",
		src: "Exchange",
	},
	"25637": {
		msg: "Set-ServerComponentState Exchange cmdlet issued",
		src: "Exchange",
	},
	"25638": {
		msg: "Set-ServerMonitor Exchange cmdlet issued",
		src: "Exchange",
	},
	"25639": {
		msg: "Set-SiteMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25640": {
		msg: "Set-SiteMailboxProvisioningPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25641": {
		msg: "Set-TransportService Exchange cmdlet issued",
		src: "Exchange",
	},
	"25642": {
		msg: "Set-UMCallAnsweringRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25643": {
		msg: "Set-UMCallRouterSettings Exchange cmdlet issued",
		src: "Exchange",
	},
	"25644": {
		msg: "Set-UMService Exchange cmdlet issued",
		src: "Exchange",
	},
	"25645": {
		msg: "Set-UserPhoto Exchange cmdlet issued",
		src: "Exchange",
	},
	"25646": {
		msg: "Set-WorkloadPolicy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25647": {
		msg: "Start-MailboxSearch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25648": {
		msg: "Start-MigrationBatch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25649": {
		msg: "Stop-MailboxSearch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25650": {
		msg: "Stop-MigrationBatch Exchange cmdlet issued",
		src: "Exchange",
	},
	"25651": {
		msg: "Suspend-MailboxExportRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25652": {
		msg: "Suspend-MailboxImportRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25653": {
		msg: "Suspend-PublicFolderMigrationRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25654": {
		msg: "Test-ArchiveConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25655": {
		msg: "Test-MigrationServerAvailability Exchange cmdlet issued",
		src: "Exchange",
	},
	"25656": {
		msg: "Test-OAuthConnectivity Exchange cmdlet issued",
		src: "Exchange",
	},
	"25657": {
		msg: "Test-SiteMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25658": {
		msg: "Update-HybridConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25659": {
		msg: "Update-PublicFolderMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25660": {
		msg: "Update-SiteMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25661": {
		msg: "Add-AttachmentFilterEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25662": {
		msg: "Remove-AttachmentFilterEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25663": {
		msg: "New-AddressRewriteEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25664": {
		msg: "Remove-AddressRewriteEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25665": {
		msg: "Set-AddressRewriteEntry Exchange cmdlet issued",
		src: "Exchange",
	},
	"25666": {
		msg: "Set-AttachmentFilterListConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25667": {
		msg: "Set-MailboxSentItemsConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25668": {
		msg: "Update-MovedMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25669": {
		msg: "Disable-MalwareFilterRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25670": {
		msg: "Enable-MalwareFilterRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25671": {
		msg: "New-MalwareFilterRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25672": {
		msg: "Remove-MalwareFilterRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25673": {
		msg: "Set-MalwareFilterRule Exchange cmdlet issued",
		src: "Exchange",
	},
	"25674": {
		msg: "Remove-MailboxRepairRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25675": {
		msg: "Remove-ServerMonitoringOverride Exchange cmdlet issued",
		src: "Exchange",
	},
	"25676": {
		msg: "Update-ExchangeHelp Exchange cmdlet issued",
		src: "Exchange",
	},
	"25677": {
		msg: "Update-StoreMailboxState Exchange cmdlet issued",
		src: "Exchange",
	},
	"25678": {
		msg: "Disable-PushNotificationProxy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25679": {
		msg: "Enable-PushNotificationProxy Exchange cmdlet issued",
		src: "Exchange",
	},
	"25680": {
		msg: "New-PublicFolderMoveRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25681": {
		msg: "Remove-PublicFolderMoveRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25682": {
		msg: "Resume-PublicFolderMoveRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25683": {
		msg: "Set-PublicFolderMoveRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25684": {
		msg: "Suspend-PublicFolderMoveRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25685": {
		msg: "Update-DatabaseSchema Exchange cmdlet issued",
		src: "Exchange",
	},
	"25686": {
		msg: "Set-SearchDocumentFormat Exchange cmdlet issued",
		src: "Exchange",
	},
	"25687": {
		msg: "New-AuthRedirect Exchange cmdlet issued",
		src: "Exchange",
	},
	"25688": {
		msg: "New-CompliancePolicySyncNotification Exchange cmdlet issued",
		src: "Exchange",
	},
	"25689": {
		msg: "New-ComplianceServiceVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25690": {
		msg: "New-DatabaseAvailabilityGroupConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25691": {
		msg: "New-DataClassification Exchange cmdlet issued",
		src: "Exchange",
	},
	"25692": {
		msg: "New-Fingerprint Exchange cmdlet issued",
		src: "Exchange",
	},
	"25693": {
		msg: "New-IntraOrganizationConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25694": {
		msg: "New-MailboxDeliveryVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25695": {
		msg: "New-MapiVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25696": {
		msg: "New-OutlookServiceVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25697": {
		msg: "New-RestVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25698": {
		msg: "New-SearchDocumentFormat Exchange cmdlet issued",
		src: "Exchange",
	},
	"25699": {
		msg: "New-SettingOverride Exchange cmdlet issued",
		src: "Exchange",
	},
	"25700": {
		msg: "New-SiteMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25701": {
		msg: "Remove-AuthRedirect Exchange cmdlet issued",
		src: "Exchange",
	},
	"25702": {
		msg: "Remove-CompliancePolicySyncNotification Exchange cmdlet issued",
		src: "Exchange",
	},
	"25703": {
		msg: "Remove-ComplianceServiceVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25704": {
		msg: "Remove-DatabaseAvailabilityGroupConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25705": {
		msg: "Remove-DataClassification Exchange cmdlet issued",
		src: "Exchange",
	},
	"25706": {
		msg: "Remove-IntraOrganizationConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25707": {
		msg: "Remove-MailboxDeliveryVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25708": {
		msg: "Remove-MapiVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25709": {
		msg: "Remove-OutlookServiceVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25710": {
		msg: "Remove-PublicFolderMailboxMigrationRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25711": {
		msg: "Remove-PushNotificationSubscription Exchange cmdlet issued",
		src: "Exchange",
	},
	"25712": {
		msg: "Remove-RestVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25713": {
		msg: "Remove-SearchDocumentFormat Exchange cmdlet issued",
		src: "Exchange",
	},
	"25714": {
		msg: "Remove-SettingOverride Exchange cmdlet issued",
		src: "Exchange",
	},
	"25715": {
		msg: "Remove-SyncMailPublicFolder Exchange cmdlet issued",
		src: "Exchange",
	},
	"25716": {
		msg: "Resume-PublicFolderMailboxMigrationRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25717": {
		msg: "Send-MapiSubmitSystemProbe Exchange cmdlet issued",
		src: "Exchange",
	},
	"25718": {
		msg: "Set-AuthRedirect Exchange cmdlet issued",
		src: "Exchange",
	},
	"25719": {
		msg: "Set-ClientAccessService Exchange cmdlet issued",
		src: "Exchange",
	},
	"25720": {
		msg: "Set-Clutter Exchange cmdlet issued",
		src: "Exchange",
	},
	"25721": {
		msg: "Set-ComplianceServiceVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25722": {
		msg: "Set-ConsumerMailbox Exchange cmdlet issued",
		src: "Exchange",
	},
	"25723": {
		msg: "Set-DatabaseAvailabilityGroupConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25724": {
		msg: "Set-DataClassification Exchange cmdlet issued",
		src: "Exchange",
	},
	"25725": {
		msg: "Set-IntraOrganizationConnector Exchange cmdlet issued",
		src: "Exchange",
	},
	"25726": {
		msg: "Set-LogExportVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25727": {
		msg: "Set-MailboxDeliveryVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25728": {
		msg: "Set-MapiVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25729": {
		msg: "Set-OutlookServiceVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25730": {
		msg: "Set-PublicFolderMailboxMigrationRequest Exchange cmdlet issued",
		src: "Exchange",
	},
	"25731": {
		msg: "Set-RestVirtualDirectory Exchange cmdlet issued",
		src: "Exchange",
	},
	"25732": {
		msg: "Set-SettingOverride Exchange cmdlet issued",
		src: "Exchange",
	},
	"25733": {
		msg: "Set-SmimeConfig Exchange cmdlet issued",
		src: "Exchange",
	},
	"25734": {
		msg: "Set-SubmissionMalwareFilteringServer Exchange cmdlet issued",
		src: "Exchange",
	},
	"25735": {
		msg: "Set-UMMailboxConfiguration Exchange cmdlet issued",
		src: "Exchange",
	},
	"25736": {
		msg: "Set-UnifiedAuditSetting Exchange cmdlet issued",
		src: "Exchange",
	},
	"25737": {
		msg: "Start-AuditAssistant Exchange cmdlet issued",
		src: "Exchange",
	},
	"25738": {
		msg: "Start-UMPhoneSession Exchange cmdlet issued",
		src: "Exchange",
	},
	"25739": {
		msg: "Stop-UMPhoneSession Exchange cmdlet issued",
		src: "Exchange",
	},
	"25740": {
		msg: "Test-DataClassification Exchange cmdlet issued",
		src: "Exchange",
	},
	"25741": {
		msg: "Test-TextExtraction Exchange cmdlet issued",
		src: "Exchange",
	},
}
