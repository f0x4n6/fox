package events

var Database = map[string]struct {
	Msg string
	Src string
}{
	"1100": {
		Msg: "The event logging service has shut down",
		Src: "Windows",
	},
	"1101": {
		Msg: "Audit events have been dropped by the transport.",
		Src: "Windows",
	},
	"1102": {
		Msg: "The audit log was cleared",
		Src: "Windows",
	},
	"1104": {
		Msg: "The security Log is now full",
		Src: "Windows",
	},
	"1105": {
		Msg: "Event log automatic backup",
		Src: "Windows",
	},
	"1108": {
		Msg: "The event logging service encountered an error",
		Src: "Windows",
	},
	"4608": {
		Msg: "Windows is starting up",
		Src: "Windows",
	},
	"4609": {
		Msg: "Windows is shutting down",
		Src: "Windows",
	},
	"4610": {
		Msg: "An authentication package has been loaded by the Local Security Authority",
		Src: "Windows",
	},
	"4611": {
		Msg: "A trusted logon process has been registered with the Local Security Authority",
		Src: "Windows",
	},
	"4612": {
		Msg: "Internal resources allocated for the queuing of audit messages have been exhausted, leading to the loss of some audits.",
		Src: "Windows",
	},
	"4614": {
		Msg: "A notification package has been loaded by the Security Account Manager.",
		Src: "Windows",
	},
	"4615": {
		Msg: "Invalid use of LPC port",
		Src: "Windows",
	},
	"4616": {
		Msg: "The system time was changed.",
		Src: "Windows",
	},
	"4618": {
		Msg: "A monitored security event pattern has occurred",
		Src: "Windows",
	},
	"4621": {
		Msg: "Administrator recovered system from CrashOnAuditFail",
		Src: "Windows",
	},
	"4622": {
		Msg: "A security package has been loaded by the Local Security Authority.",
		Src: "Windows",
	},
	"4624": {
		Msg: "An account was successfully logged on",
		Src: "Windows",
	},
	"4625": {
		Msg: "An account failed to log on",
		Src: "Windows",
	},
	"4626": {
		Msg: "User/Device claims information",
		Src: "Windows",
	},
	"4627": {
		Msg: "Group membership information.",
		Src: "Windows",
	},
	"4634": {
		Msg: "An account was logged off",
		Src: "Windows",
	},
	"4646": {
		Msg: "IKE DoS-prevention mode started",
		Src: "Windows",
	},
	"4647": {
		Msg: "User initiated logoff",
		Src: "Windows",
	},
	"4648": {
		Msg: "A logon was attempted using explicit credentials",
		Src: "Windows",
	},
	"4649": {
		Msg: "A replay attack was detected",
		Src: "Windows",
	},
	"4650": {
		Msg: "An IPsec Main Mode security association was established",
		Src: "Windows",
	},
	"4651": {
		Msg: "An IPsec Main Mode security association was established",
		Src: "Windows",
	},
	"4652": {
		Msg: "An IPsec Main Mode negotiation failed",
		Src: "Windows",
	},
	"4653": {
		Msg: "An IPsec Main Mode negotiation failed",
		Src: "Windows",
	},
	"4654": {
		Msg: "An IPsec Quick Mode negotiation failed",
		Src: "Windows",
	},
	"4655": {
		Msg: "An IPsec Main Mode security association ended",
		Src: "Windows",
	},
	"4656": {
		Msg: "A handle to an object was requested",
		Src: "Windows",
	},
	"4657": {
		Msg: "A registry value was modified",
		Src: "Windows",
	},
	"4658": {
		Msg: "The handle to an object was closed",
		Src: "Windows",
	},
	"4659": {
		Msg: "A handle to an object was requested with intent to delete",
		Src: "Windows",
	},
	"4660": {
		Msg: "An object was deleted",
		Src: "Windows",
	},
	"4661": {
		Msg: "A handle to an object was requested",
		Src: "Windows",
	},
	"4662": {
		Msg: "An operation was performed on an object",
		Src: "Windows",
	},
	"4663": {
		Msg: "An attempt was made to access an object",
		Src: "Windows",
	},
	"4664": {
		Msg: "An attempt was made to create a hard link",
		Src: "Windows",
	},
	"4665": {
		Msg: "An attempt was made to create an application client context.",
		Src: "Windows",
	},
	"4666": {
		Msg: "An application attempted an operation",
		Src: "Windows",
	},
	"4667": {
		Msg: "An application client context was deleted",
		Src: "Windows",
	},
	"4668": {
		Msg: "An application was initialized",
		Src: "Windows",
	},
	"4670": {
		Msg: "Permissions on an object were changed",
		Src: "Windows",
	},
	"4671": {
		Msg: "An application attempted to access a blocked ordinal through the TBS",
		Src: "Windows",
	},
	"4672": {
		Msg: "Special privileges assigned to new logon",
		Src: "Windows",
	},
	"4673": {
		Msg: "A privileged service was called",
		Src: "Windows",
	},
	"4674": {
		Msg: "An operation was attempted on a privileged object",
		Src: "Windows",
	},
	"4675": {
		Msg: "SIDs were filtered",
		Src: "Windows",
	},
	"4688": {
		Msg: "A new process has been created",
		Src: "Windows",
	},
	"4689": {
		Msg: "A process has exited",
		Src: "Windows",
	},
	"4690": {
		Msg: "An attempt was made to duplicate a handle to an object",
		Src: "Windows",
	},
	"4691": {
		Msg: "Indirect access to an object was requested",
		Src: "Windows",
	},
	"4692": {
		Msg: "Backup of data protection master key was attempted",
		Src: "Windows",
	},
	"4693": {
		Msg: "Recovery of data protection master key was attempted",
		Src: "Windows",
	},
	"4694": {
		Msg: "Protection of auditable protected data was attempted",
		Src: "Windows",
	},
	"4695": {
		Msg: "Unprotection of auditable protected data was attempted",
		Src: "Windows",
	},
	"4696": {
		Msg: "A primary token was assigned to process",
		Src: "Windows",
	},
	"4697": {
		Msg: "A service was installed in the system",
		Src: "Windows",
	},
	"4698": {
		Msg: "A scheduled task was created",
		Src: "Windows",
	},
	"4699": {
		Msg: "A scheduled task was deleted",
		Src: "Windows",
	},
	"4700": {
		Msg: "A scheduled task was enabled",
		Src: "Windows",
	},
	"4701": {
		Msg: "A scheduled task was disabled",
		Src: "Windows",
	},
	"4702": {
		Msg: "A scheduled task was updated",
		Src: "Windows",
	},
	"4703": {
		Msg: "A token right was adjusted",
		Src: "Windows",
	},
	"4704": {
		Msg: "A user right was assigned",
		Src: "Windows",
	},
	"4705": {
		Msg: "A user right was removed",
		Src: "Windows",
	},
	"4706": {
		Msg: "A new trust was created to a domain",
		Src: "Windows",
	},
	"4707": {
		Msg: "A trust to a domain was removed",
		Src: "Windows",
	},
	"4709": {
		Msg: "IPsec Services was started",
		Src: "Windows",
	},
	"4710": {
		Msg: "IPsec Services was disabled",
		Src: "Windows",
	},
	"4711": {
		Msg: "PAStore Engine (1%)",
		Src: "Windows",
	},
	"4712": {
		Msg: "IPsec Services encountered a potentially serious failure",
		Src: "Windows",
	},
	"4713": {
		Msg: "Kerberos policy was changed",
		Src: "Windows",
	},
	"4714": {
		Msg: "Encrypted data recovery policy was changed",
		Src: "Windows",
	},
	"4715": {
		Msg: "The audit policy (SACL) on an object was changed",
		Src: "Windows",
	},
	"4716": {
		Msg: "Trusted domain information was modified",
		Src: "Windows",
	},
	"4717": {
		Msg: "System security access was granted to an account",
		Src: "Windows",
	},
	"4718": {
		Msg: "System security access was removed from an account",
		Src: "Windows",
	},
	"4719": {
		Msg: "System audit policy was changed",
		Src: "Windows",
	},
	"4720": {
		Msg: "A user account was created",
		Src: "Windows",
	},
	"4722": {
		Msg: "A user account was enabled",
		Src: "Windows",
	},
	"4723": {
		Msg: "An attempt was made to change an account's password",
		Src: "Windows",
	},
	"4724": {
		Msg: "An attempt was made to reset an accounts password",
		Src: "Windows",
	},
	"4725": {
		Msg: "A user account was disabled",
		Src: "Windows",
	},
	"4726": {
		Msg: "A user account was deleted",
		Src: "Windows",
	},
	"4727": {
		Msg: "A security-enabled global group was created",
		Src: "Windows",
	},
	"4728": {
		Msg: "A member was added to a security-enabled global group",
		Src: "Windows",
	},
	"4729": {
		Msg: "A member was removed from a security-enabled global group",
		Src: "Windows",
	},
	"4730": {
		Msg: "A security-enabled global group was deleted",
		Src: "Windows",
	},
	"4731": {
		Msg: "A security-enabled local group was created",
		Src: "Windows",
	},
	"4732": {
		Msg: "A member was added to a security-enabled local group",
		Src: "Windows",
	},
	"4733": {
		Msg: "A member was removed from a security-enabled local group",
		Src: "Windows",
	},
	"4734": {
		Msg: "A security-enabled local group was deleted",
		Src: "Windows",
	},
	"4735": {
		Msg: "A security-enabled local group was changed",
		Src: "Windows",
	},
	"4737": {
		Msg: "A security-enabled global group was changed",
		Src: "Windows",
	},
	"4738": {
		Msg: "A user account was changed",
		Src: "Windows",
	},
	"4739": {
		Msg: "Domain Policy was changed",
		Src: "Windows",
	},
	"4740": {
		Msg: "A user account was locked out",
		Src: "Windows",
	},
	"4741": {
		Msg: "A computer account was created",
		Src: "Windows",
	},
	"4742": {
		Msg: "A computer account was changed",
		Src: "Windows",
	},
	"4743": {
		Msg: "A computer account was deleted",
		Src: "Windows",
	},
	"4744": {
		Msg: "A security-disabled local group was created",
		Src: "Windows",
	},
	"4745": {
		Msg: "A security-disabled local group was changed",
		Src: "Windows",
	},
	"4746": {
		Msg: "A member was added to a security-disabled local group",
		Src: "Windows",
	},
	"4747": {
		Msg: "A member was removed from a security-disabled local group",
		Src: "Windows",
	},
	"4748": {
		Msg: "A security-disabled local group was deleted",
		Src: "Windows",
	},
	"4749": {
		Msg: "A security-disabled global group was created",
		Src: "Windows",
	},
	"4750": {
		Msg: "A security-disabled global group was changed",
		Src: "Windows",
	},
	"4751": {
		Msg: "A member was added to a security-disabled global group",
		Src: "Windows",
	},
	"4752": {
		Msg: "A member was removed from a security-disabled global group",
		Src: "Windows",
	},
	"4753": {
		Msg: "A security-disabled global group was deleted",
		Src: "Windows",
	},
	"4754": {
		Msg: "A security-enabled universal group was created",
		Src: "Windows",
	},
	"4755": {
		Msg: "A security-enabled universal group was changed",
		Src: "Windows",
	},
	"4756": {
		Msg: "A member was added to a security-enabled universal group",
		Src: "Windows",
	},
	"4757": {
		Msg: "A member was removed from a security-enabled universal group",
		Src: "Windows",
	},
	"4758": {
		Msg: "A security-enabled universal group was deleted",
		Src: "Windows",
	},
	"4759": {
		Msg: "A security-disabled universal group was created",
		Src: "Windows",
	},
	"4760": {
		Msg: "A security-disabled universal group was changed",
		Src: "Windows",
	},
	"4761": {
		Msg: "A member was added to a security-disabled universal group",
		Src: "Windows",
	},
	"4762": {
		Msg: "A member was removed from a security-disabled universal group",
		Src: "Windows",
	},
	"4763": {
		Msg: "A security-disabled universal group was deleted",
		Src: "Windows",
	},
	"4764": {
		Msg: "A groups type was changed",
		Src: "Windows",
	},
	"4765": {
		Msg: "SID History was added to an account",
		Src: "Windows",
	},
	"4766": {
		Msg: "An attempt to add SID History to an account failed",
		Src: "Windows",
	},
	"4767": {
		Msg: "A user account was unlocked",
		Src: "Windows",
	},
	"4768": {
		Msg: "A Kerberos authentication ticket (TGT) was requested",
		Src: "Windows",
	},
	"4769": {
		Msg: "A Kerberos service ticket was requested",
		Src: "Windows",
	},
	"4770": {
		Msg: "A Kerberos service ticket was renewed",
		Src: "Windows",
	},
	"4771": {
		Msg: "Kerberos pre-authentication failed",
		Src: "Windows",
	},
	"4772": {
		Msg: "A Kerberos authentication ticket request failed",
		Src: "Windows",
	},
	"4773": {
		Msg: "A Kerberos service ticket request failed",
		Src: "Windows",
	},
	"4774": {
		Msg: "An account was mapped for logon",
		Src: "Windows",
	},
	"4775": {
		Msg: "An account could not be mapped for logon",
		Src: "Windows",
	},
	"4776": {
		Msg: "The domain controller attempted to validate the credentials for an account",
		Src: "Windows",
	},
	"4777": {
		Msg: "The domain controller failed to validate the credentials for an account",
		Src: "Windows",
	},
	"4778": {
		Msg: "A session was reconnected to a Window Station",
		Src: "Windows",
	},
	"4779": {
		Msg: "A session was disconnected from a Window Station",
		Src: "Windows",
	},
	"4780": {
		Msg: "The ACL was set on accounts which are members of administrators groups",
		Src: "Windows",
	},
	"4781": {
		Msg: "The Data of an account was changed",
		Src: "Windows",
	},
	"4782": {
		Msg: "The password hash an account was accessed",
		Src: "Windows",
	},
	"4783": {
		Msg: "A basic application group was created",
		Src: "Windows",
	},
	"4784": {
		Msg: "A basic application group was changed",
		Src: "Windows",
	},
	"4785": {
		Msg: "A member was added to a basic application group",
		Src: "Windows",
	},
	"4786": {
		Msg: "A member was removed from a basic application group",
		Src: "Windows",
	},
	"4787": {
		Msg: "A non-member was added to a basic application group",
		Src: "Windows",
	},
	"4788": {
		Msg: "A non-member was removed from a basic application group..",
		Src: "Windows",
	},
	"4789": {
		Msg: "A basic application group was deleted",
		Src: "Windows",
	},
	"4790": {
		Msg: "An LDAP query group was created",
		Src: "Windows",
	},
	"4791": {
		Msg: "A basic application group was changed",
		Src: "Windows",
	},
	"4792": {
		Msg: "An LDAP query group was deleted",
		Src: "Windows",
	},
	"4793": {
		Msg: "The Password Policy Checking API was called",
		Src: "Windows",
	},
	"4794": {
		Msg: "An attempt was made to set the Directory Services Restore Mode administrator password",
		Src: "Windows",
	},
	"4797": {
		Msg: "An attempt was made to query the existence of a blank password for an account",
		Src: "Windows",
	},
	"4798": {
		Msg: "A user's local group membership was enumerated.",
		Src: "Windows",
	},
	"4799": {
		Msg: "A security-enabled local group membership was enumerated",
		Src: "Windows",
	},
	"4800": {
		Msg: "The workstation was locked",
		Src: "Windows",
	},
	"4801": {
		Msg: "The workstation was unlocked",
		Src: "Windows",
	},
	"4802": {
		Msg: "The screen saver was invoked",
		Src: "Windows",
	},
	"4803": {
		Msg: "The screen saver was dismissed",
		Src: "Windows",
	},
	"4816": {
		Msg: "RPC detected an integrity violation while decrypting an incoming message",
		Src: "Windows",
	},
	"4817": {
		Msg: "Auditing settings on object were changed.",
		Src: "Windows",
	},
	"4818": {
		Msg: "Proposed Central Access Policy does not grant the same access permissions as the current Central Access Policy",
		Src: "Windows",
	},
	"4819": {
		Msg: "Central Access Policies on the machine have been changed",
		Src: "Windows",
	},
	"4820": {
		Msg: "A Kerberos Ticket-granting-ticket (TGT) was denied because the device does not meet the access control restrictions",
		Src: "Windows",
	},
	"4821": {
		Msg: "A Kerberos service ticket was denied because the user, device, or both does not meet the access control restrictions",
		Src: "Windows",
	},
	"4822": {
		Msg: "NTLM authentication failed because the account was a member of the Protected User group",
		Src: "Windows",
	},
	"4823": {
		Msg: "NTLM authentication failed because access control restrictions are required",
		Src: "Windows",
	},
	"4824": {
		Msg: "Kerberos preauthentication by using DES or RC4 failed because the account was a member of the Protected User group",
		Src: "Windows",
	},
	"4825": {
		Msg: "A user was denied the access to Remote Desktop. By default, users are allowed to connect only if they are members of the Remote Desktop Users group or Administrators group",
		Src: "Windows",
	},
	"4826": {
		Msg: "Boot Configuration Data loaded",
		Src: "Windows",
	},
	"4830": {
		Msg: "SID History was removed from an account",
		Src: "Windows",
	},
	"4864": {
		Msg: "A Dataspace collision was detected",
		Src: "Windows",
	},
	"4865": {
		Msg: "A trusted forest information entry was added",
		Src: "Windows",
	},
	"4866": {
		Msg: "A trusted forest information entry was removed",
		Src: "Windows",
	},
	"4867": {
		Msg: "A trusted forest information entry was modified",
		Src: "Windows",
	},
	"4868": {
		Msg: "The certificate manager denied a pending certificate request",
		Src: "Windows",
	},
	"4869": {
		Msg: "Certificate Services received a resubmitted certificate request",
		Src: "Windows",
	},
	"4870": {
		Msg: "Certificate Services revoked a certificate",
		Src: "Windows",
	},
	"4871": {
		Msg: "Certificate Services received a request to publish the certificate revocation list (CRL)",
		Src: "Windows",
	},
	"4872": {
		Msg: "Certificate Services published the certificate revocation list (CRL)",
		Src: "Windows",
	},
	"4873": {
		Msg: "A certificate request extension changed",
		Src: "Windows",
	},
	"4874": {
		Msg: "One or more certificate request attributes changed.",
		Src: "Windows",
	},
	"4875": {
		Msg: "Certificate Services received a request to shut down",
		Src: "Windows",
	},
	"4876": {
		Msg: "Certificate Services backup started",
		Src: "Windows",
	},
	"4877": {
		Msg: "Certificate Services backup completed",
		Src: "Windows",
	},
	"4878": {
		Msg: "Certificate Services restore started",
		Src: "Windows",
	},
	"4879": {
		Msg: "Certificate Services restore completed",
		Src: "Windows",
	},
	"4880": {
		Msg: "Certificate Services started",
		Src: "Windows",
	},
	"4881": {
		Msg: "Certificate Services stopped",
		Src: "Windows",
	},
	"4882": {
		Msg: "The security permissions for Certificate Services changed",
		Src: "Windows",
	},
	"4883": {
		Msg: "Certificate Services retrieved an archived key",
		Src: "Windows",
	},
	"4884": {
		Msg: "Certificate Services imported a certificate into its database",
		Src: "Windows",
	},
	"4885": {
		Msg: "The audit filter for Certificate Services changed",
		Src: "Windows",
	},
	"4886": {
		Msg: "Certificate Services received a certificate request",
		Src: "Windows",
	},
	"4887": {
		Msg: "Certificate Services approved a certificate request and issued a certificate",
		Src: "Windows",
	},
	"4888": {
		Msg: "Certificate Services denied a certificate request",
		Src: "Windows",
	},
	"4889": {
		Msg: "Certificate Services set the status of a certificate request to pending",
		Src: "Windows",
	},
	"4890": {
		Msg: "The certificate manager settings for Certificate Services changed.",
		Src: "Windows",
	},
	"4891": {
		Msg: "A configuration entry changed in Certificate Services",
		Src: "Windows",
	},
	"4892": {
		Msg: "A property of Certificate Services changed",
		Src: "Windows",
	},
	"4893": {
		Msg: "Certificate Services archived a key",
		Src: "Windows",
	},
	"4894": {
		Msg: "Certificate Services imported and archived a key",
		Src: "Windows",
	},
	"4895": {
		Msg: "Certificate Services published the CA certificate to Active Directory Domain Services",
		Src: "Windows",
	},
	"4896": {
		Msg: "One or more rows have been deleted from the certificate database",
		Src: "Windows",
	},
	"4897": {
		Msg: "Role separation enabled",
		Src: "Windows",
	},
	"4898": {
		Msg: "Certificate Services loaded a template",
		Src: "Windows",
	},
	"4899": {
		Msg: "A Certificate Services template was updated",
		Src: "Windows",
	},
	"4900": {
		Msg: "Certificate Services template security was updated",
		Src: "Windows",
	},
	"4902": {
		Msg: "The Per-user audit policy table was created",
		Src: "Windows",
	},
	"4904": {
		Msg: "An attempt was made to register a security event Orig",
		Src: "Windows",
	},
	"4905": {
		Msg: "An attempt was made to unregister a security event Orig",
		Src: "Windows",
	},
	"4906": {
		Msg: "The CrashOnAuditFail value has changed",
		Src: "Windows",
	},
	"4907": {
		Msg: "Auditing settings on object were changed",
		Src: "Windows",
	},
	"4908": {
		Msg: "Special Groups Logon table modified",
		Src: "Windows",
	},
	"4909": {
		Msg: "The local policy settings for the TBS were changed",
		Src: "Windows",
	},
	"4910": {
		Msg: "The group policy settings for the TBS were changed",
		Src: "Windows",
	},
	"4911": {
		Msg: "ReOrig attributes of the object were changed",
		Src: "Windows",
	},
	"4912": {
		Msg: "Per User Audit Policy was changed",
		Src: "Windows",
	},
	"4913": {
		Msg: "Central Access Policy on the object was changed",
		Src: "Windows",
	},
	"4928": {
		Msg: "An Active Directory replica Orig naming context was established",
		Src: "Windows",
	},
	"4929": {
		Msg: "An Active Directory replica Orig naming context was removed",
		Src: "Windows",
	},
	"4930": {
		Msg: "An Active Directory replica Orig naming context was modified",
		Src: "Windows",
	},
	"4931": {
		Msg: "An Active Directory replica destination naming context was modified",
		Src: "Windows",
	},
	"4932": {
		Msg: "Synchronization of a replica of an Active Directory naming context has begun",
		Src: "Windows",
	},
	"4933": {
		Msg: "Synchronization of a replica of an Active Directory naming context has ended",
		Src: "Windows",
	},
	"4934": {
		Msg: "Attributes of an Active Directory object were replicated",
		Src: "Windows",
	},
	"4935": {
		Msg: "Replication failure begins",
		Src: "Windows",
	},
	"4936": {
		Msg: "Replication failure ends",
		Src: "Windows",
	},
	"4937": {
		Msg: "A lingering object was removed from a replica",
		Src: "Windows",
	},
	"4944": {
		Msg: "The following policy was active when the Windows Firewall started",
		Src: "Windows",
	},
	"4945": {
		Msg: "A rule was listed when the Windows Firewall started",
		Src: "Windows",
	},
	"4946": {
		Msg: "A change has been made to Windows Firewall exception list. A rule was added",
		Src: "Windows",
	},
	"4947": {
		Msg: "A change has been made to Windows Firewall exception list. A rule was modified",
		Src: "Windows",
	},
	"4948": {
		Msg: "A change has been made to Windows Firewall exception list. A rule was deleted",
		Src: "Windows",
	},
	"4949": {
		Msg: "Windows Firewall settings were restored to the default values",
		Src: "Windows",
	},
	"4950": {
		Msg: "A Windows Firewall setting has changed",
		Src: "Windows",
	},
	"4951": {
		Msg: "A rule has been ignored because its major version number was not recognized by Windows Firewall",
		Src: "Windows",
	},
	"4952": {
		Msg: "Parts of a rule have been ignored because its minor version number was not recognized by Windows Firewall",
		Src: "Windows",
	},
	"4953": {
		Msg: "A rule has been ignored by Windows Firewall because it could not parse the rule",
		Src: "Windows",
	},
	"4954": {
		Msg: "Windows Firewall Group Policy settings has changed. The new settings have been applied",
		Src: "Windows",
	},
	"4956": {
		Msg: "Windows Firewall has changed the active profile",
		Src: "Windows",
	},
	"4957": {
		Msg: "Windows Firewall did not apply the following rule",
		Src: "Windows",
	},
	"4958": {
		Msg: "Windows Firewall did not apply the following rule because the rule referred to items not configured on this computer",
		Src: "Windows",
	},
	"4960": {
		Msg: "IPsec dropped an inbound packet that failed an integrity check",
		Src: "Windows",
	},
	"4961": {
		Msg: "IPsec dropped an inbound packet that failed a replay check",
		Src: "Windows",
	},
	"4962": {
		Msg: "IPsec dropped an inbound packet that failed a replay check",
		Src: "Windows",
	},
	"4963": {
		Msg: "IPsec dropped an inbound clear text packet that should have been secured",
		Src: "Windows",
	},
	"4964": {
		Msg: "Special groups have been assigned to a new logon",
		Src: "Windows",
	},
	"4965": {
		Msg: "IPsec received a packet from a remote computer with an incorrect Security Parameter Index (SPI).",
		Src: "Windows",
	},
	"4976": {
		Msg: "During Main Mode negotiation, IPsec received an invalid negotiation packet.",
		Src: "Windows",
	},
	"4977": {
		Msg: "During Quick Mode negotiation, IPsec received an invalid negotiation packet.",
		Src: "Windows",
	},
	"4978": {
		Msg: "During Extended Mode negotiation, IPsec received an invalid negotiation packet.",
		Src: "Windows",
	},
	"4979": {
		Msg: "IPsec Main Mode and Extended Mode security associations were established.",
		Src: "Windows",
	},
	"4980": {
		Msg: "IPsec Main Mode and Extended Mode security associations were established",
		Src: "Windows",
	},
	"4981": {
		Msg: "IPsec Main Mode and Extended Mode security associations were established",
		Src: "Windows",
	},
	"4982": {
		Msg: "IPsec Main Mode and Extended Mode security associations were established",
		Src: "Windows",
	},
	"4983": {
		Msg: "An IPsec Extended Mode negotiation failed",
		Src: "Windows",
	},
	"4984": {
		Msg: "An IPsec Extended Mode negotiation failed",
		Src: "Windows",
	},
	"4985": {
		Msg: "The state of a transaction has changed",
		Src: "Windows",
	},
	"5024": {
		Msg: "The Windows Firewall Service has started successfully",
		Src: "Windows",
	},
	"5025": {
		Msg: "The Windows Firewall Service has been stopped",
		Src: "Windows",
	},
	"5027": {
		Msg: "The Windows Firewall Service was unable to retrieve the security policy from the local storage",
		Src: "Windows",
	},
	"5028": {
		Msg: "The Windows Firewall Service was unable to parse the new security policy.",
		Src: "Windows",
	},
	"5029": {
		Msg: "The Windows Firewall Service failed to initialize the driver",
		Src: "Windows",
	},
	"5030": {
		Msg: "The Windows Firewall Service failed to start",
		Src: "Windows",
	},
	"5031": {
		Msg: "The Windows Firewall Service blocked an application from accepting incoming connections on the network.",
		Src: "Windows",
	},
	"5032": {
		Msg: "Windows Firewall was unable to notify the user that it blocked an application from accepting incoming connections on the network",
		Src: "Windows",
	},
	"5033": {
		Msg: "The Windows Firewall Driver has started successfully",
		Src: "Windows",
	},
	"5034": {
		Msg: "The Windows Firewall Driver has been stopped",
		Src: "Windows",
	},
	"5035": {
		Msg: "The Windows Firewall Driver failed to start",
		Src: "Windows",
	},
	"5037": {
		Msg: "The Windows Firewall Driver detected critical runtime error. Terminating",
		Src: "Windows",
	},
	"5038": {
		Msg: "Code integrity determined that the image hash of a file is not valid",
		Src: "Windows",
	},
	"5039": {
		Msg: "A registry key was virtualized.",
		Src: "Windows",
	},
	"5040": {
		Msg: "A change has been made to IPsec settings. An Authentication Set was added.",
		Src: "Windows",
	},
	"5041": {
		Msg: "A change has been made to IPsec settings. An Authentication Set was modified",
		Src: "Windows",
	},
	"5042": {
		Msg: "A change has been made to IPsec settings. An Authentication Set was deleted",
		Src: "Windows",
	},
	"5043": {
		Msg: "A change has been made to IPsec settings. A Connection Security Rule was added",
		Src: "Windows",
	},
	"5044": {
		Msg: "A change has been made to IPsec settings. A Connection Security Rule was modified",
		Src: "Windows",
	},
	"5045": {
		Msg: "A change has been made to IPsec settings. A Connection Security Rule was deleted",
		Src: "Windows",
	},
	"5046": {
		Msg: "A change has been made to IPsec settings. A Crypto Set was added",
		Src: "Windows",
	},
	"5047": {
		Msg: "A change has been made to IPsec settings. A Crypto Set was modified",
		Src: "Windows",
	},
	"5048": {
		Msg: "A change has been made to IPsec settings. A Crypto Set was deleted",
		Src: "Windows",
	},
	"5049": {
		Msg: "An IPsec Security Association was deleted",
		Src: "Windows",
	},
	"5050": {
		Msg: "An attempt to programmatically disable the Windows Firewall using a call to INetFwProfile.FirewallEnabled(FALSE",
		Src: "Windows",
	},
	"5051": {
		Msg: "A file was virtualized",
		Src: "Windows",
	},
	"5056": {
		Msg: "A cryptographic self test was performed",
		Src: "Windows",
	},
	"5057": {
		Msg: "A cryptographic primitive operation failed",
		Src: "Windows",
	},
	"5058": {
		Msg: "Key file operation",
		Src: "Windows",
	},
	"5059": {
		Msg: "Key migration operation",
		Src: "Windows",
	},
	"5060": {
		Msg: "Verification operation failed",
		Src: "Windows",
	},
	"5061": {
		Msg: "Cryptographic operation",
		Src: "Windows",
	},
	"5062": {
		Msg: "A kernel-mode cryptographic self test was performed",
		Src: "Windows",
	},
	"5063": {
		Msg: "A cryptographic provider operation was attempted",
		Src: "Windows",
	},
	"5064": {
		Msg: "A cryptographic context operation was attempted",
		Src: "Windows",
	},
	"5065": {
		Msg: "A cryptographic context modification was attempted",
		Src: "Windows",
	},
	"5066": {
		Msg: "A cryptographic function operation was attempted",
		Src: "Windows",
	},
	"5067": {
		Msg: "A cryptographic function modification was attempted",
		Src: "Windows",
	},
	"5068": {
		Msg: "A cryptographic function provider operation was attempted",
		Src: "Windows",
	},
	"5069": {
		Msg: "A cryptographic function property operation was attempted",
		Src: "Windows",
	},
	"5070": {
		Msg: "A cryptographic function property operation was attempted",
		Src: "Windows",
	},
	"5071": {
		Msg: "Key access denied by Microsoft key distribution service",
		Src: "Windows",
	},
	"5120": {
		Msg: "OCSP Responder Service Started",
		Src: "Windows",
	},
	"5121": {
		Msg: "OCSP Responder Service Stopped",
		Src: "Windows",
	},
	"5122": {
		Msg: "A Configuration entry changed in the OCSP Responder Service",
		Src: "Windows",
	},
	"5123": {
		Msg: "A configuration entry changed in the OCSP Responder Service",
		Src: "Windows",
	},
	"5124": {
		Msg: "A security setting was updated on OCSP Responder Service",
		Src: "Windows",
	},
	"5125": {
		Msg: "A request was submitted to OCSP Responder Service",
		Src: "Windows",
	},
	"5126": {
		Msg: "Signing Certificate was automatically updated by the OCSP Responder Service",
		Src: "Windows",
	},
	"5127": {
		Msg: "The OCSP Revocation Provider successfully updated the revocation information",
		Src: "Windows",
	},
	"5136": {
		Msg: "A directory service object was modified",
		Src: "Windows",
	},
	"5137": {
		Msg: "A directory service object was created",
		Src: "Windows",
	},
	"5138": {
		Msg: "A directory service object was undeleted",
		Src: "Windows",
	},
	"5139": {
		Msg: "A directory service object was moved",
		Src: "Windows",
	},
	"5140": {
		Msg: "A network share object was accessed",
		Src: "Windows",
	},
	"5141": {
		Msg: "A directory service object was deleted",
		Src: "Windows",
	},
	"5142": {
		Msg: "A network share object was added.",
		Src: "Windows",
	},
	"5143": {
		Msg: "A network share object was modified",
		Src: "Windows",
	},
	"5144": {
		Msg: "A network share object was deleted.",
		Src: "Windows",
	},
	"5145": {
		Msg: "A network share object was checked to see whether client can be granted desired access",
		Src: "Windows",
	},
	"5146": {
		Msg: "The Windows Filtering Platform has blocked a packet",
		Src: "Windows",
	},
	"5147": {
		Msg: "A more restrictive Windows Filtering Platform filter has blocked a packet",
		Src: "Windows",
	},
	"5148": {
		Msg: "The Windows Filtering Platform has detected a DoS attack and entered a defensive mode; packets associated with this attack will be discarded.",
		Src: "Windows",
	},
	"5149": {
		Msg: "The DoS attack has subsided and normal processing is being resumed.",
		Src: "Windows",
	},
	"5150": {
		Msg: "The Windows Filtering Platform has blocked a packet.",
		Src: "Windows",
	},
	"5151": {
		Msg: "A more restrictive Windows Filtering Platform filter has blocked a packet.",
		Src: "Windows",
	},
	"5152": {
		Msg: "The Windows Filtering Platform blocked a packet",
		Src: "Windows",
	},
	"5153": {
		Msg: "A more restrictive Windows Filtering Platform filter has blocked a packet",
		Src: "Windows",
	},
	"5154": {
		Msg: "The Windows Filtering Platform has permitted an application or service to listen on a port for incoming connections",
		Src: "Windows",
	},
	"5155": {
		Msg: "The Windows Filtering Platform has blocked an application or service from listening on a port for incoming connections",
		Src: "Windows",
	},
	"5156": {
		Msg: "The Windows Filtering Platform has allowed a connection",
		Src: "Windows",
	},
	"5157": {
		Msg: "The Windows Filtering Platform has blocked a connection",
		Src: "Windows",
	},
	"5158": {
		Msg: "The Windows Filtering Platform has permitted a bind to a local port",
		Src: "Windows",
	},
	"5159": {
		Msg: "The Windows Filtering Platform has blocked a bind to a local port",
		Src: "Windows",
	},
	"5168": {
		Msg: "Spn check for SMB/SMB2 fails.",
		Src: "Windows",
	},
	"5169": {
		Msg: "A directory service object was modified",
		Src: "Windows",
	},
	"5170": {
		Msg: "A directory service object was modified during a background cleanup task",
		Src: "Windows",
	},
	"5376": {
		Msg: "Credential Manager credentials were backed up",
		Src: "Windows",
	},
	"5377": {
		Msg: "Credential Manager credentials were restored from a backup",
		Src: "Windows",
	},
	"5378": {
		Msg: "The requested credentials delegation was disallowed by policy",
		Src: "Windows",
	},
	"5440": {
		Msg: "The following callout was present when the Windows Filtering Platform Base Filtering Engine started",
		Src: "Windows",
	},
	"5441": {
		Msg: "The following filter was present when the Windows Filtering Platform Base Filtering Engine started",
		Src: "Windows",
	},
	"5442": {
		Msg: "The following provider was present when the Windows Filtering Platform Base Filtering Engine started",
		Src: "Windows",
	},
	"5443": {
		Msg: "The following provider context was present when the Windows Filtering Platform Base Filtering Engine started",
		Src: "Windows",
	},
	"5444": {
		Msg: "The following sub-layer was present when the Windows Filtering Platform Base Filtering Engine started",
		Src: "Windows",
	},
	"5446": {
		Msg: "A Windows Filtering Platform callout has been changed",
		Src: "Windows",
	},
	"5447": {
		Msg: "A Windows Filtering Platform filter has been changed",
		Src: "Windows",
	},
	"5448": {
		Msg: "A Windows Filtering Platform provider has been changed",
		Src: "Windows",
	},
	"5449": {
		Msg: "A Windows Filtering Platform provider context has been changed",
		Src: "Windows",
	},
	"5450": {
		Msg: "A Windows Filtering Platform sub-layer has been changed",
		Src: "Windows",
	},
	"5451": {
		Msg: "An IPsec Quick Mode security association was established",
		Src: "Windows",
	},
	"5452": {
		Msg: "An IPsec Quick Mode security association ended",
		Src: "Windows",
	},
	"5453": {
		Msg: "An IPsec negotiation with a remote computer failed because the IKE and AuthIP IPsec Keying Modules (IKEEXT) service is not started",
		Src: "Windows",
	},
	"5456": {
		Msg: "PAStore Engine applied Active Directory storage IPsec policy on the computer",
		Src: "Windows",
	},
	"5457": {
		Msg: "PAStore Engine failed to apply Active Directory storage IPsec policy on the computer",
		Src: "Windows",
	},
	"5458": {
		Msg: "PAStore Engine applied locally cached copy of Active Directory storage IPsec policy on the computer",
		Src: "Windows",
	},
	"5459": {
		Msg: "PAStore Engine failed to apply locally cached copy of Active Directory storage IPsec policy on the computer",
		Src: "Windows",
	},
	"5460": {
		Msg: "PAStore Engine applied local registry storage IPsec policy on the computer",
		Src: "Windows",
	},
	"5461": {
		Msg: "PAStore Engine failed to apply local registry storage IPsec policy on the computer",
		Src: "Windows",
	},
	"5462": {
		Msg: "PAStore Engine failed to apply some rules of the active IPsec policy on the computer",
		Src: "Windows",
	},
	"5463": {
		Msg: "PAStore Engine polled for changes to the active IPsec policy and detected no changes",
		Src: "Windows",
	},
	"5464": {
		Msg: "PAStore Engine polled for changes to the active IPsec policy, detected changes, and applied them to IPsec Services",
		Src: "Windows",
	},
	"5465": {
		Msg: "PAStore Engine received a control for forced reloading of IPsec policy and processed the control successfully",
		Src: "Windows",
	},
	"5466": {
		Msg: "PAStore Engine polled for changes to the Active Directory IPsec policy, determined that Active Directory cannot be reached, and will use the cached copy of the Active Directory IPsec policy instead",
		Src: "Windows",
	},
	"5467": {
		Msg: "PAStore Engine polled for changes to the Active Directory IPsec policy, determined that Active Directory can be reached, and found no changes to the policy",
		Src: "Windows",
	},
	"5468": {
		Msg: "PAStore Engine polled for changes to the Active Directory IPsec policy, determined that Active Directory can be reached, found changes to the policy, and applied those changes",
		Src: "Windows",
	},
	"5471": {
		Msg: "PAStore Engine loaded local storage IPsec policy on the computer",
		Src: "Windows",
	},
	"5472": {
		Msg: "PAStore Engine failed to load local storage IPsec policy on the computer",
		Src: "Windows",
	},
	"5473": {
		Msg: "PAStore Engine loaded directory storage IPsec policy on the computer",
		Src: "Windows",
	},
	"5474": {
		Msg: "PAStore Engine failed to load directory storage IPsec policy on the computer",
		Src: "Windows",
	},
	"5477": {
		Msg: "PAStore Engine failed to add quick mode filter",
		Src: "Windows",
	},
	"5478": {
		Msg: "IPsec Services has started successfully",
		Src: "Windows",
	},
	"5479": {
		Msg: "IPsec Services has been shut down successfully",
		Src: "Windows",
	},
	"5480": {
		Msg: "IPsec Services failed to get the complete list of network interfaces on the computer",
		Src: "Windows",
	},
	"5483": {
		Msg: "IPsec Services failed to initialize RPC server. IPsec Services could not be started",
		Src: "Windows",
	},
	"5484": {
		Msg: "IPsec Services has experienced a critical failure and has been shut down",
		Src: "Windows",
	},
	"5485": {
		Msg: "IPsec Services failed to process some IPsec filters on a plug-and-play event for network interfaces",
		Src: "Windows",
	},
	"5632": {
		Msg: "A request was made to authenticate to a wireless network",
		Src: "Windows",
	},
	"5633": {
		Msg: "A request was made to authenticate to a wired network",
		Src: "Windows",
	},
	"5712": {
		Msg: "A Remote Procedure Call (RPC) was attempted",
		Src: "Windows",
	},
	"5888": {
		Msg: "An object in the COM+ Catalog was modified",
		Src: "Windows",
	},
	"5889": {
		Msg: "An object was deleted from the COM+ Catalog",
		Src: "Windows",
	},
	"5890": {
		Msg: "An object was added to the COM+ Catalog",
		Src: "Windows",
	},
	"6144": {
		Msg: "Security policy in the group policy objects has been applied successfully",
		Src: "Windows",
	},
	"6145": {
		Msg: "One or more errors occured while processing security policy in the group policy objects",
		Src: "Windows",
	},
	"6272": {
		Msg: "Network Policy Server granted access to a user",
		Src: "Windows",
	},
	"6273": {
		Msg: "Network Policy Server denied access to a user",
		Src: "Windows",
	},
	"6274": {
		Msg: "Network Policy Server discarded the request for a user",
		Src: "Windows",
	},
	"6275": {
		Msg: "Network Policy Server discarded the accounting request for a user",
		Src: "Windows",
	},
	"6276": {
		Msg: "Network Policy Server quarantined a user",
		Src: "Windows",
	},
	"6277": {
		Msg: "Network Policy Server granted access to a user but put it on probation because the host did not meet the defined health policy",
		Src: "Windows",
	},
	"6278": {
		Msg: "Network Policy Server granted full access to a user because the host met the defined health policy",
		Src: "Windows",
	},
	"6279": {
		Msg: "Network Policy Server locked the user account due to repeated failed authentication attempts",
		Src: "Windows",
	},
	"6280": {
		Msg: "Network Policy Server unlocked the user account",
		Src: "Windows",
	},
	"6281": {
		Msg: "Code Integrity determined that the page hashes of an image file are not valid...",
		Src: "Windows",
	},
	"6400": {
		Msg: "BranchCache: Received an incorrectly formatted response while discovering availability of content.",
		Src: "Windows",
	},
	"6401": {
		Msg: "BranchCache: Received invalid data from a peer. Data discarded.",
		Src: "Windows",
	},
	"6402": {
		Msg: "BranchCache: The message to the hosted cache offering it data is incorrectly formatted.",
		Src: "Windows",
	},
	"6403": {
		Msg: "BranchCache: The hosted cache sent an incorrectly formatted response to the client's message to offer it data.",
		Src: "Windows",
	},
	"6404": {
		Msg: "BranchCache: Hosted cache could not be authenticated using the provisioned SSL certificate.",
		Src: "Windows",
	},
	"6405": {
		Msg: "BranchCache: %2 instance(s) of event id %1 occurred.",
		Src: "Windows",
	},
	"6406": {
		Msg: "%1 registered to Windows Firewall to control filtering for the following:",
		Src: "Windows",
	},
	"6407": {
		Msg: "%1",
		Src: "Windows",
	},
	"6408": {
		Msg: "Registered product %1 failed and Windows Firewall is now controlling the filtering for %2.",
		Src: "Windows",
	},
	"6409": {
		Msg: "BranchCache: A service connection point object could not be parsed",
		Src: "Windows",
	},
	"6410": {
		Msg: "Code integrity determined that a file does not meet the security requirements to load into a process. This could be due to the use of shared sections or other issues",
		Src: "Windows",
	},
	"6416": {
		Msg: "A new external device was recognized by the system.",
		Src: "Windows",
	},
	"6417": {
		Msg: "The FIPS mode crypto selftests succeeded",
		Src: "Windows",
	},
	"6418": {
		Msg: "The FIPS mode crypto selftests failed",
		Src: "Windows",
	},
	"6419": {
		Msg: "A request was made to disable a device",
		Src: "Windows",
	},
	"6420": {
		Msg: "A device was disabled",
		Src: "Windows",
	},
	"6421": {
		Msg: "A request was made to enable a device",
		Src: "Windows",
	},
	"6422": {
		Msg: "A device was enabled",
		Src: "Windows",
	},
	"6423": {
		Msg: "The installation of this device is forbidden by system policy",
		Src: "Windows",
	},
	"6424": {
		Msg: "The installation of this device was allowed, after having previously been forbidden by policy",
		Src: "Windows",
	},
	"8191": {
		Msg: "Highest System-Defined Audit Message Value",
		Src: "Windows",
	},
	"11": {
		Msg: "Site collection audit policy changed",
		Src: "SharePoint",
	},
	"12": {
		Msg: "Audit policy changed",
		Src: "SharePoint",
	},
	"13": {
		Msg: "Document checked in",
		Src: "SharePoint",
	},
	"14": {
		Msg: "Document checked out",
		Src: "SharePoint",
	},
	"15": {
		Msg: "Child object deleted",
		Src: "SharePoint",
	},
	"16": {
		Msg: "Child object moved",
		Src: "SharePoint",
	},
	"17": {
		Msg: "Object copied",
		Src: "SharePoint",
	},
	"18": {
		Msg: "Custom event",
		Src: "SharePoint",
	},
	"19": {
		Msg: "Object deleted",
		Src: "SharePoint",
	},
	"20": {
		Msg: "SharePoint audit logs deleted",
		Src: "SharePoint",
	},
	"21": {
		Msg: "Object moved",
		Src: "SharePoint",
	},
	"22": {
		Msg: "Object profile changed",
		Src: "SharePoint",
	},
	"23": {
		Msg: "SharePoint object structure changed",
		Src: "SharePoint",
	},
	"24": {
		Msg: "Search performed",
		Src: "SharePoint",
	},
	"25": {
		Msg: "SharePoint group created",
		Src: "SharePoint",
	},
	"26": {
		Msg: "SharePoint group deleted",
		Src: "SharePoint",
	},
	"27": {
		Msg: "SharePoint group member added",
		Src: "SharePoint",
	},
	"28": {
		Msg: "SharePoint group member removed",
		Src: "SharePoint",
	},
	"29": {
		Msg: "Unique permissions created",
		Src: "SharePoint",
	},
	"30": {
		Msg: "Unique permissions removed",
		Src: "SharePoint",
	},
	"31": {
		Msg: "Permissions updated",
		Src: "SharePoint",
	},
	"32": {
		Msg: "Permissions removed",
		Src: "SharePoint",
	},
	"33": {
		Msg: "Unique permission levels created",
		Src: "SharePoint",
	},
	"34": {
		Msg: "Permission level created",
		Src: "SharePoint",
	},
	"35": {
		Msg: "Permission level deleted",
		Src: "SharePoint",
	},
	"36": {
		Msg: "Permission level modified",
		Src: "SharePoint",
	},
	"37": {
		Msg: "SharePoint site collection administrator added",
		Src: "SharePoint",
	},
	"38": {
		Msg: "SharePoint site collection administrator removed",
		Src: "SharePoint",
	},
	"39": {
		Msg: "Object restored",
		Src: "SharePoint",
	},
	"40": {
		Msg: "Site collection updated",
		Src: "SharePoint",
	},
	"41": {
		Msg: "Web updated",
		Src: "SharePoint",
	},
	"42": {
		Msg: "Document library updated",
		Src: "SharePoint",
	},
	"43": {
		Msg: "Document updated",
		Src: "SharePoint",
	},
	"44": {
		Msg: "List updated",
		Src: "SharePoint",
	},
	"45": {
		Msg: "List item updated",
		Src: "SharePoint",
	},
	"46": {
		Msg: "Folder updated",
		Src: "SharePoint",
	},
	"47": {
		Msg: "Document viewed",
		Src: "SharePoint",
	},
	"48": {
		Msg: "Document library viewed",
		Src: "SharePoint",
	},
	"49": {
		Msg: "List viewed",
		Src: "SharePoint",
	},
	"50": {
		Msg: "Object viewed",
		Src: "SharePoint",
	},
	"51": {
		Msg: "Workflow accessed",
		Src: "SharePoint",
	},
	"52": {
		Msg: "Information management policy created",
		Src: "SharePoint",
	},
	"53": {
		Msg: "Information management policy changed",
		Src: "SharePoint",
	},
	"54": {
		Msg: "Site collection information management policy created",
		Src: "SharePoint",
	},
	"55": {
		Msg: "Site collection information management policy changed",
		Src: "SharePoint",
	},
	"56": {
		Msg: "Export of objects started",
		Src: "SharePoint",
	},
	"57": {
		Msg: "Export of objects completed",
		Src: "SharePoint",
	},
	"58": {
		Msg: "Import of objects started",
		Src: "SharePoint",
	},
	"59": {
		Msg: "Import of objects completed",
		Src: "SharePoint",
	},
	"60": {
		Msg: "Possible tampering warning",
		Src: "SharePoint",
	},
	"61": {
		Msg: "Retention policy processed",
		Src: "SharePoint",
	},
	"62": {
		Msg: "Document fragment updated",
		Src: "SharePoint",
	},
	"63": {
		Msg: "Content type imported",
		Src: "SharePoint",
	},
	"64": {
		Msg: "Information management policy deleted",
		Src: "SharePoint",
	},
	"65": {
		Msg: "Item declared as a record",
		Src: "SharePoint",
	},
	"66": {
		Msg: "Item undeclared as a record",
		Src: "SharePoint",
	},
	"24000": {
		Msg: "SQL audit event",
		Src: "SQL Server",
	},
	"24001": {
		Msg: "Login succeeded (action_id LGIS)",
		Src: "SQL Server",
	},
	"24002": {
		Msg: "Logout succeeded (action_id LGO)",
		Src: "SQL Server",
	},
	"24003": {
		Msg: "Login failed (action_id LGIF)",
		Src: "SQL Server",
	},
	"24004": {
		Msg: "Change own password succeeded (action_id PWCS; class_type LX)",
		Src: "SQL Server",
	},
	"24005": {
		Msg: "Change own password failed (action_id PWCS; class_type LX)",
		Src: "SQL Server",
	},
	"24006": {
		Msg: "Change password succeeded (action_id PWC class_type LX)",
		Src: "SQL Server",
	},
	"24007": {
		Msg: "Change password failed (action_id PWC class_type LX)",
		Src: "SQL Server",
	},
	"24008": {
		Msg: "Reset own password succeeded (action_id PWRS; class_type LX)",
		Src: "SQL Server",
	},
	"24009": {
		Msg: "Reset own password failed (action_id PWRS; class_type LX)",
		Src: "SQL Server",
	},
	"24010": {
		Msg: "Reset password succeeded (action_id PWR; class_type LX)",
		Src: "SQL Server",
	},
	"24011": {
		Msg: "Reset password failed (action_id PWR; class_type LX)",
		Src: "SQL Server",
	},
	"24012": {
		Msg: "Must change password (action_id PWMC)",
		Src: "SQL Server",
	},
	"24013": {
		Msg: "Account unlocked (action_id PWU)",
		Src: "SQL Server",
	},
	"24014": {
		Msg: "Change application role password succeeded (action_id PWC; class_type AR)",
		Src: "SQL Server",
	},
	"24015": {
		Msg: "Change application role password failed (action_id PWC class_type AR)",
		Src: "SQL Server",
	},
	"24016": {
		Msg: "Add member to server role succeeded (action_id APRL class_type SG)",
		Src: "SQL Server",
	},
	"24017": {
		Msg: "Add member to server role failed (action_id APRL class_type SG)",
		Src: "SQL Server",
	},
	"24018": {
		Msg: "Remove member from server role succeeded (action_id DPRL class_type SG)",
		Src: "SQL Server",
	},
	"24019": {
		Msg: "Remove member from server role failed (action_id DPRL class_type SG)",
		Src: "SQL Server",
	},
	"24020": {
		Msg: "Add member to database role succeeded (action_id APRL class_type RL)",
		Src: "SQL Server",
	},
	"24021": {
		Msg: "Add member to database role failed (action_id APRL class_type RL)",
		Src: "SQL Server",
	},
	"24022": {
		Msg: "Remove member from database role succeeded (action_id DPRL class_type RL)",
		Src: "SQL Server",
	},
	"24023": {
		Msg: "Remove member from database role failed (action_id DPRL class_type RL)",
		Src: "SQL Server",
	},
	"24024": {
		Msg: "Issued database backup command (action_id BA class_type DB)",
		Src: "SQL Server",
	},
	"24025": {
		Msg: "Issued transaction log backup command (action_id BAL)",
		Src: "SQL Server",
	},
	"24026": {
		Msg: "Issued database restore command (action_id RS class_type DB)",
		Src: "SQL Server",
	},
	"24027": {
		Msg: "Issued transaction log restore command (action_id RS class_type DB)",
		Src: "SQL Server",
	},
	"24028": {
		Msg: "Issued database console command (action_id DBCC)",
		Src: "SQL Server",
	},
	"24029": {
		Msg: "Issued a bulk administration command (action_id ADBO)",
		Src: "SQL Server",
	},
	"24030": {
		Msg: "Issued an alter connection command (action_id ALCN)",
		Src: "SQL Server",
	},
	"24031": {
		Msg: "Issued an alter resources command (action_id ALRS)",
		Src: "SQL Server",
	},
	"24032": {
		Msg: "Issued an alter server state command (action_id ALSS)",
		Src: "SQL Server",
	},
	"24033": {
		Msg: "Issued an alter server settings command (action_id ALST)",
		Src: "SQL Server",
	},
	"24034": {
		Msg: "Issued a view server state command (action_id VSST)",
		Src: "SQL Server",
	},
	"24035": {
		Msg: "Issued an external access assembly command (action_id XA)",
		Src: "SQL Server",
	},
	"24036": {
		Msg: "Issued an unsafe assembly command (action_id XU)",
		Src: "SQL Server",
	},
	"24037": {
		Msg: "Issued an alter reOrig governor command (action_id ALRS class_type RG)",
		Src: "SQL Server",
	},
	"24038": {
		Msg: "Issued a database authenticate command (action_id AUTH)",
		Src: "SQL Server",
	},
	"24039": {
		Msg: "Issued a database checkpoint command (action_id CP)",
		Src: "SQL Server",
	},
	"24040": {
		Msg: "Issued a database show plan command (action_id SPLN)",
		Src: "SQL Server",
	},
	"24041": {
		Msg: "Issued a subscribe to query information command (action_id SUQN)",
		Src: "SQL Server",
	},
	"24042": {
		Msg: "Issued a view database state command (action_id VDST)",
		Src: "SQL Server",
	},
	"24043": {
		Msg: "Issued a change server audit command (action_id AL class_type A)",
		Src: "SQL Server",
	},
	"24044": {
		Msg: "Issued a change server audit specification command (action_id AL class_type SA)",
		Src: "SQL Server",
	},
	"24045": {
		Msg: "Issued a change database audit specification command (action_id AL class_type DA)",
		Src: "SQL Server",
	},
	"24046": {
		Msg: "Issued a create server audit command (action_id CR class_type A)",
		Src: "SQL Server",
	},
	"24047": {
		Msg: "Issued a create server audit specification command (action_id CR class_type SA)",
		Src: "SQL Server",
	},
	"24048": {
		Msg: "Issued a create database audit specification command (action_id CR class_type DA)",
		Src: "SQL Server",
	},
	"24049": {
		Msg: "Issued a delete server audit command (action_id DR class_type A)",
		Src: "SQL Server",
	},
	"24050": {
		Msg: "Issued a delete server audit specification command (action_id DR class_type SA)",
		Src: "SQL Server",
	},
	"24051": {
		Msg: "Issued a delete database audit specification command (action_id DR class_type DA)",
		Src: "SQL Server",
	},
	"24052": {
		Msg: "Audit failure (action_id AUSF)",
		Src: "SQL Server",
	},
	"24053": {
		Msg: "Audit session changed (action_id AUSC)",
		Src: "SQL Server",
	},
	"24054": {
		Msg: "Started SQL server (action_id SVSR)",
		Src: "SQL Server",
	},
	"24055": {
		Msg: "Paused SQL server (action_id SVPD)",
		Src: "SQL Server",
	},
	"24056": {
		Msg: "Resumed SQL server (action_id SVCN)",
		Src: "SQL Server",
	},
	"24057": {
		Msg: "Stopped SQL server (action_id SVSD)",
		Src: "SQL Server",
	},
	"24058": {
		Msg: "Issued a create server object command (action_id CR; class_type AG, EP, SD, SE, T)",
		Src: "SQL Server",
	},
	"24059": {
		Msg: "Issued a change server object command (action_id AL; class_type AG, EP, SD, SE, T)",
		Src: "SQL Server",
	},
	"24060": {
		Msg: "Issued a delete server object command (action_id DR; class_type AG, EP, SD, SE, T)",
		Src: "SQL Server",
	},
	"24061": {
		Msg: "Issued a create server setting command (action_id CR class_type SR)",
		Src: "SQL Server",
	},
	"24062": {
		Msg: "Issued a change server setting command (action_id AL class_type SR)",
		Src: "SQL Server",
	},
	"24063": {
		Msg: "Issued a delete server setting command (action_id DR class_type SR)",
		Src: "SQL Server",
	},
	"24064": {
		Msg: "Issued a create server cryptographic provider command (action_id CR class_type CP)",
		Src: "SQL Server",
	},
	"24065": {
		Msg: "Issued a delete server cryptographic provider command (action_id DR class_type CP)",
		Src: "SQL Server",
	},
	"24066": {
		Msg: "Issued a change server cryptographic provider command (action_id AL class_type CP)",
		Src: "SQL Server",
	},
	"24067": {
		Msg: "Issued a create server credential command (action_id CR class_type CD)",
		Src: "SQL Server",
	},
	"24068": {
		Msg: "Issued a delete server credential command (action_id DR class_type CD)",
		Src: "SQL Server",
	},
	"24069": {
		Msg: "Issued a change server credential command (action_id AL class_type CD)",
		Src: "SQL Server",
	},
	"24070": {
		Msg: "Issued a change server master key command (action_id AL class_type MK)",
		Src: "SQL Server",
	},
	"24071": {
		Msg: "Issued a back up server master key command (action_id BA class_type MK)",
		Src: "SQL Server",
	},
	"24072": {
		Msg: "Issued a restore server master key command (action_id RS class_type MK)",
		Src: "SQL Server",
	},
	"24073": {
		Msg: "Issued a map server credential to login command (action_id CMLG)",
		Src: "SQL Server",
	},
	"24074": {
		Msg: "Issued a remove map between server credential and login command (action_id NMLG)",
		Src: "SQL Server",
	},
	"24075": {
		Msg: "Issued a create server principal command (action_id CR class_type LX, SL)",
		Src: "SQL Server",
	},
	"24076": {
		Msg: "Issued a delete server principal command (action_id DR class_type LX, SL)",
		Src: "SQL Server",
	},
	"24077": {
		Msg: "Issued a change server principal credentials command (action_id CCLG)",
		Src: "SQL Server",
	},
	"24078": {
		Msg: "Issued a disable server principal command (action_id LGDA)",
		Src: "SQL Server",
	},
	"24079": {
		Msg: "Issued a change server principal default database command (action_id LGDB)",
		Src: "SQL Server",
	},
	"24080": {
		Msg: "Issued an enable server principal command (action_id LGEA)",
		Src: "SQL Server",
	},
	"24081": {
		Msg: "Issued a change server principal default language command (action_id LGLG)",
		Src: "SQL Server",
	},
	"24082": {
		Msg: "Issued a change server principal password expiration command (action_id PWEX)",
		Src: "SQL Server",
	},
	"24083": {
		Msg: "Issued a change server principal password policy command (action_id PWPL)",
		Src: "SQL Server",
	},
	"24084": {
		Msg: "Issued a change server principal Data command (action_id LGNM)",
		Src: "SQL Server",
	},
	"24085": {
		Msg: "Issued a create database command (action_id CR class_type DB)",
		Src: "SQL Server",
	},
	"24086": {
		Msg: "Issued a change database command (action_id AL class_type DB)",
		Src: "SQL Server",
	},
	"24087": {
		Msg: "Issued a delete database command (action_id DR class_type DB)",
		Src: "SQL Server",
	},
	"24088": {
		Msg: "Issued a create certificate command (action_id CR class_type CR)",
		Src: "SQL Server",
	},
	"24089": {
		Msg: "Issued a change certificate command (action_id AL class_type CR)",
		Src: "SQL Server",
	},
	"24090": {
		Msg: "Issued a delete certificate command (action_id DR class_type CR)",
		Src: "SQL Server",
	},
	"24091": {
		Msg: "Issued a back up certificate command (action_id BA class_type CR)",
		Src: "SQL Server",
	},
	"24092": {
		Msg: "Issued an access certificate command (action_id AS class_type CR)",
		Src: "SQL Server",
	},
	"24093": {
		Msg: "Issued a create asymmetric key command (action_id CR class_type AK)",
		Src: "SQL Server",
	},
	"24094": {
		Msg: "Issued a change asymmetric key command (action_id AL class_type AK)",
		Src: "SQL Server",
	},
	"24095": {
		Msg: "Issued a delete asymmetric key command (action_id DR class_type AK)",
		Src: "SQL Server",
	},
	"24096": {
		Msg: "Issued an access asymmetric key command (action_id AS class_type AK)",
		Src: "SQL Server",
	},
	"24097": {
		Msg: "Issued a create database master key command (action_id CR class_type MK)",
		Src: "SQL Server",
	},
	"24098": {
		Msg: "Issued a change database master key command (action_id AL class_type MK)",
		Src: "SQL Server",
	},
	"24099": {
		Msg: "Issued a delete database master key command (action_id DR class_type MK)",
		Src: "SQL Server",
	},
	"24100": {
		Msg: "Issued a back up database master key command (action_id BA class_type MK)",
		Src: "SQL Server",
	},
	"24101": {
		Msg: "Issued a restore database master key command (action_id RS class_type MK)",
		Src: "SQL Server",
	},
	"24102": {
		Msg: "Issued an open database master key command (action_id OP class_type MK)",
		Src: "SQL Server",
	},
	"24103": {
		Msg: "Issued a create database symmetric key command (action_id CR class_type SK)",
		Src: "SQL Server",
	},
	"24104": {
		Msg: "Issued a change database symmetric key command (action_id AL class_type SK)",
		Src: "SQL Server",
	},
	"24105": {
		Msg: "Issued a delete database symmetric key command (action_id DR class_type SK)",
		Src: "SQL Server",
	},
	"24106": {
		Msg: "Issued a back up database symmetric key command (action_id BA class_type SK)",
		Src: "SQL Server",
	},
	"24107": {
		Msg: "Issued an open database symmetric key command (action_id OP class_type SK)",
		Src: "SQL Server",
	},
	"24108": {
		Msg: "Issued a create database object command (action_id CR)",
		Src: "SQL Server",
	},
	"24109": {
		Msg: "Issued a change database object command (action_id AL)",
		Src: "SQL Server",
	},
	"24110": {
		Msg: "Issued a delete database object command (action_id DR)",
		Src: "SQL Server",
	},
	"24111": {
		Msg: "Issued an access database object command (action_id AS)",
		Src: "SQL Server",
	},
	"24112": {
		Msg: "Issued a create assembly command (action_id CR class_type AS)",
		Src: "SQL Server",
	},
	"24113": {
		Msg: "Issued a change assembly command (action_id AL class_type AS)",
		Src: "SQL Server",
	},
	"24114": {
		Msg: "Issued a delete assembly command (action_id DR class_type AS)",
		Src: "SQL Server",
	},
	"24115": {
		Msg: "Issued a create schema command (action_id CR class_type SC)",
		Src: "SQL Server",
	},
	"24116": {
		Msg: "Issued a change schema command (action_id AL class_type SC)",
		Src: "SQL Server",
	},
	"24117": {
		Msg: "Issued a delete schema command (action_id DR class_type SC)",
		Src: "SQL Server",
	},
	"24118": {
		Msg: "Issued a create database encryption key command (action_id CR class_type DK)",
		Src: "SQL Server",
	},
	"24119": {
		Msg: "Issued a change database encryption key command (action_id AL class_type DK)",
		Src: "SQL Server",
	},
	"24120": {
		Msg: "Issued a delete database encryption key command (action_id DR class_type DK)",
		Src: "SQL Server",
	},
	"24121": {
		Msg: "Issued a create database user command (action_id CR; class_type US)",
		Src: "SQL Server",
	},
	"24122": {
		Msg: "Issued a change database user command (action_id AL; class_type US)",
		Src: "SQL Server",
	},
	"24123": {
		Msg: "Issued a delete database user command (action_id DR; class_type US)",
		Src: "SQL Server",
	},
	"24124": {
		Msg: "Issued a create database role command (action_id CR class_type RL)",
		Src: "SQL Server",
	},
	"24125": {
		Msg: "Issued a change database role command (action_id AL class_type RL)",
		Src: "SQL Server",
	},
	"24126": {
		Msg: "Issued a delete database role command (action_id DR class_type RL)",
		Src: "SQL Server",
	},
	"24127": {
		Msg: "Issued a create application role command (action_id CR class_type AR)",
		Src: "SQL Server",
	},
	"24128": {
		Msg: "Issued a change application role command (action_id AL class_type AR)",
		Src: "SQL Server",
	},
	"24129": {
		Msg: "Issued a delete application role command (action_id DR class_type AR)",
		Src: "SQL Server",
	},
	"24130": {
		Msg: "Issued a change database user login command (action_id USAF)",
		Src: "SQL Server",
	},
	"24131": {
		Msg: "Issued an auto-change database user login command (action_id USLG)",
		Src: "SQL Server",
	},
	"24132": {
		Msg: "Issued a create schema object command (action_id CR class_type D)",
		Src: "SQL Server",
	},
	"24133": {
		Msg: "Issued a change schema object command (action_id AL class_type D)",
		Src: "SQL Server",
	},
	"24134": {
		Msg: "Issued a delete schema object command (action_id DR class_type D)",
		Src: "SQL Server",
	},
	"24135": {
		Msg: "Issued a transfer schema object command (action_id TRO class_type D)",
		Src: "SQL Server",
	},
	"24136": {
		Msg: "Issued a create schema type command (action_id CR class_type TY)",
		Src: "SQL Server",
	},
	"24137": {
		Msg: "Issued a change schema type command (action_id AL class_type TY)",
		Src: "SQL Server",
	},
	"24138": {
		Msg: "Issued a delete schema type command (action_id DR class_type TY)",
		Src: "SQL Server",
	},
	"24139": {
		Msg: "Issued a transfer schema type command (action_id TRO class_type TY)",
		Src: "SQL Server",
	},
	"24140": {
		Msg: "Issued a create XML schema collection command (action_id CR class_type SX)",
		Src: "SQL Server",
	},
	"24141": {
		Msg: "Issued a change XML schema collection command (action_id AL class_type SX)",
		Src: "SQL Server",
	},
	"24142": {
		Msg: "Issued a delete XML schema collection command (action_id DR class_type SX)",
		Src: "SQL Server",
	},
	"24143": {
		Msg: "Issued a transfer XML schema collection command (action_id TRO class_type SX)",
		Src: "SQL Server",
	},
	"24144": {
		Msg: "Issued an impersonate within server scope command (action_id IMP; class_type LX)",
		Src: "SQL Server",
	},
	"24145": {
		Msg: "Issued an impersonate within database scope command (action_id IMP; class_type US)",
		Src: "SQL Server",
	},
	"24146": {
		Msg: "Issued a change server object owner command (action_id TO class_type SG)",
		Src: "SQL Server",
	},
	"24147": {
		Msg: "Issued a change database owner command (action_id TO class_type DB)",
		Src: "SQL Server",
	},
	"24148": {
		Msg: "Issued a change schema owner command (action_id TO class_type SC)",
		Src: "SQL Server",
	},
	"24150": {
		Msg: "Issued a change role owner command (action_id TO class_type RL)",
		Src: "SQL Server",
	},
	"24151": {
		Msg: "Issued a change database object owner command (action_id TO)",
		Src: "SQL Server",
	},
	"24152": {
		Msg: "Issued a change symmetric key owner command (action_id TO class_type SK)",
		Src: "SQL Server",
	},
	"24153": {
		Msg: "Issued a change certificate owner command (action_id TO class_type CR)",
		Src: "SQL Server",
	},
	"24154": {
		Msg: "Issued a change asymmetric key owner command (action_id TO class_type AK)",
		Src: "SQL Server",
	},
	"24155": {
		Msg: "Issued a change schema object owner command (action_id TO class_type OB)",
		Src: "SQL Server",
	},
	"24156": {
		Msg: "Issued a change schema type owner command (action_id TO class_type TY)",
		Src: "SQL Server",
	},
	"24157": {
		Msg: "Issued a change XML schema collection owner command (action_id TO class_type SX)",
		Src: "SQL Server",
	},
	"24158": {
		Msg: "Grant server permissions succeeded (action_id G class_type SR)",
		Src: "SQL Server",
	},
	"24159": {
		Msg: "Grant server permissions failed (action_id G class_type SR)",
		Src: "SQL Server",
	},
	"24160": {
		Msg: "Grant server permissions with grant succeeded (action_id GWG class_type SR)",
		Src: "SQL Server",
	},
	"24161": {
		Msg: "Grant server permissions with grant failed (action_id GWG class_type SR)",
		Src: "SQL Server",
	},
	"24162": {
		Msg: "Deny server permissions succeeded (action_id D class_type SR)",
		Src: "SQL Server",
	},
	"24163": {
		Msg: "Deny server permissions failed (action_id D class_type SR)",
		Src: "SQL Server",
	},
	"24164": {
		Msg: "Deny server permissions with cascade succeeded (action_id DWC class_type SR)",
		Src: "SQL Server",
	},
	"24165": {
		Msg: "Deny server permissions with cascade failed (action_id DWC class_type SR)",
		Src: "SQL Server",
	},
	"24166": {
		Msg: "Revoke server permissions succeeded (action_id R class_type SR)",
		Src: "SQL Server",
	},
	"24167": {
		Msg: "Revoke server permissions failed (action_id R class_type SR)",
		Src: "SQL Server",
	},
	"24168": {
		Msg: "Revoke server permissions with grant succeeded (action_id RWG class_type SR)",
		Src: "SQL Server",
	},
	"24169": {
		Msg: "Revoke server permissions with grant failed (action_id RWG class_type SR)",
		Src: "SQL Server",
	},
	"24170": {
		Msg: "Revoke server permissions with cascade succeeded (action_id RWC class_type SR)",
		Src: "SQL Server",
	},
	"24171": {
		Msg: "Revoke server permissions with cascade failed (action_id RWC class_type SR)",
		Src: "SQL Server",
	},
	"24172": {
		Msg: "Issued grant server object permissions command (action_id G; class_type LX)",
		Src: "SQL Server",
	},
	"24173": {
		Msg: "Issued grant server object permissions with grant command (action_id GWG; class_type LX)",
		Src: "SQL Server",
	},
	"24174": {
		Msg: "Issued deny server object permissions command (action_id D; class_type LX)",
		Src: "SQL Server",
	},
	"24175": {
		Msg: "Issued deny server object permissions with cascade command (action_id DWC; class_type LX)",
		Src: "SQL Server",
	},
	"24176": {
		Msg: "Issued revoke server object permissions command (action_id R; class_type LX)",
		Src: "SQL Server",
	},
	"24177": {
		Msg: "Issued revoke server object permissions with grant command (action_id; RWG class_type LX)",
		Src: "SQL Server",
	},
	"24178": {
		Msg: "Issued revoke server object permissions with cascade command (action_id RWC; class_type LX)",
		Src: "SQL Server",
	},
	"24179": {
		Msg: "Grant database permissions succeeded (action_id G class_type DB)",
		Src: "SQL Server",
	},
	"24180": {
		Msg: "Grant database permissions failed (action_id G class_type DB)",
		Src: "SQL Server",
	},
	"24181": {
		Msg: "Grant database permissions with grant succeeded (action_id GWG class_type DB)",
		Src: "SQL Server",
	},
	"24182": {
		Msg: "Grant database permissions with grant failed (action_id GWG class_type DB)",
		Src: "SQL Server",
	},
	"24183": {
		Msg: "Deny database permissions succeeded (action_id D class_type DB)",
		Src: "SQL Server",
	},
	"24184": {
		Msg: "Deny database permissions failed (action_id D class_type DB)",
		Src: "SQL Server",
	},
	"24185": {
		Msg: "Deny database permissions with cascade succeeded (action_id DWC class_type DB)",
		Src: "SQL Server",
	},
	"24186": {
		Msg: "Deny database permissions with cascade failed (action_id DWC class_type DB)",
		Src: "SQL Server",
	},
	"24187": {
		Msg: "Revoke database permissions succeeded (action_id R class_type DB)",
		Src: "SQL Server",
	},
	"24188": {
		Msg: "Revoke database permissions failed (action_id R class_type DB)",
		Src: "SQL Server",
	},
	"24189": {
		Msg: "Revoke database permissions with grant succeeded (action_id RWG class_type DB)",
		Src: "SQL Server",
	},
	"24190": {
		Msg: "Revoke database permissions with grant failed (action_id RWG class_type DB)",
		Src: "SQL Server",
	},
	"24191": {
		Msg: "Revoke database permissions with cascade succeeded (action_id RWC class_type DB)",
		Src: "SQL Server",
	},
	"24192": {
		Msg: "Revoke database permissions with cascade failed (action_id RWC class_type DB)",
		Src: "SQL Server",
	},
	"24193": {
		Msg: "Issued grant database object permissions command (action_id G class_type US)",
		Src: "SQL Server",
	},
	"24194": {
		Msg: "Issued grant database object permissions with grant command (action_id GWG; class_type US)",
		Src: "SQL Server",
	},
	"24195": {
		Msg: "Issued deny database object permissions command (action_id D; class_type US)",
		Src: "SQL Server",
	},
	"24196": {
		Msg: "Issued deny database object permissions with cascade command (action_id DWC; class_type US)",
		Src: "SQL Server",
	},
	"24197": {
		Msg: "Issued revoke database object permissions command (action_id R; class_type US)",
		Src: "SQL Server",
	},
	"24198": {
		Msg: "Issued revoke database object permissions with grant command (action_id RWG; class_type US)",
		Src: "SQL Server",
	},
	"24199": {
		Msg: "Issued revoke database object permissions with cascade command (action_id RWC; class_type US)",
		Src: "SQL Server",
	},
	"24200": {
		Msg: "Issued grant schema permissions command (action_id G class_type SC)",
		Src: "SQL Server",
	},
	"24201": {
		Msg: "Issued grant schema permissions with grant command (action_id GWG class_type SC)",
		Src: "SQL Server",
	},
	"24202": {
		Msg: "Issued deny schema permissions command (action_id D class_type SC)",
		Src: "SQL Server",
	},
	"24203": {
		Msg: "Issued deny schema permissions with cascade command (action_id DWC class_type SC)",
		Src: "SQL Server",
	},
	"24204": {
		Msg: "Issued revoke schema permissions command (action_id R class_type SC)",
		Src: "SQL Server",
	},
	"24205": {
		Msg: "Issued revoke schema permissions with grant command (action_id RWG class_type SC)",
		Src: "SQL Server",
	},
	"24206": {
		Msg: "Issued revoke schema permissions with cascade command (action_id RWC class_type SC)",
		Src: "SQL Server",
	},
	"24207": {
		Msg: "Issued grant assembly permissions command (action_id G class_type AS)",
		Src: "SQL Server",
	},
	"24208": {
		Msg: "Issued grant assembly permissions with grant command (action_id GWG class_type AS)",
		Src: "SQL Server",
	},
	"24209": {
		Msg: "Issued deny assembly permissions command (action_id D class_type AS)",
		Src: "SQL Server",
	},
	"24210": {
		Msg: "Issued deny assembly permissions with cascade command (action_id DWC class_type AS)",
		Src: "SQL Server",
	},
	"24211": {
		Msg: "Issued revoke assembly permissions command (action_id R class_type AS)",
		Src: "SQL Server",
	},
	"24212": {
		Msg: "Issued revoke assembly permissions with grant command (action_id RWG class_type AS)",
		Src: "SQL Server",
	},
	"24213": {
		Msg: "Issued revoke assembly permissions with cascade command (action_id RWC class_type AS)",
		Src: "SQL Server",
	},
	"24214": {
		Msg: "Issued grant database role permissions command (action_id G class_type RL)",
		Src: "SQL Server",
	},
	"24215": {
		Msg: "Issued grant database role permissions with grant command (action_id GWG class_type RL)",
		Src: "SQL Server",
	},
	"24216": {
		Msg: "Issued deny database role permissions command (action_id D class_type RL)",
		Src: "SQL Server",
	},
	"24217": {
		Msg: "Issued deny database role permissions with cascade command (action_id DWC class_type RL)",
		Src: "SQL Server",
	},
	"24218": {
		Msg: "Issued revoke database role permissions command (action_id R class_type RL)",
		Src: "SQL Server",
	},
	"24219": {
		Msg: "Issued revoke database role permissions with grant command (action_id RWG class_type RL)",
		Src: "SQL Server",
	},
	"24220": {
		Msg: "Issued revoke database role permissions with cascade command (action_id RWC class_type RL)",
		Src: "SQL Server",
	},
	"24221": {
		Msg: "Issued grant application role permissions command (action_id G class_type AR)",
		Src: "SQL Server",
	},
	"24222": {
		Msg: "Issued grant application role permissions with grant command (action_id GWG class_type AR)",
		Src: "SQL Server",
	},
	"24223": {
		Msg: "Issued deny application role permissions command (action_id D class_type AR)",
		Src: "SQL Server",
	},
	"24224": {
		Msg: "Issued deny application role permissions with cascade command (action_id DWC class_type AR)",
		Src: "SQL Server",
	},
	"24225": {
		Msg: "Issued revoke application role permissions command (action_id R class_type AR)",
		Src: "SQL Server",
	},
	"24226": {
		Msg: "Issued revoke application role permissions with grant command (action_id RWG class_type AR)",
		Src: "SQL Server",
	},
	"24227": {
		Msg: "Issued revoke application role permissions with cascade command (action_id RWC class_type AR)",
		Src: "SQL Server",
	},
	"24228": {
		Msg: "Issued grant symmetric key permissions command (action_id G class_type SK)",
		Src: "SQL Server",
	},
	"24229": {
		Msg: "Issued grant symmetric key permissions with grant command (action_id GWG class_type SK)",
		Src: "SQL Server",
	},
	"24230": {
		Msg: "Issued deny symmetric key permissions command (action_id D class_type SK)",
		Src: "SQL Server",
	},
	"24231": {
		Msg: "Issued deny symmetric key permissions with cascade command (action_id DWC class_type SK)",
		Src: "SQL Server",
	},
	"24232": {
		Msg: "Issued revoke symmetric key permissions command (action_id R class_type SK)",
		Src: "SQL Server",
	},
	"24233": {
		Msg: "Issued revoke symmetric key permissions with grant command (action_id RWG class_type SK)",
		Src: "SQL Server",
	},
	"24234": {
		Msg: "Issued revoke symmetric key permissions with cascade command (action_id RWC class_type SK)",
		Src: "SQL Server",
	},
	"24235": {
		Msg: "Issued grant certificate permissions command (action_id G class_type CR)",
		Src: "SQL Server",
	},
	"24236": {
		Msg: "Issued grant certificate permissions with grant command (action_id GWG class_type CR)",
		Src: "SQL Server",
	},
	"24237": {
		Msg: "Issued deny certificate permissions command (action_id D class_type CR)",
		Src: "SQL Server",
	},
	"24238": {
		Msg: "Issued deny certificate permissions with cascade command (action_id DWC class_type CR)",
		Src: "SQL Server",
	},
	"24239": {
		Msg: "Issued revoke certificate permissions command (action_id R class_type CR)",
		Src: "SQL Server",
	},
	"24240": {
		Msg: "Issued revoke certificate permissions with grant command (action_id RWG class_type CR)",
		Src: "SQL Server",
	},
	"24241": {
		Msg: "Issued revoke certificate permissions with cascade command (action_id RWC class_type CR)",
		Src: "SQL Server",
	},
	"24242": {
		Msg: "Issued grant asymmetric key permissions command (action_id G class_type AK)",
		Src: "SQL Server",
	},
	"24243": {
		Msg: "Issued grant asymmetric key permissions with grant command (action_id GWG class_type AK)",
		Src: "SQL Server",
	},
	"24244": {
		Msg: "Issued deny asymmetric key permissions command (action_id D class_type AK)",
		Src: "SQL Server",
	},
	"24245": {
		Msg: "Issued deny asymmetric key permissions with cascade command (action_id DWC class_type AK)",
		Src: "SQL Server",
	},
	"24246": {
		Msg: "Issued revoke asymmetric key permissions command (action_id R class_type AK)",
		Src: "SQL Server",
	},
	"24247": {
		Msg: "Issued revoke asymmetric key permissions with grant command (action_id RWG class_type AK)",
		Src: "SQL Server",
	},
	"24248": {
		Msg: "Issued revoke asymmetric key permissions with cascade command (action_id RWC class_type AK)",
		Src: "SQL Server",
	},
	"24249": {
		Msg: "Issued grant schema object permissions command (action_id G class_type OB)",
		Src: "SQL Server",
	},
	"24250": {
		Msg: "Issued grant schema object permissions with grant command (action_id GWG class_type OB)",
		Src: "SQL Server",
	},
	"24251": {
		Msg: "Issued deny schema object permissions command (action_id D class_type OB)",
		Src: "SQL Server",
	},
	"24252": {
		Msg: "Issued deny schema object permissions with cascade command (action_id DWC class_type OB)",
		Src: "SQL Server",
	},
	"24253": {
		Msg: "Issued revoke schema object permissions command (action_id R class_type OB)",
		Src: "SQL Server",
	},
	"24254": {
		Msg: "Issued revoke schema object permissions with grant command (action_id RWG class_type OB)",
		Src: "SQL Server",
	},
	"24255": {
		Msg: "Issued revoke schema object permissions with cascade command (action_id RWC class_type OB)",
		Src: "SQL Server",
	},
	"24256": {
		Msg: "Issued grant schema type permissions command (action_id G class_type TY)",
		Src: "SQL Server",
	},
	"24257": {
		Msg: "Issued grant schema type permissions with grant command (action_id GWG class_type TY)",
		Src: "SQL Server",
	},
	"24258": {
		Msg: "Issued deny schema type permissions command (action_id D class_type TY)",
		Src: "SQL Server",
	},
	"24259": {
		Msg: "Issued deny schema type permissions with cascade command (action_id DWC class_type TY)",
		Src: "SQL Server",
	},
	"24260": {
		Msg: "Issued revoke schema type permissions command (action_id R class_type TY)",
		Src: "SQL Server",
	},
	"24261": {
		Msg: "Issued revoke schema type permissions with grant command (action_id RWG class_type TY)",
		Src: "SQL Server",
	},
	"24262": {
		Msg: "Issued revoke schema type permissions with cascade command (action_id RWC class_type TY)",
		Src: "SQL Server",
	},
	"24263": {
		Msg: "Issued grant XML schema collection permissions command (action_id G class_type SX)",
		Src: "SQL Server",
	},
	"24264": {
		Msg: "Issued grant XML schema collection permissions with grant command (action_id GWG class_type SX)",
		Src: "SQL Server",
	},
	"24265": {
		Msg: "Issued deny XML schema collection permissions command (action_id D class_type SX)",
		Src: "SQL Server",
	},
	"24266": {
		Msg: "Issued deny XML schema collection permissions with cascade command (action_id DWC class_type SX)",
		Src: "SQL Server",
	},
	"24267": {
		Msg: "Issued revoke XML schema collection permissions command (action_id R class_type SX)",
		Src: "SQL Server",
	},
	"24268": {
		Msg: "Issued revoke XML schema collection permissions with grant command (action_id RWG class_type SX)",
		Src: "SQL Server",
	},
	"24269": {
		Msg: "Issued revoke XML schema collection permissions with cascade command (action_id RWC class_type SX)",
		Src: "SQL Server",
	},
	"24270": {
		Msg: "Issued reference database object permissions command (action_id RF)",
		Src: "SQL Server",
	},
	"24271": {
		Msg: "Issued send service request command (action_id SN)",
		Src: "SQL Server",
	},
	"24272": {
		Msg: "Issued check permissions with schema command (action_id VWCT)",
		Src: "SQL Server",
	},
	"24273": {
		Msg: "Issued use service broker transport security command (action_id LGB)",
		Src: "SQL Server",
	},
	"24274": {
		Msg: "Issued use database mirroring transport security command (action_id LGM)",
		Src: "SQL Server",
	},
	"24275": {
		Msg: "Issued alter trace command (action_id ALTR)",
		Src: "SQL Server",
	},
	"24276": {
		Msg: "Issued start trace command (action_id TASA)",
		Src: "SQL Server",
	},
	"24277": {
		Msg: "Issued stop trace command (action_id TASP)",
		Src: "SQL Server",
	},
	"24278": {
		Msg: "Issued enable trace C2 audit mode command (action_id C2ON)",
		Src: "SQL Server",
	},
	"24279": {
		Msg: "Issued disable trace C2 audit mode command (action_id C2OF)",
		Src: "SQL Server",
	},
	"24280": {
		Msg: "Issued server full-text command (action_id FT)",
		Src: "SQL Server",
	},
	"24281": {
		Msg: "Issued select command (action_id SL)",
		Src: "SQL Server",
	},
	"24282": {
		Msg: "Issued update command (action_id UP)",
		Src: "SQL Server",
	},
	"24283": {
		Msg: "Issued insert command (action_id IN)",
		Src: "SQL Server",
	},
	"24284": {
		Msg: "Issued delete command (action_id DL)",
		Src: "SQL Server",
	},
	"24285": {
		Msg: "Issued execute command (action_id EX)",
		Src: "SQL Server",
	},
	"24286": {
		Msg: "Issued receive command (action_id RC)",
		Src: "SQL Server",
	},
	"24287": {
		Msg: "Issued check references command (action_id RF)",
		Src: "SQL Server",
	},
	"24288": {
		Msg: "Issued a create user-defined server role command (action_id CR class_type SG)",
		Src: "SQL Server",
	},
	"24289": {
		Msg: "Issued a change user-defined server role command (action_id AL class_type SG)",
		Src: "SQL Server",
	},
	"24290": {
		Msg: "Issued a delete user-defined server role command (action_id DR class_type SG)",
		Src: "SQL Server",
	},
	"24291": {
		Msg: "Issued grant user-defined server role permissions command (action_id G class_type SG)",
		Src: "SQL Server",
	},
	"24292": {
		Msg: "Issued grant user-defined server role permissions with grant command (action_id GWG class_type SG)",
		Src: "SQL Server",
	},
	"24293": {
		Msg: "Issued deny user-defined server role permissions command (action_id D class_type SG)",
		Src: "SQL Server",
	},
	"24294": {
		Msg: "Issued deny user-defined server role permissions with cascade command (action_id DWC class_type SG)",
		Src: "SQL Server",
	},
	"24295": {
		Msg: "Issued revoke user-defined server role permissions command (action_id R class_type SG)",
		Src: "SQL Server",
	},
	"24296": {
		Msg: "Issued revoke user-defined server role permissions with grant command (action_id RWG class_type SG)",
		Src: "SQL Server",
	},
	"24297": {
		Msg: "Issued revoke user-defined server role permissions with cascade command (action_id RWC class_type SG)",
		Src: "SQL Server",
	},
	"24298": {
		Msg: "Database login succeeded (action_id DBAS)",
		Src: "SQL Server",
	},
	"24299": {
		Msg: "Database login failed (action_id DBAF)",
		Src: "SQL Server",
	},
	"24300": {
		Msg: "Database logout successful (action_id DAGL)",
		Src: "SQL Server",
	},
	"24301": {
		Msg: "Change password succeeded (action_id PWC; class_type US)",
		Src: "SQL Server",
	},
	"24302": {
		Msg: "Change password failed (action_id PWC; class_type US)",
		Src: "SQL Server",
	},
	"24303": {
		Msg: "Change own password succeeded (action_id PWCS; class_type US)",
		Src: "SQL Server",
	},
	"24304": {
		Msg: "Change own password failed (action_id PWCS; class_type US)",
		Src: "SQL Server",
	},
	"24305": {
		Msg: "Reset own password succeeded (action_id PWRS; class_type US)",
		Src: "SQL Server",
	},
	"24306": {
		Msg: "Reset own password failed (action_id PWRS; class_type US)",
		Src: "SQL Server",
	},
	"24307": {
		Msg: "Reset password succeeded (action_id PWR; class_type US)",
		Src: "SQL Server",
	},
	"24308": {
		Msg: "Reset password failed (action_id PWR; class_type US)",
		Src: "SQL Server",
	},
	"24309": {
		Msg: "Copy password (action_id USTC)",
		Src: "SQL Server",
	},
	"24310": {
		Msg: "User-defined SQL audit event (action_id UDAU)",
		Src: "SQL Server",
	},
	"24311": {
		Msg: "Issued a change database audit command (action_id AL class_type DU)",
		Src: "SQL Server",
	},
	"24312": {
		Msg: "Issued a create database audit command (action_id CR class_type DU)",
		Src: "SQL Server",
	},
	"24313": {
		Msg: "Issued a delete database audit command (action_id DR class_type DU)",
		Src: "SQL Server",
	},
	"24314": {
		Msg: "Issued a begin transaction command (action_id TXBG)",
		Src: "SQL Server",
	},
	"24315": {
		Msg: "Issued a commit transaction command (action_id TXCM)",
		Src: "SQL Server",
	},
	"24316": {
		Msg: "Issued a rollback transaction command (action_id TXRB)",
		Src: "SQL Server",
	},
	"24317": {
		Msg: "Issued a create column master key command (action_id CR; class_type CM)",
		Src: "SQL Server",
	},
	"24318": {
		Msg: "Issued a delete column master key command (action_id DR; class_type CM)",
		Src: "SQL Server",
	},
	"24319": {
		Msg: "A column master key was viewed (action_id VW; class_type CM)",
		Src: "SQL Server",
	},
	"24320": {
		Msg: "Issued a create column encryption key command (action_id CR; class_type CK)",
		Src: "SQL Server",
	},
	"24321": {
		Msg: "Issued a change column encryption key command (action_id AL; class_type CK)",
		Src: "SQL Server",
	},
	"24322": {
		Msg: "Issued a delete column encryption key command (action_id DR; class_type CK)",
		Src: "SQL Server",
	},
	"24323": {
		Msg: "A column encryption key was viewed (action_id VW; class_type CK)",
		Src: "SQL Server",
	},
	"24324": {
		Msg: "Issued a create database credential command (action_id CR; class_type DC)",
		Src: "SQL Server",
	},
	"24325": {
		Msg: "Issued a change database credential command (action_id AL; class_type DC)",
		Src: "SQL Server",
	},
	"24326": {
		Msg: "Issued a delete database credential command (action_id DR; class_type DC)",
		Src: "SQL Server",
	},
	"24327": {
		Msg: "Issued a change database scoped configuration command (action_id AL; class_type DS)",
		Src: "SQL Server",
	},
	"24328": {
		Msg: "Issued a create external data Orig command (action_id CR; class_type ED)",
		Src: "SQL Server",
	},
	"24329": {
		Msg: "Issued a change external data Orig command (action_id AL; class_type ED)",
		Src: "SQL Server",
	},
	"24330": {
		Msg: "Issued a delete external data Orig command (action_id DR; class_type ED)",
		Src: "SQL Server",
	},
	"24331": {
		Msg: "Issued a create external file format command (action_id CR; class_type EF)",
		Src: "SQL Server",
	},
	"24332": {
		Msg: "Issued a delete external file format command (action_id DR; class_type EF)",
		Src: "SQL Server",
	},
	"24333": {
		Msg: "Issued a create external reOrig pool command (action_id CR; class_type ER)",
		Src: "SQL Server",
	},
	"24334": {
		Msg: "Issued a change external reOrig pool command (action_id AL; class_type ER)",
		Src: "SQL Server",
	},
	"24335": {
		Msg: "Issued a delete external reOrig pool command (action_id DR; class_type ER)",
		Src: "SQL Server",
	},
	"24337": {
		Msg: "Global transaction login (action_id LGG)",
		Src: "SQL Server",
	},
	"24338": {
		Msg: "Grant permissions on a database scoped credential succeeded (action_id G; class_type DC)",
		Src: "SQL Server",
	},
	"24339": {
		Msg: "Grant permissions on a database scoped credential failed (action_id G; class_type DC)",
		Src: "SQL Server",
	},
	"24340": {
		Msg: "Grant permissions on a database scoped credential with grant succeeded (action_id GWG; class_type DC)",
		Src: "SQL Server",
	},
	"24341": {
		Msg: "Grant permissions on a database scoped credential with grant failed (action_id GWG; class_type DC)",
		Src: "SQL Server",
	},
	"24342": {
		Msg: "Deny permissions on a database scoped credential succeeded (action_id D; class_type DC)",
		Src: "SQL Server",
	},
	"24343": {
		Msg: "Deny permissions on a database scoped credential failed (action_id D; class_type DC)",
		Src: "SQL Server",
	},
	"24344": {
		Msg: "Deny permissions on a database scoped credential with cascade succeeded (action_id DWC; class_type DC)",
		Src: "SQL Server",
	},
	"24345": {
		Msg: "Deny permissions on a database scoped credential with cascade failed (action_id DWC; class_type DC)",
		Src: "SQL Server",
	},
	"24346": {
		Msg: "Revoke permissions on a database scoped credential succeeded (action_id R; class_type DC)",
		Src: "SQL Server",
	},
	"24347": {
		Msg: "Revoke permissions on a database scoped credential failed (action_id R; class_type DC)",
		Src: "SQL Server",
	},
	"24348": {
		Msg: "Revoke permissions with cascade on a database scoped credential succeeded (action_id RWC; class_type DC)",
		Src: "SQL Server",
	},
	"24349": {
		Msg: "Issued a change assembly owner command (action_id TO class_type AS)",
		Src: "SQL Server",
	},
	"24350": {
		Msg: "Revoke permissions with cascade on a database scoped credential failed (action_id RWC; class_type DC)",
		Src: "SQL Server",
	},
	"24351": {
		Msg: "Revoke permissions with grant on a database scoped credential succeeded (action_id RWG; class_type DC)",
		Src: "SQL Server",
	},
	"24352": {
		Msg: "Revoke permissions with grant on a database scoped credential failed (action_id RWG; class_type DC)",
		Src: "SQL Server",
	},
	"24353": {
		Msg: "Issued a change database scoped credential owner command (action_id TO; class_type DC)",
		Src: "SQL Server",
	},
	"24354": {
		Msg: "Issued a create external library command (action_id CR; class_type EL)",
		Src: "SQL Server",
	},
	"24355": {
		Msg: "Issued a change external library command (action_id AL; class_type EL)",
		Src: "SQL Server",
	},
	"24356": {
		Msg: "Issued a drop external library command (action_id DR; class_type EL)",
		Src: "SQL Server",
	},
	"24357": {
		Msg: "Grant permissions on an external library succeeded (action_id G; class_type EL)",
		Src: "SQL Server",
	},
	"24358": {
		Msg: "Grant permissions on an external library failed (action_id G; class_type EL)",
		Src: "SQL Server",
	},
	"24359": {
		Msg: "Grant permissions on an external library with grant succeeded (action_id GWG; class_type EL)",
		Src: "SQL Server",
	},
	"24360": {
		Msg: "Grant permissions on an external library with grant failed (action_id GWG; class_type EL)",
		Src: "SQL Server",
	},
	"24361": {
		Msg: "Deny permissions on an external library succeeded (action_id D; class_type EL)",
		Src: "SQL Server",
	},
	"24362": {
		Msg: "Deny permissions on an external library failed (action_id D; class_type EL)",
		Src: "SQL Server",
	},
	"24363": {
		Msg: "Deny permissions on an external library with cascade succeeded (action_id DWC; class_type EL)",
		Src: "SQL Server",
	},
	"24364": {
		Msg: "Deny permissions on an external library with cascade failed (action_id DWC; class_type EL)",
		Src: "SQL Server",
	},
	"24365": {
		Msg: "Revoke permissions on an external library succeeded (action_id R; class_type EL)",
		Src: "SQL Server",
	},
	"24366": {
		Msg: "Revoke permissions on an external library failed (action_id R; class_type EL)",
		Src: "SQL Server",
	},
	"24367": {
		Msg: "Revoke permissions with cascade on an external library succeeded (action_id RWC; class_type EL)",
		Src: "SQL Server",
	},
	"24368": {
		Msg: "Revoke permissions with cascade on an external library failed (action_id RWC; class_type EL)",
		Src: "SQL Server",
	},
	"24369": {
		Msg: "Revoke permissions with grant on an external library succeeded (action_id RWG; class_type EL)",
		Src: "SQL Server",
	},
	"24370": {
		Msg: "Revoke permissions with grant on an external library failed (action_id RWG; class_type EL)",
		Src: "SQL Server",
	},
	"24371": {
		Msg: "Issued a create database scoped reOrig governor command (action_id CR; class_type DR)",
		Src: "SQL Server",
	},
	"24372": {
		Msg: "Issued a change database scoped reOrig governor command (action_id AL; class_type DR)",
		Src: "SQL Server",
	},
	"24373": {
		Msg: "Issued a drop database scoped reOrig governor command (action_id DR; class_type DR)",
		Src: "SQL Server",
	},
	"24374": {
		Msg: "Issued a database bulk administration command (action_id DABO; class_type DB)",
		Src: "SQL Server",
	},
	"24375": {
		Msg: "Command to change permission failed (action_id D, DWC, G, GWG, R, RWC, RWG; class_type DC, EL)",
		Src: "SQL Server",
	},
	"25000": {
		Msg: "Undocumented Exchange mailbox operation",
		Src: "Exchange",
	},
	"25001": {
		Msg: "Operation Copy - Copy item to another Exchange mailbox folder",
		Src: "Exchange",
	},
	"25002": {
		Msg: "Operation Create - Create item in Exchange mailbox",
		Src: "Exchange",
	},
	"25003": {
		Msg: "Operation FolderBind - Access Exchange mailbox folder",
		Src: "Exchange",
	},
	"25004": {
		Msg: "Operation HardDelete - Delete Exchange mailbox item permanently from Recoverable Items folder",
		Src: "Exchange",
	},
	"25005": {
		Msg: "Operation MessageBind - Access Exchange mailbox item",
		Src: "Exchange",
	},
	"25006": {
		Msg: "Operation Move - Move item to another Exchange mailbox folder",
		Src: "Exchange",
	},
	"25007": {
		Msg: "Operation MoveToDeletedItems - Move Exchange mailbox item to Deleted Items folder",
		Src: "Exchange",
	},
	"25008": {
		Msg: "Operation SendAs - Send message using Send As Exchange mailbox permissions",
		Src: "Exchange",
	},
	"25009": {
		Msg: "Operation SendOnBehalf - Send message using Send on Behalf Exchange mailbox permissions",
		Src: "Exchange",
	},
	"25010": {
		Msg: "Operation SoftDelete - Delete Exchange mailbox item from Deleted Items folder",
		Src: "Exchange",
	},
	"25011": {
		Msg: "Operation Update - Update Exchange mailbox item's properties",
		Src: "Exchange",
	},
	"25012": {
		Msg: "Information Event - Mailbox audit policy applied",
		Src: "Exchange",
	},
	"25100": {
		Msg: "Undocumented Exchange admin operation",
		Src: "Exchange",
	},
	"25101": {
		Msg: "Add-ADPermission Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25102": {
		Msg: "Add-AvailabilityAddressSpace Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25103": {
		Msg: "Add-ContentFilterPhrase Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25104": {
		Msg: "Add-DatabaseAvailabilityGroupServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25105": {
		Msg: "Add-DistributionGroupMember Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25106": {
		Msg: "Add-FederatedDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25107": {
		Msg: "Add-IPAllowListEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25108": {
		Msg: "Add-IPAllowListProvider Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25109": {
		Msg: "Add-IPBlockListEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25110": {
		Msg: "Add-IPBlockListProvider Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25111": {
		Msg: "Add-MailboxDatabaseCopy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25112": {
		Msg: "Add-MailboxFolderPermission Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25113": {
		Msg: "Add-MailboxPermission Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25114": {
		Msg: "Add-ManagementRoleEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25115": {
		Msg: "Add-PublicFolderAdministrativePermission Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25116": {
		Msg: "Add-PublicFolderClientPermission Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25117": {
		Msg: "Add-RoleGroupMember Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25118": {
		Msg: "Clean-MailboxDatabase Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25119": {
		Msg: "Clear-ActiveSyncDevice Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25120": {
		Msg: "Clear-TextMessagingAccount Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25121": {
		Msg: "Compare-TextMessagingVerificationCode Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25122": {
		Msg: "Connect-Mailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25123": {
		Msg: "Disable-AddressListPaging Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25124": {
		Msg: "Disable-CmdletExtensionAgent Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25125": {
		Msg: "Disable-DistributionGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25126": {
		Msg: "Disable-InboxRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25127": {
		Msg: "Disable-JournalRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25128": {
		Msg: "Disable-Mailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25129": {
		Msg: "Disable-MailContact Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25130": {
		Msg: "Disable-MailPublicFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25131": {
		Msg: "Disable-MailUser Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25132": {
		Msg: "Disable-OutlookAnywhere Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25133": {
		Msg: "Disable-OutlookProtectionRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25134": {
		Msg: "Disable-RemoteMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25135": {
		Msg: "Disable-ServiceEmailChannel Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25136": {
		Msg: "Disable-TransportAgent Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25137": {
		Msg: "Disable-TransportRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25138": {
		Msg: "Disable-UMAutoAttendant Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25139": {
		Msg: "Disable-UMIPGateway Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25140": {
		Msg: "Disable-UMMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25141": {
		Msg: "Disable-UMServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25142": {
		Msg: "Dismount-Database Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25143": {
		Msg: "Enable-AddressListPaging Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25144": {
		Msg: "Enable-AntispamUpdates Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25145": {
		Msg: "Enable-CmdletExtensionAgent Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25146": {
		Msg: "Enable-DistributionGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25147": {
		Msg: "Enable-ExchangeCertificate Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25148": {
		Msg: "Enable-InboxRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25149": {
		Msg: "Enable-JournalRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25150": {
		Msg: "Enable-Mailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25151": {
		Msg: "Enable-MailContact Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25152": {
		Msg: "Enable-MailPublicFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25153": {
		Msg: "Enable-MailUser Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25154": {
		Msg: "Enable-OutlookAnywhere Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25155": {
		Msg: "Enable-OutlookProtectionRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25156": {
		Msg: "Enable-RemoteMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25157": {
		Msg: "Enable-ServiceEmailChannel Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25158": {
		Msg: "Enable-TransportAgent Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25159": {
		Msg: "Enable-TransportRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25160": {
		Msg: "Enable-UMAutoAttendant Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25161": {
		Msg: "Enable-UMIPGateway Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25162": {
		Msg: "Enable-UMMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25163": {
		Msg: "Enable-UMServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25164": {
		Msg: "Export-ActiveSyncLog Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25165": {
		Msg: "Export-AutoDiscoverConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25166": {
		Msg: "Export-ExchangeCertificate Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25167": {
		Msg: "Export-JournalRuleCollection Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25168": {
		Msg: "Export-MailboxDiagnosticLogs Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25169": {
		Msg: "Export-Message Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25170": {
		Msg: "Export-RecipientDataProperty Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25171": {
		Msg: "Export-TransportRuleCollection Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25172": {
		Msg: "Export-UMCallDataRecord Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25173": {
		Msg: "Export-UMPrompt Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25174": {
		Msg: "Import-ExchangeCertificate Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25175": {
		Msg: "Import-JournalRuleCollection Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25176": {
		Msg: "Import-RecipientDataProperty Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25177": {
		Msg: "Import-TransportRuleCollection Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25178": {
		Msg: "Import-UMPrompt Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25179": {
		Msg: "Install-TransportAgent Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25180": {
		Msg: "Mount-Database Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25181": {
		Msg: "Move-ActiveMailboxDatabase Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25182": {
		Msg: "Move-AddressList Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25183": {
		Msg: "Move-DatabasePath Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25184": {
		Msg: "Move-OfflineAddressBook Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25185": {
		Msg: "New-AcceptedDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25186": {
		Msg: "New-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25187": {
		Msg: "New-ActiveSyncMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25188": {
		Msg: "New-ActiveSyncVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25189": {
		Msg: "New-AddressList Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25190": {
		Msg: "New-AdminAuditLogSearch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25191": {
		Msg: "New-AutodiscoverVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25192": {
		Msg: "New-AvailabilityReportOutage Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25193": {
		Msg: "New-ClientAccessArray Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25194": {
		Msg: "New-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25195": {
		Msg: "New-DatabaseAvailabilityGroupNetwork Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25196": {
		Msg: "New-DeliveryAgentConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25197": {
		Msg: "New-DistributionGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25198": {
		Msg: "New-DynamicDistributionGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25199": {
		Msg: "New-EcpVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25200": {
		Msg: "New-EdgeSubscription Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25201": {
		Msg: "New-EdgeSyncServiceConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25202": {
		Msg: "New-EmailAddressPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25203": {
		Msg: "New-ExchangeCertificate Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25204": {
		Msg: "New-FederationTrust Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25205": {
		Msg: "New-ForeignConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25206": {
		Msg: "New-GlobalAddressList Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25207": {
		Msg: "New-InboxRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25208": {
		Msg: "New-JournalRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25209": {
		Msg: "New-Mailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25210": {
		Msg: "New-MailboxAuditLogSearch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25211": {
		Msg: "New-MailboxDatabase Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25212": {
		Msg: "New-MailboxFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25213": {
		Msg: "New-MailboxRepairRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25214": {
		Msg: "New-MailboxRestoreRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25215": {
		Msg: "New-MailContact Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25216": {
		Msg: "New-MailMessage Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25217": {
		Msg: "New-MailUser Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25218": {
		Msg: "New-ManagedContentSettings Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25219": {
		Msg: "New-ManagedFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25220": {
		Msg: "New-ManagedFolderMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25221": {
		Msg: "New-ManagementRole Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25222": {
		Msg: "New-ManagementRoleAssignment Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25223": {
		Msg: "New-ManagementScope Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25224": {
		Msg: "New-MessageClassification Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25225": {
		Msg: "New-MoveRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25226": {
		Msg: "New-OabVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25227": {
		Msg: "New-OfflineAddressBook Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25228": {
		Msg: "New-OrganizationRelationship Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25229": {
		Msg: "New-OutlookProtectionRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25230": {
		Msg: "New-OutlookProvider Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25231": {
		Msg: "New-OwaMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25232": {
		Msg: "New-OwaVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25233": {
		Msg: "New-PublicFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25234": {
		Msg: "New-PublicFolderDatabase Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25235": {
		Msg: "New-PublicFolderDatabaseRepairRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25236": {
		Msg: "New-ReceiveConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25237": {
		Msg: "New-RemoteDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25238": {
		Msg: "New-RemoteMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25239": {
		Msg: "New-RetentionPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25240": {
		Msg: "New-RetentionPolicyTag Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25241": {
		Msg: "New-RoleAssignmentPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25242": {
		Msg: "New-RoleGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25243": {
		Msg: "New-RoutingGroupConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25244": {
		Msg: "New-RpcClientAccess Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25245": {
		Msg: "New-SendConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25246": {
		Msg: "New-SharingPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25247": {
		Msg: "New-SystemMessage Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25248": {
		Msg: "New-ThrottlingPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25249": {
		Msg: "New-TransportRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25250": {
		Msg: "New-UMAutoAttendant Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25251": {
		Msg: "New-UMDialPlan Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25252": {
		Msg: "New-UMHuntGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25253": {
		Msg: "New-UMIPGateway Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25254": {
		Msg: "New-UMMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25255": {
		Msg: "New-WebServicesVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25256": {
		Msg: "New-X400AuthoritativeDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25257": {
		Msg: "Remove-AcceptedDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25258": {
		Msg: "Remove-ActiveSyncDevice Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25259": {
		Msg: "Remove-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25260": {
		Msg: "Remove-ActiveSyncDeviceClass Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25261": {
		Msg: "Remove-ActiveSyncMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25262": {
		Msg: "Remove-ActiveSyncVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25263": {
		Msg: "Remove-AddressList Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25264": {
		Msg: "Remove-ADPermission Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25265": {
		Msg: "Remove-AutodiscoverVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25266": {
		Msg: "Remove-AvailabilityAddressSpace Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25267": {
		Msg: "Remove-AvailabilityReportOutage Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25268": {
		Msg: "Remove-ClientAccessArray Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25269": {
		Msg: "Remove-ContentFilterPhrase Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25270": {
		Msg: "Remove-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25271": {
		Msg: "Remove-DatabaseAvailabilityGroupNetwork Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25272": {
		Msg: "Remove-DatabaseAvailabilityGroupServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25273": {
		Msg: "Remove-DeliveryAgentConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25274": {
		Msg: "Remove-DistributionGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25275": {
		Msg: "Remove-DistributionGroupMember Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25276": {
		Msg: "Remove-DynamicDistributionGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25277": {
		Msg: "Remove-EcpVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25278": {
		Msg: "Remove-EdgeSubscription Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25279": {
		Msg: "Remove-EmailAddressPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25280": {
		Msg: "Remove-ExchangeCertificate Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25281": {
		Msg: "Remove-FederatedDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25282": {
		Msg: "Remove-FederationTrust Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25283": {
		Msg: "Remove-ForeignConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25284": {
		Msg: "Remove-GlobalAddressList Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25285": {
		Msg: "Remove-InboxRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25286": {
		Msg: "Remove-IPAllowListEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25287": {
		Msg: "Remove-IPAllowListProvider Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25288": {
		Msg: "Remove-IPBlockListEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25289": {
		Msg: "Remove-IPBlockListProvider Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25290": {
		Msg: "Remove-JournalRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25291": {
		Msg: "Remove-Mailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25292": {
		Msg: "Remove-MailboxDatabase Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25293": {
		Msg: "Remove-MailboxDatabaseCopy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25294": {
		Msg: "Remove-MailboxFolderPermission Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25295": {
		Msg: "Remove-MailboxPermission Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25296": {
		Msg: "Remove-MailboxRestoreRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25297": {
		Msg: "Remove-MailContact Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25298": {
		Msg: "Remove-MailUser Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25299": {
		Msg: "Remove-ManagedContentSettings Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25300": {
		Msg: "Remove-ManagedFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25301": {
		Msg: "Remove-ManagedFolderMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25302": {
		Msg: "Remove-ManagementRole Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25303": {
		Msg: "Remove-ManagementRoleAssignment Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25304": {
		Msg: "Remove-ManagementRoleEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25305": {
		Msg: "Remove-ManagementScope Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25306": {
		Msg: "Remove-Message Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25307": {
		Msg: "Remove-MessageClassification Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25308": {
		Msg: "Remove-MoveRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25309": {
		Msg: "Remove-OabVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25310": {
		Msg: "Remove-OfflineAddressBook Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25311": {
		Msg: "Remove-OrganizationRelationship Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25312": {
		Msg: "Remove-OutlookProtectionRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25313": {
		Msg: "Remove-OutlookProvider Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25314": {
		Msg: "Remove-OwaMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25315": {
		Msg: "Remove-OwaVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25316": {
		Msg: "Remove-PublicFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25317": {
		Msg: "Remove-PublicFolderAdministrativePermission Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25318": {
		Msg: "Remove-PublicFolderClientPermission Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25319": {
		Msg: "Remove-PublicFolderDatabase Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25320": {
		Msg: "Remove-ReceiveConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25321": {
		Msg: "Remove-RemoteDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25322": {
		Msg: "Remove-RemoteMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25323": {
		Msg: "Remove-RetentionPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25324": {
		Msg: "Remove-RetentionPolicyTag Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25325": {
		Msg: "Remove-RoleAssignmentPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25326": {
		Msg: "Remove-RoleGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25327": {
		Msg: "Remove-RoleGroupMember Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25328": {
		Msg: "Remove-RoutingGroupConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25329": {
		Msg: "Remove-RpcClientAccess Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25330": {
		Msg: "Remove-SendConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25331": {
		Msg: "Remove-SharingPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25332": {
		Msg: "Remove-StoreMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25333": {
		Msg: "Remove-SystemMessage Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25334": {
		Msg: "Remove-ThrottlingPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25335": {
		Msg: "Remove-TransportRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25336": {
		Msg: "Remove-UMAutoAttendant Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25337": {
		Msg: "Remove-UMDialPlan Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25338": {
		Msg: "Remove-UMHuntGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25339": {
		Msg: "Remove-UMIPGateway Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25340": {
		Msg: "Remove-UMMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25341": {
		Msg: "Remove-WebServicesVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25342": {
		Msg: "Remove-X400AuthoritativeDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25343": {
		Msg: "Restore-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25344": {
		Msg: "Restore-DetailsTemplate Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25345": {
		Msg: "Restore-Mailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25346": {
		Msg: "Resume-MailboxDatabaseCopy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25347": {
		Msg: "Resume-MailboxExportRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25348": {
		Msg: "Resume-MailboxRestoreRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25349": {
		Msg: "Resume-Message Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25350": {
		Msg: "Resume-MoveRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25351": {
		Msg: "Resume-PublicFolderReplication Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25352": {
		Msg: "Resume-Queue Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25353": {
		Msg: "Retry-Queue Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25354": {
		Msg: "Send-TextMessagingVerificationCode Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25355": {
		Msg: "Set-AcceptedDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25356": {
		Msg: "Set-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25357": {
		Msg: "Set-ActiveSyncMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25358": {
		Msg: "Set-ActiveSyncOrganizationSettings Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25359": {
		Msg: "Set-ActiveSyncVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25360": {
		Msg: "Set-AddressList Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25361": {
		Msg: "Set-AdminAuditLogConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25362": {
		Msg: "Set-ADServerSettings Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25363": {
		Msg: "Set-ADSite Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25364": {
		Msg: "Set-AdSiteLink Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25365": {
		Msg: "Set-AutodiscoverVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25366": {
		Msg: "Set-AvailabilityConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25367": {
		Msg: "Set-AvailabilityReportOutage Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25368": {
		Msg: "Set-CalendarNotification Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25369": {
		Msg: "Set-CalendarProcessing Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25370": {
		Msg: "Set-CASMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25371": {
		Msg: "Set-ClientAccessArray Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25372": {
		Msg: "Set-ClientAccessServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25373": {
		Msg: "Set-CmdletExtensionAgent Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25374": {
		Msg: "Set-Contact Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25375": {
		Msg: "Set-ContentFilterConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25376": {
		Msg: "Set-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25377": {
		Msg: "Set-DatabaseAvailabilityGroupNetwork Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25378": {
		Msg: "Set-DeliveryAgentConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25379": {
		Msg: "Set-DetailsTemplate Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25380": {
		Msg: "Set-DistributionGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25381": {
		Msg: "Set-DynamicDistributionGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25382": {
		Msg: "Set-EcpVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25383": {
		Msg: "Set-EdgeSyncServiceConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25384": {
		Msg: "Set-EmailAddressPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25385": {
		Msg: "Set-EventLogLevel Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25386": {
		Msg: "Set-ExchangeAssistanceConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25387": {
		Msg: "Set-ExchangeServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25388": {
		Msg: "Set-FederatedOrganizationIdentifier Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25389": {
		Msg: "Set-FederationTrust Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25390": {
		Msg: "Set-ForeignConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25391": {
		Msg: "Set-GlobalAddressList Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25392": {
		Msg: "Set-Group Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25393": {
		Msg: "Set-ImapSettings Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25394": {
		Msg: "Set-InboxRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25395": {
		Msg: "Set-IPAllowListConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25396": {
		Msg: "Set-IPAllowListProvider Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25397": {
		Msg: "Set-IPAllowListProvidersConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25398": {
		Msg: "Set-IPBlockListConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25399": {
		Msg: "Set-IPBlockListProvider Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25400": {
		Msg: "Set-IPBlockListProvidersConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25401": {
		Msg: "Set-IRMConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25402": {
		Msg: "Set-JournalRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25403": {
		Msg: "Set-Mailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25404": {
		Msg: "Set-MailboxAuditBypassAssociation Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25405": {
		Msg: "Set-MailboxAutoReplyConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25406": {
		Msg: "Set-MailboxCalendarConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25407": {
		Msg: "Set-MailboxCalendarFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25408": {
		Msg: "Set-MailboxDatabase Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25409": {
		Msg: "Set-MailboxDatabaseCopy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25410": {
		Msg: "Set-MailboxFolderPermission Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25411": {
		Msg: "Set-MailboxJunkEmailConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25412": {
		Msg: "Set-MailboxMessageConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25413": {
		Msg: "Set-MailboxRegionalConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25414": {
		Msg: "Set-MailboxRestoreRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25415": {
		Msg: "Set-MailboxServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25416": {
		Msg: "Set-MailboxSpellingConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25417": {
		Msg: "Set-MailContact Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25418": {
		Msg: "Set-MailPublicFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25419": {
		Msg: "Set-MailUser Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25420": {
		Msg: "Set-ManagedContentSettings Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25421": {
		Msg: "Set-ManagedFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25422": {
		Msg: "Set-ManagedFolderMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25423": {
		Msg: "Set-ManagementRoleAssignment Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25424": {
		Msg: "Set-ManagementRoleEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25425": {
		Msg: "Set-ManagementScope Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25426": {
		Msg: "Set-MessageClassification Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25427": {
		Msg: "Set-MoveRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25428": {
		Msg: "Set-OabVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25429": {
		Msg: "Set-OfflineAddressBook Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25430": {
		Msg: "Set-OrganizationConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25431": {
		Msg: "Set-OrganizationRelationship Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25432": {
		Msg: "Set-OutlookAnywhere Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25433": {
		Msg: "Set-OutlookProtectionRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25434": {
		Msg: "Set-OutlookProvider Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25435": {
		Msg: "Set-OwaMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25436": {
		Msg: "Set-OwaVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25437": {
		Msg: "Set-PopSettings Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25438": {
		Msg: "Set-PowerShellVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25439": {
		Msg: "Set-PublicFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25440": {
		Msg: "Set-PublicFolderDatabase Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25441": {
		Msg: "Set-ReceiveConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25442": {
		Msg: "Set-RecipientFilterConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25443": {
		Msg: "Set-RemoteDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25444": {
		Msg: "Set-RemoteMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25445": {
		Msg: "Set-ReOrigConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25446": {
		Msg: "Set-RetentionPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25447": {
		Msg: "Set-RetentionPolicyTag Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25448": {
		Msg: "Set-RoleAssignmentPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25449": {
		Msg: "Set-RoleGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25450": {
		Msg: "Set-RoutingGroupConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25451": {
		Msg: "Set-RpcClientAccess Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25452": {
		Msg: "Set-SendConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25453": {
		Msg: "Set-SenderFilterConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25454": {
		Msg: "Set-SenderIdConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25455": {
		Msg: "Set-SenderReputationConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25456": {
		Msg: "Set-SharingPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25457": {
		Msg: "Set-SystemMessage Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25458": {
		Msg: "Set-TextMessagingAccount Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25459": {
		Msg: "Set-ThrottlingPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25460": {
		Msg: "Set-ThrottlingPolicyAssociation Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25461": {
		Msg: "Set-TransportAgent Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25462": {
		Msg: "Set-TransportConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25463": {
		Msg: "Set-TransportRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25464": {
		Msg: "Set-TransportServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25465": {
		Msg: "Set-UMAutoAttendant Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25466": {
		Msg: "Set-UMDialPlan Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25467": {
		Msg: "Set-UMIPGateway Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25468": {
		Msg: "Set-UMMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25469": {
		Msg: "Set-UMMailboxPIN Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25470": {
		Msg: "Set-UMMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25471": {
		Msg: "Set-UmServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25472": {
		Msg: "Set-User Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25473": {
		Msg: "Set-WebServicesVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25474": {
		Msg: "Set-X400AuthoritativeDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25475": {
		Msg: "Start-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25476": {
		Msg: "Start-EdgeSynchronization Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25477": {
		Msg: "Start-ManagedFolderAssistant Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25478": {
		Msg: "Start-RetentionAutoTagLearning Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25479": {
		Msg: "Stop-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25480": {
		Msg: "Stop-ManagedFolderAssistant Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25481": {
		Msg: "Suspend-MailboxDatabaseCopy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25482": {
		Msg: "Suspend-MailboxRestoreRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25483": {
		Msg: "Suspend-Message Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25484": {
		Msg: "Suspend-MoveRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25485": {
		Msg: "Suspend-PublicFolderReplication Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25486": {
		Msg: "Suspend-Queue Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25487": {
		Msg: "Test-ActiveSyncConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25488": {
		Msg: "Test-AssistantHealth Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25489": {
		Msg: "Test-CalendarConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25490": {
		Msg: "Test-EcpConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25491": {
		Msg: "Test-EdgeSynchronization Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25492": {
		Msg: "Test-ExchangeSearch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25493": {
		Msg: "Test-FederationTrust Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25494": {
		Msg: "Test-FederationTrustCertificate Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25495": {
		Msg: "Test-ImapConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25496": {
		Msg: "Test-IPAllowListProvider Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25497": {
		Msg: "Test-IPBlockListProvider Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25498": {
		Msg: "Test-IRMConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25499": {
		Msg: "Test-Mailflow Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25500": {
		Msg: "Test-MAPIConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25501": {
		Msg: "Test-MRSHealth Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25502": {
		Msg: "Test-OrganizationRelationship Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25503": {
		Msg: "Test-OutlookConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25504": {
		Msg: "Test-OutlookWebServices Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25505": {
		Msg: "Test-OwaConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25506": {
		Msg: "Test-PopConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25507": {
		Msg: "Test-PowerShellConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25508": {
		Msg: "Test-ReplicationHealth Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25509": {
		Msg: "Test-SenderId Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25510": {
		Msg: "Test-ServiceHealth Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25511": {
		Msg: "Test-SmtpConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25512": {
		Msg: "Test-SystemHealth Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25513": {
		Msg: "Test-UMConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25514": {
		Msg: "Test-WebServicesConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25515": {
		Msg: "Uninstall-TransportAgent Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25516": {
		Msg: "Update-AddressList Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25517": {
		Msg: "Update-DistributionGroupMember Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25518": {
		Msg: "Update-EmailAddressPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25519": {
		Msg: "Update-FileDistributionService Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25520": {
		Msg: "Update-GlobalAddressList Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25521": {
		Msg: "Update-MailboxDatabaseCopy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25522": {
		Msg: "Update-OfflineAddressBook Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25523": {
		Msg: "Update-PublicFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25524": {
		Msg: "Update-PublicFolderHierarchy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25525": {
		Msg: "Update-Recipient Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25526": {
		Msg: "Update-RoleGroupMember Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25527": {
		Msg: "Update-SafeList Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25528": {
		Msg: "Write-AdminAuditLog Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25529": {
		Msg: "Add-GlobalMonitoringOverride Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25530": {
		Msg: "Add-ResubmitRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25531": {
		Msg: "Add-ServerMonitoringOverride Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25532": {
		Msg: "Clear-MobileDevice Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25533": {
		Msg: "Complete-MigrationBatch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25534": {
		Msg: "Disable-App Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25535": {
		Msg: "Disable-MailboxQuarantine Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25536": {
		Msg: "Disable-UMCallAnsweringRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25537": {
		Msg: "Disable-UMService Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25538": {
		Msg: "Dump-ProvisioningCache Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25539": {
		Msg: "Enable-App Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25540": {
		Msg: "Enable-MailboxQuarantine Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25541": {
		Msg: "Enable-UMCallAnsweringRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25542": {
		Msg: "Enable-UMService Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25543": {
		Msg: "Export-DlpPolicyCollection Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25544": {
		Msg: "Export-MigrationReport Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25545": {
		Msg: "Import-DlpPolicyCollection Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25546": {
		Msg: "Import-DlpPolicyTemplate Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25547": {
		Msg: "Invoke-MonitoringProbe Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25548": {
		Msg: "New-AddressBookPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25549": {
		Msg: "New-App Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25550": {
		Msg: "New-AuthServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25551": {
		Msg: "New-ClassificationRuleCollection Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25552": {
		Msg: "New-DlpPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25553": {
		Msg: "New-HybridConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25554": {
		Msg: "New-MailboxExportRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25555": {
		Msg: "New-MailboxImportRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25556": {
		Msg: "New-MailboxSearch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25557": {
		Msg: "New-MalwareFilterPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25558": {
		Msg: "New-MigrationBatch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25559": {
		Msg: "New-MigrationEndpoint Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25560": {
		Msg: "New-MobileDeviceMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25561": {
		Msg: "New-OnPremisesOrganization Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25562": {
		Msg: "New-PartnerApplication Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25563": {
		Msg: "New-PolicyTipConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25564": {
		Msg: "New-PowerShellVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25565": {
		Msg: "New-PublicFolderMigrationRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25566": {
		Msg: "New-ReOrigPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25567": {
		Msg: "New-SiteMailboxProvisioningPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25568": {
		Msg: "New-SyncMailPublicFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25569": {
		Msg: "New-UMCallAnsweringRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25570": {
		Msg: "New-WorkloadManagementPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25571": {
		Msg: "New-WorkloadPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25572": {
		Msg: "Redirect-Message Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25573": {
		Msg: "Remove-AddressBookPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25574": {
		Msg: "Remove-App Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25575": {
		Msg: "Remove-AuthServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25576": {
		Msg: "Remove-ClassificationRuleCollection Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25577": {
		Msg: "Remove-DlpPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25578": {
		Msg: "Remove-DlpPolicyTemplate Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25579": {
		Msg: "Remove-GlobalMonitoringOverride Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25580": {
		Msg: "Remove-HybridConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25581": {
		Msg: "Remove-LinkedUser Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25582": {
		Msg: "Remove-MailboxExportRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25583": {
		Msg: "Remove-MailboxImportRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25584": {
		Msg: "Remove-MailboxSearch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25585": {
		Msg: "Remove-MalwareFilterPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25586": {
		Msg: "Remove-MalwareFilterRecoveryItem Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25587": {
		Msg: "Remove-MigrationBatch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25588": {
		Msg: "Remove-MigrationEndpoint Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25589": {
		Msg: "Remove-MigrationUser Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25590": {
		Msg: "Remove-MobileDevice Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25591": {
		Msg: "Remove-MobileDeviceMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25592": {
		Msg: "Remove-OnPremisesOrganization Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25593": {
		Msg: "Remove-PartnerApplication Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25594": {
		Msg: "Remove-PolicyTipConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25595": {
		Msg: "Remove-PowerShellVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25596": {
		Msg: "Remove-PublicFolderMigrationRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25597": {
		Msg: "Remove-ReOrigPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25598": {
		Msg: "Remove-ResubmitRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25599": {
		Msg: "Remove-SiteMailboxProvisioningPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25600": {
		Msg: "Remove-UMCallAnsweringRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25601": {
		Msg: "Remove-UserPhoto Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25602": {
		Msg: "Remove-WorkloadManagementPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25603": {
		Msg: "Remove-WorkloadPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25604": {
		Msg: "Reset-ProvisioningCache Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25605": {
		Msg: "Resume-MailboxImportRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25606": {
		Msg: "Resume-MalwareFilterRecoveryItem Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25607": {
		Msg: "Resume-PublicFolderMigrationRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25608": {
		Msg: "Set-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25609": {
		Msg: "Set-AddressBookPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25610": {
		Msg: "Set-App Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25611": {
		Msg: "Set-AuthConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25612": {
		Msg: "Set-AuthServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25613": {
		Msg: "Set-ClassificationRuleCollection Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25614": {
		Msg: "Set-DlpPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25615": {
		Msg: "Set-FrontendTransportService Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25616": {
		Msg: "Set-HybridConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25617": {
		Msg: "Set-HybridMailflow Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25618": {
		Msg: "Set-MailboxExportRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25619": {
		Msg: "Set-MailboxImportRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25620": {
		Msg: "Set-MailboxSearch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25621": {
		Msg: "Set-MailboxTransportService Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25622": {
		Msg: "Set-MalwareFilteringServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25623": {
		Msg: "Set-MalwareFilterPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25624": {
		Msg: "Set-MigrationBatch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25625": {
		Msg: "Set-MigrationConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25626": {
		Msg: "Set-MigrationEndpoint Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25627": {
		Msg: "Set-MobileDeviceMailboxPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25628": {
		Msg: "Set-Notification Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25629": {
		Msg: "Set-OnPremisesOrganization Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25630": {
		Msg: "Set-PartnerApplication Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25631": {
		Msg: "Set-PendingFederatedDomain Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25632": {
		Msg: "Set-PolicyTipConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25633": {
		Msg: "Set-PublicFolderMigrationRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25634": {
		Msg: "Set-ReOrigPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25635": {
		Msg: "Set-ResubmitRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25636": {
		Msg: "Set-RMSTemplate Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25637": {
		Msg: "Set-ServerComponentState Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25638": {
		Msg: "Set-ServerMonitor Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25639": {
		Msg: "Set-SiteMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25640": {
		Msg: "Set-SiteMailboxProvisioningPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25641": {
		Msg: "Set-TransportService Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25642": {
		Msg: "Set-UMCallAnsweringRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25643": {
		Msg: "Set-UMCallRouterSettings Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25644": {
		Msg: "Set-UMService Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25645": {
		Msg: "Set-UserPhoto Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25646": {
		Msg: "Set-WorkloadPolicy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25647": {
		Msg: "Start-MailboxSearch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25648": {
		Msg: "Start-MigrationBatch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25649": {
		Msg: "Stop-MailboxSearch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25650": {
		Msg: "Stop-MigrationBatch Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25651": {
		Msg: "Suspend-MailboxExportRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25652": {
		Msg: "Suspend-MailboxImportRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25653": {
		Msg: "Suspend-PublicFolderMigrationRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25654": {
		Msg: "Test-ArchiveConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25655": {
		Msg: "Test-MigrationServerAvailability Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25656": {
		Msg: "Test-OAuthConnectivity Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25657": {
		Msg: "Test-SiteMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25658": {
		Msg: "Update-HybridConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25659": {
		Msg: "Update-PublicFolderMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25660": {
		Msg: "Update-SiteMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25661": {
		Msg: "Add-AttachmentFilterEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25662": {
		Msg: "Remove-AttachmentFilterEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25663": {
		Msg: "New-AddressRewriteEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25664": {
		Msg: "Remove-AddressRewriteEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25665": {
		Msg: "Set-AddressRewriteEntry Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25666": {
		Msg: "Set-AttachmentFilterListConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25667": {
		Msg: "Set-MailboxSentItemsConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25668": {
		Msg: "Update-MovedMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25669": {
		Msg: "Disable-MalwareFilterRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25670": {
		Msg: "Enable-MalwareFilterRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25671": {
		Msg: "New-MalwareFilterRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25672": {
		Msg: "Remove-MalwareFilterRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25673": {
		Msg: "Set-MalwareFilterRule Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25674": {
		Msg: "Remove-MailboxRepairRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25675": {
		Msg: "Remove-ServerMonitoringOverride Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25676": {
		Msg: "Update-ExchangeHelp Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25677": {
		Msg: "Update-StoreMailboxState Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25678": {
		Msg: "Disable-PushNotificationProxy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25679": {
		Msg: "Enable-PushNotificationProxy Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25680": {
		Msg: "New-PublicFolderMoveRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25681": {
		Msg: "Remove-PublicFolderMoveRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25682": {
		Msg: "Resume-PublicFolderMoveRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25683": {
		Msg: "Set-PublicFolderMoveRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25684": {
		Msg: "Suspend-PublicFolderMoveRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25685": {
		Msg: "Update-DatabaseSchema Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25686": {
		Msg: "Set-SearchDocumentFormat Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25687": {
		Msg: "New-AuthRedirect Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25688": {
		Msg: "New-CompliancePolicySyncNotification Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25689": {
		Msg: "New-ComplianceServiceVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25690": {
		Msg: "New-DatabaseAvailabilityGroupConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25691": {
		Msg: "New-DataClassification Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25692": {
		Msg: "New-Fingerprint Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25693": {
		Msg: "New-IntraOrganizationConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25694": {
		Msg: "New-MailboxDeliveryVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25695": {
		Msg: "New-MapiVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25696": {
		Msg: "New-OutlookServiceVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25697": {
		Msg: "New-RestVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25698": {
		Msg: "New-SearchDocumentFormat Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25699": {
		Msg: "New-SettingOverride Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25700": {
		Msg: "New-SiteMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25701": {
		Msg: "Remove-AuthRedirect Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25702": {
		Msg: "Remove-CompliancePolicySyncNotification Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25703": {
		Msg: "Remove-ComplianceServiceVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25704": {
		Msg: "Remove-DatabaseAvailabilityGroupConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25705": {
		Msg: "Remove-DataClassification Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25706": {
		Msg: "Remove-IntraOrganizationConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25707": {
		Msg: "Remove-MailboxDeliveryVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25708": {
		Msg: "Remove-MapiVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25709": {
		Msg: "Remove-OutlookServiceVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25710": {
		Msg: "Remove-PublicFolderMailboxMigrationRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25711": {
		Msg: "Remove-PushNotificationSubscription Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25712": {
		Msg: "Remove-RestVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25713": {
		Msg: "Remove-SearchDocumentFormat Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25714": {
		Msg: "Remove-SettingOverride Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25715": {
		Msg: "Remove-SyncMailPublicFolder Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25716": {
		Msg: "Resume-PublicFolderMailboxMigrationRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25717": {
		Msg: "Send-MapiSubmitSystemProbe Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25718": {
		Msg: "Set-AuthRedirect Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25719": {
		Msg: "Set-ClientAccessService Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25720": {
		Msg: "Set-Clutter Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25721": {
		Msg: "Set-ComplianceServiceVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25722": {
		Msg: "Set-ConsumerMailbox Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25723": {
		Msg: "Set-DatabaseAvailabilityGroupConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25724": {
		Msg: "Set-DataClassification Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25725": {
		Msg: "Set-IntraOrganizationConnector Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25726": {
		Msg: "Set-LogExportVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25727": {
		Msg: "Set-MailboxDeliveryVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25728": {
		Msg: "Set-MapiVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25729": {
		Msg: "Set-OutlookServiceVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25730": {
		Msg: "Set-PublicFolderMailboxMigrationRequest Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25731": {
		Msg: "Set-RestVirtualDirectory Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25732": {
		Msg: "Set-SettingOverride Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25733": {
		Msg: "Set-SmimeConfig Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25734": {
		Msg: "Set-SubmissionMalwareFilteringServer Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25735": {
		Msg: "Set-UMMailboxConfiguration Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25736": {
		Msg: "Set-UnifiedAuditSetting Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25737": {
		Msg: "Start-AuditAssistant Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25738": {
		Msg: "Start-UMPhoneSession Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25739": {
		Msg: "Stop-UMPhoneSession Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25740": {
		Msg: "Test-DataClassification Exchange cmdlet issued",
		Src: "Exchange",
	},
	"25741": {
		Msg: "Test-TextExtraction Exchange cmdlet issued",
		Src: "Exchange",
	},
}
