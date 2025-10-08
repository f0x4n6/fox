package evtx

import (
	"regexp"
	"time"

	"github.com/cuhsat/fox/internal/pkg/files/extract"
)

var events = map[string]struct {
	Data string
	Orig string
}{
	"1100": {
		Data: "The event logging service has shut down",
		Orig: "Windows",
	},
	"1101": {
		Data: "Audit events have been dropped by the transport.",
		Orig: "Windows",
	},
	"1102": {
		Data: "The audit log was cleared",
		Orig: "Windows",
	},
	"1104": {
		Data: "The security Log is now full",
		Orig: "Windows",
	},
	"1105": {
		Data: "Event log automatic backup",
		Orig: "Windows",
	},
	"1108": {
		Data: "The event logging service encountered an error",
		Orig: "Windows",
	},
	"4608": {
		Data: "Windows is starting up",
		Orig: "Windows",
	},
	"4609": {
		Data: "Windows is shutting down",
		Orig: "Windows",
	},
	"4610": {
		Data: "An authentication package has been loaded by the Local Security Authority",
		Orig: "Windows",
	},
	"4611": {
		Data: "A trusted logon process has been registered with the Local Security Authority",
		Orig: "Windows",
	},
	"4612": {
		Data: "Internal reOrigs allocated for the queuing of audit messages have been exhausted, leading to the loss of some audits.",
		Orig: "Windows",
	},
	"4614": {
		Data: "A notification package has been loaded by the Security Account Manager.",
		Orig: "Windows",
	},
	"4615": {
		Data: "Invalid use of LPC port",
		Orig: "Windows",
	},
	"4616": {
		Data: "The system time was changed.",
		Orig: "Windows",
	},
	"4618": {
		Data: "A monitored security event pattern has occurred",
		Orig: "Windows",
	},
	"4621": {
		Data: "Administrator recovered system from CrashOnAuditFail",
		Orig: "Windows",
	},
	"4622": {
		Data: "A security package has been loaded by the Local Security Authority.",
		Orig: "Windows",
	},
	"4624": {
		Data: "An account was successfully logged on",
		Orig: "Windows",
	},
	"4625": {
		Data: "An account failed to log on",
		Orig: "Windows",
	},
	"4626": {
		Data: "User/Device claims information",
		Orig: "Windows",
	},
	"4627": {
		Data: "Group membership information.",
		Orig: "Windows",
	},
	"4634": {
		Data: "An account was logged off",
		Orig: "Windows",
	},
	"4646": {
		Data: "IKE DoS-prevention mode started",
		Orig: "Windows",
	},
	"4647": {
		Data: "User initiated logoff",
		Orig: "Windows",
	},
	"4648": {
		Data: "A logon was attempted using explicit credentials",
		Orig: "Windows",
	},
	"4649": {
		Data: "A replay attack was detected",
		Orig: "Windows",
	},
	"4650": {
		Data: "An IPsec Main Mode security association was established",
		Orig: "Windows",
	},
	"4651": {
		Data: "An IPsec Main Mode security association was established",
		Orig: "Windows",
	},
	"4652": {
		Data: "An IPsec Main Mode negotiation failed",
		Orig: "Windows",
	},
	"4653": {
		Data: "An IPsec Main Mode negotiation failed",
		Orig: "Windows",
	},
	"4654": {
		Data: "An IPsec Quick Mode negotiation failed",
		Orig: "Windows",
	},
	"4655": {
		Data: "An IPsec Main Mode security association ended",
		Orig: "Windows",
	},
	"4656": {
		Data: "A handle to an object was requested",
		Orig: "Windows",
	},
	"4657": {
		Data: "A registry value was modified",
		Orig: "Windows",
	},
	"4658": {
		Data: "The handle to an object was closed",
		Orig: "Windows",
	},
	"4659": {
		Data: "A handle to an object was requested with intent to delete",
		Orig: "Windows",
	},
	"4660": {
		Data: "An object was deleted",
		Orig: "Windows",
	},
	"4661": {
		Data: "A handle to an object was requested",
		Orig: "Windows",
	},
	"4662": {
		Data: "An operation was performed on an object",
		Orig: "Windows",
	},
	"4663": {
		Data: "An attempt was made to access an object",
		Orig: "Windows",
	},
	"4664": {
		Data: "An attempt was made to create a hard link",
		Orig: "Windows",
	},
	"4665": {
		Data: "An attempt was made to create an application client context.",
		Orig: "Windows",
	},
	"4666": {
		Data: "An application attempted an operation",
		Orig: "Windows",
	},
	"4667": {
		Data: "An application client context was deleted",
		Orig: "Windows",
	},
	"4668": {
		Data: "An application was initialized",
		Orig: "Windows",
	},
	"4670": {
		Data: "Permissions on an object were changed",
		Orig: "Windows",
	},
	"4671": {
		Data: "An application attempted to access a blocked ordinal through the TBS",
		Orig: "Windows",
	},
	"4672": {
		Data: "Special privileges assigned to new logon",
		Orig: "Windows",
	},
	"4673": {
		Data: "A privileged service was called",
		Orig: "Windows",
	},
	"4674": {
		Data: "An operation was attempted on a privileged object",
		Orig: "Windows",
	},
	"4675": {
		Data: "SIDs were filtered",
		Orig: "Windows",
	},
	"4688": {
		Data: "A new process has been created",
		Orig: "Windows",
	},
	"4689": {
		Data: "A process has exited",
		Orig: "Windows",
	},
	"4690": {
		Data: "An attempt was made to duplicate a handle to an object",
		Orig: "Windows",
	},
	"4691": {
		Data: "Indirect access to an object was requested",
		Orig: "Windows",
	},
	"4692": {
		Data: "Backup of data protection master key was attempted",
		Orig: "Windows",
	},
	"4693": {
		Data: "Recovery of data protection master key was attempted",
		Orig: "Windows",
	},
	"4694": {
		Data: "Protection of auditable protected data was attempted",
		Orig: "Windows",
	},
	"4695": {
		Data: "Unprotection of auditable protected data was attempted",
		Orig: "Windows",
	},
	"4696": {
		Data: "A primary token was assigned to process",
		Orig: "Windows",
	},
	"4697": {
		Data: "A service was installed in the system",
		Orig: "Windows",
	},
	"4698": {
		Data: "A scheduled task was created",
		Orig: "Windows",
	},
	"4699": {
		Data: "A scheduled task was deleted",
		Orig: "Windows",
	},
	"4700": {
		Data: "A scheduled task was enabled",
		Orig: "Windows",
	},
	"4701": {
		Data: "A scheduled task was disabled",
		Orig: "Windows",
	},
	"4702": {
		Data: "A scheduled task was updated",
		Orig: "Windows",
	},
	"4703": {
		Data: "A token right was adjusted",
		Orig: "Windows",
	},
	"4704": {
		Data: "A user right was assigned",
		Orig: "Windows",
	},
	"4705": {
		Data: "A user right was removed",
		Orig: "Windows",
	},
	"4706": {
		Data: "A new trust was created to a domain",
		Orig: "Windows",
	},
	"4707": {
		Data: "A trust to a domain was removed",
		Orig: "Windows",
	},
	"4709": {
		Data: "IPsec Services was started",
		Orig: "Windows",
	},
	"4710": {
		Data: "IPsec Services was disabled",
		Orig: "Windows",
	},
	"4711": {
		Data: "PAStore Engine (1%)",
		Orig: "Windows",
	},
	"4712": {
		Data: "IPsec Services encountered a potentially serious failure",
		Orig: "Windows",
	},
	"4713": {
		Data: "Kerberos policy was changed",
		Orig: "Windows",
	},
	"4714": {
		Data: "Encrypted data recovery policy was changed",
		Orig: "Windows",
	},
	"4715": {
		Data: "The audit policy (SACL) on an object was changed",
		Orig: "Windows",
	},
	"4716": {
		Data: "Trusted domain information was modified",
		Orig: "Windows",
	},
	"4717": {
		Data: "System security access was granted to an account",
		Orig: "Windows",
	},
	"4718": {
		Data: "System security access was removed from an account",
		Orig: "Windows",
	},
	"4719": {
		Data: "System audit policy was changed",
		Orig: "Windows",
	},
	"4720": {
		Data: "A user account was created",
		Orig: "Windows",
	},
	"4722": {
		Data: "A user account was enabled",
		Orig: "Windows",
	},
	"4723": {
		Data: "An attempt was made to change an account's password",
		Orig: "Windows",
	},
	"4724": {
		Data: "An attempt was made to reset an accounts password",
		Orig: "Windows",
	},
	"4725": {
		Data: "A user account was disabled",
		Orig: "Windows",
	},
	"4726": {
		Data: "A user account was deleted",
		Orig: "Windows",
	},
	"4727": {
		Data: "A security-enabled global group was created",
		Orig: "Windows",
	},
	"4728": {
		Data: "A member was added to a security-enabled global group",
		Orig: "Windows",
	},
	"4729": {
		Data: "A member was removed from a security-enabled global group",
		Orig: "Windows",
	},
	"4730": {
		Data: "A security-enabled global group was deleted",
		Orig: "Windows",
	},
	"4731": {
		Data: "A security-enabled local group was created",
		Orig: "Windows",
	},
	"4732": {
		Data: "A member was added to a security-enabled local group",
		Orig: "Windows",
	},
	"4733": {
		Data: "A member was removed from a security-enabled local group",
		Orig: "Windows",
	},
	"4734": {
		Data: "A security-enabled local group was deleted",
		Orig: "Windows",
	},
	"4735": {
		Data: "A security-enabled local group was changed",
		Orig: "Windows",
	},
	"4737": {
		Data: "A security-enabled global group was changed",
		Orig: "Windows",
	},
	"4738": {
		Data: "A user account was changed",
		Orig: "Windows",
	},
	"4739": {
		Data: "Domain Policy was changed",
		Orig: "Windows",
	},
	"4740": {
		Data: "A user account was locked out",
		Orig: "Windows",
	},
	"4741": {
		Data: "A computer account was created",
		Orig: "Windows",
	},
	"4742": {
		Data: "A computer account was changed",
		Orig: "Windows",
	},
	"4743": {
		Data: "A computer account was deleted",
		Orig: "Windows",
	},
	"4744": {
		Data: "A security-disabled local group was created",
		Orig: "Windows",
	},
	"4745": {
		Data: "A security-disabled local group was changed",
		Orig: "Windows",
	},
	"4746": {
		Data: "A member was added to a security-disabled local group",
		Orig: "Windows",
	},
	"4747": {
		Data: "A member was removed from a security-disabled local group",
		Orig: "Windows",
	},
	"4748": {
		Data: "A security-disabled local group was deleted",
		Orig: "Windows",
	},
	"4749": {
		Data: "A security-disabled global group was created",
		Orig: "Windows",
	},
	"4750": {
		Data: "A security-disabled global group was changed",
		Orig: "Windows",
	},
	"4751": {
		Data: "A member was added to a security-disabled global group",
		Orig: "Windows",
	},
	"4752": {
		Data: "A member was removed from a security-disabled global group",
		Orig: "Windows",
	},
	"4753": {
		Data: "A security-disabled global group was deleted",
		Orig: "Windows",
	},
	"4754": {
		Data: "A security-enabled universal group was created",
		Orig: "Windows",
	},
	"4755": {
		Data: "A security-enabled universal group was changed",
		Orig: "Windows",
	},
	"4756": {
		Data: "A member was added to a security-enabled universal group",
		Orig: "Windows",
	},
	"4757": {
		Data: "A member was removed from a security-enabled universal group",
		Orig: "Windows",
	},
	"4758": {
		Data: "A security-enabled universal group was deleted",
		Orig: "Windows",
	},
	"4759": {
		Data: "A security-disabled universal group was created",
		Orig: "Windows",
	},
	"4760": {
		Data: "A security-disabled universal group was changed",
		Orig: "Windows",
	},
	"4761": {
		Data: "A member was added to a security-disabled universal group",
		Orig: "Windows",
	},
	"4762": {
		Data: "A member was removed from a security-disabled universal group",
		Orig: "Windows",
	},
	"4763": {
		Data: "A security-disabled universal group was deleted",
		Orig: "Windows",
	},
	"4764": {
		Data: "A groups type was changed",
		Orig: "Windows",
	},
	"4765": {
		Data: "SID History was added to an account",
		Orig: "Windows",
	},
	"4766": {
		Data: "An attempt to add SID History to an account failed",
		Orig: "Windows",
	},
	"4767": {
		Data: "A user account was unlocked",
		Orig: "Windows",
	},
	"4768": {
		Data: "A Kerberos authentication ticket (TGT) was requested",
		Orig: "Windows",
	},
	"4769": {
		Data: "A Kerberos service ticket was requested",
		Orig: "Windows",
	},
	"4770": {
		Data: "A Kerberos service ticket was renewed",
		Orig: "Windows",
	},
	"4771": {
		Data: "Kerberos pre-authentication failed",
		Orig: "Windows",
	},
	"4772": {
		Data: "A Kerberos authentication ticket request failed",
		Orig: "Windows",
	},
	"4773": {
		Data: "A Kerberos service ticket request failed",
		Orig: "Windows",
	},
	"4774": {
		Data: "An account was mapped for logon",
		Orig: "Windows",
	},
	"4775": {
		Data: "An account could not be mapped for logon",
		Orig: "Windows",
	},
	"4776": {
		Data: "The domain controller attempted to validate the credentials for an account",
		Orig: "Windows",
	},
	"4777": {
		Data: "The domain controller failed to validate the credentials for an account",
		Orig: "Windows",
	},
	"4778": {
		Data: "A session was reconnected to a Window Station",
		Orig: "Windows",
	},
	"4779": {
		Data: "A session was disconnected from a Window Station",
		Orig: "Windows",
	},
	"4780": {
		Data: "The ACL was set on accounts which are members of administrators groups",
		Orig: "Windows",
	},
	"4781": {
		Data: "The Data of an account was changed",
		Orig: "Windows",
	},
	"4782": {
		Data: "The password hash an account was accessed",
		Orig: "Windows",
	},
	"4783": {
		Data: "A basic application group was created",
		Orig: "Windows",
	},
	"4784": {
		Data: "A basic application group was changed",
		Orig: "Windows",
	},
	"4785": {
		Data: "A member was added to a basic application group",
		Orig: "Windows",
	},
	"4786": {
		Data: "A member was removed from a basic application group",
		Orig: "Windows",
	},
	"4787": {
		Data: "A non-member was added to a basic application group",
		Orig: "Windows",
	},
	"4788": {
		Data: "A non-member was removed from a basic application group..",
		Orig: "Windows",
	},
	"4789": {
		Data: "A basic application group was deleted",
		Orig: "Windows",
	},
	"4790": {
		Data: "An LDAP query group was created",
		Orig: "Windows",
	},
	"4791": {
		Data: "A basic application group was changed",
		Orig: "Windows",
	},
	"4792": {
		Data: "An LDAP query group was deleted",
		Orig: "Windows",
	},
	"4793": {
		Data: "The Password Policy Checking API was called",
		Orig: "Windows",
	},
	"4794": {
		Data: "An attempt was made to set the Directory Services Restore Mode administrator password",
		Orig: "Windows",
	},
	"4797": {
		Data: "An attempt was made to query the existence of a blank password for an account",
		Orig: "Windows",
	},
	"4798": {
		Data: "A user's local group membership was enumerated.",
		Orig: "Windows",
	},
	"4799": {
		Data: "A security-enabled local group membership was enumerated",
		Orig: "Windows",
	},
	"4800": {
		Data: "The workstation was locked",
		Orig: "Windows",
	},
	"4801": {
		Data: "The workstation was unlocked",
		Orig: "Windows",
	},
	"4802": {
		Data: "The screen saver was invoked",
		Orig: "Windows",
	},
	"4803": {
		Data: "The screen saver was dismissed",
		Orig: "Windows",
	},
	"4816": {
		Data: "RPC detected an integrity violation while decrypting an incoming message",
		Orig: "Windows",
	},
	"4817": {
		Data: "Auditing settings on object were changed.",
		Orig: "Windows",
	},
	"4818": {
		Data: "Proposed Central Access Policy does not grant the same access permissions as the current Central Access Policy",
		Orig: "Windows",
	},
	"4819": {
		Data: "Central Access Policies on the machine have been changed",
		Orig: "Windows",
	},
	"4820": {
		Data: "A Kerberos Ticket-granting-ticket (TGT) was denied because the device does not meet the access control restrictions",
		Orig: "Windows",
	},
	"4821": {
		Data: "A Kerberos service ticket was denied because the user, device, or both does not meet the access control restrictions",
		Orig: "Windows",
	},
	"4822": {
		Data: "NTLM authentication failed because the account was a member of the Protected User group",
		Orig: "Windows",
	},
	"4823": {
		Data: "NTLM authentication failed because access control restrictions are required",
		Orig: "Windows",
	},
	"4824": {
		Data: "Kerberos preauthentication by using DES or RC4 failed because the account was a member of the Protected User group",
		Orig: "Windows",
	},
	"4825": {
		Data: "A user was denied the access to Remote Desktop. By default, users are allowed to connect only if they are members of the Remote Desktop Users group or Administrators group",
		Orig: "Windows",
	},
	"4826": {
		Data: "Boot Configuration Data loaded",
		Orig: "Windows",
	},
	"4830": {
		Data: "SID History was removed from an account",
		Orig: "Windows",
	},
	"4864": {
		Data: "A Dataspace collision was detected",
		Orig: "Windows",
	},
	"4865": {
		Data: "A trusted forest information entry was added",
		Orig: "Windows",
	},
	"4866": {
		Data: "A trusted forest information entry was removed",
		Orig: "Windows",
	},
	"4867": {
		Data: "A trusted forest information entry was modified",
		Orig: "Windows",
	},
	"4868": {
		Data: "The certificate manager denied a pending certificate request",
		Orig: "Windows",
	},
	"4869": {
		Data: "Certificate Services received a resubmitted certificate request",
		Orig: "Windows",
	},
	"4870": {
		Data: "Certificate Services revoked a certificate",
		Orig: "Windows",
	},
	"4871": {
		Data: "Certificate Services received a request to publish the certificate revocation list (CRL)",
		Orig: "Windows",
	},
	"4872": {
		Data: "Certificate Services published the certificate revocation list (CRL)",
		Orig: "Windows",
	},
	"4873": {
		Data: "A certificate request extension changed",
		Orig: "Windows",
	},
	"4874": {
		Data: "One or more certificate request attributes changed.",
		Orig: "Windows",
	},
	"4875": {
		Data: "Certificate Services received a request to shut down",
		Orig: "Windows",
	},
	"4876": {
		Data: "Certificate Services backup started",
		Orig: "Windows",
	},
	"4877": {
		Data: "Certificate Services backup completed",
		Orig: "Windows",
	},
	"4878": {
		Data: "Certificate Services restore started",
		Orig: "Windows",
	},
	"4879": {
		Data: "Certificate Services restore completed",
		Orig: "Windows",
	},
	"4880": {
		Data: "Certificate Services started",
		Orig: "Windows",
	},
	"4881": {
		Data: "Certificate Services stopped",
		Orig: "Windows",
	},
	"4882": {
		Data: "The security permissions for Certificate Services changed",
		Orig: "Windows",
	},
	"4883": {
		Data: "Certificate Services retrieved an archived key",
		Orig: "Windows",
	},
	"4884": {
		Data: "Certificate Services imported a certificate into its database",
		Orig: "Windows",
	},
	"4885": {
		Data: "The audit filter for Certificate Services changed",
		Orig: "Windows",
	},
	"4886": {
		Data: "Certificate Services received a certificate request",
		Orig: "Windows",
	},
	"4887": {
		Data: "Certificate Services approved a certificate request and issued a certificate",
		Orig: "Windows",
	},
	"4888": {
		Data: "Certificate Services denied a certificate request",
		Orig: "Windows",
	},
	"4889": {
		Data: "Certificate Services set the status of a certificate request to pending",
		Orig: "Windows",
	},
	"4890": {
		Data: "The certificate manager settings for Certificate Services changed.",
		Orig: "Windows",
	},
	"4891": {
		Data: "A configuration entry changed in Certificate Services",
		Orig: "Windows",
	},
	"4892": {
		Data: "A property of Certificate Services changed",
		Orig: "Windows",
	},
	"4893": {
		Data: "Certificate Services archived a key",
		Orig: "Windows",
	},
	"4894": {
		Data: "Certificate Services imported and archived a key",
		Orig: "Windows",
	},
	"4895": {
		Data: "Certificate Services published the CA certificate to Active Directory Domain Services",
		Orig: "Windows",
	},
	"4896": {
		Data: "One or more rows have been deleted from the certificate database",
		Orig: "Windows",
	},
	"4897": {
		Data: "Role separation enabled",
		Orig: "Windows",
	},
	"4898": {
		Data: "Certificate Services loaded a template",
		Orig: "Windows",
	},
	"4899": {
		Data: "A Certificate Services template was updated",
		Orig: "Windows",
	},
	"4900": {
		Data: "Certificate Services template security was updated",
		Orig: "Windows",
	},
	"4902": {
		Data: "The Per-user audit policy table was created",
		Orig: "Windows",
	},
	"4904": {
		Data: "An attempt was made to register a security event Orig",
		Orig: "Windows",
	},
	"4905": {
		Data: "An attempt was made to unregister a security event Orig",
		Orig: "Windows",
	},
	"4906": {
		Data: "The CrashOnAuditFail value has changed",
		Orig: "Windows",
	},
	"4907": {
		Data: "Auditing settings on object were changed",
		Orig: "Windows",
	},
	"4908": {
		Data: "Special Groups Logon table modified",
		Orig: "Windows",
	},
	"4909": {
		Data: "The local policy settings for the TBS were changed",
		Orig: "Windows",
	},
	"4910": {
		Data: "The group policy settings for the TBS were changed",
		Orig: "Windows",
	},
	"4911": {
		Data: "ReOrig attributes of the object were changed",
		Orig: "Windows",
	},
	"4912": {
		Data: "Per User Audit Policy was changed",
		Orig: "Windows",
	},
	"4913": {
		Data: "Central Access Policy on the object was changed",
		Orig: "Windows",
	},
	"4928": {
		Data: "An Active Directory replica Orig naming context was established",
		Orig: "Windows",
	},
	"4929": {
		Data: "An Active Directory replica Orig naming context was removed",
		Orig: "Windows",
	},
	"4930": {
		Data: "An Active Directory replica Orig naming context was modified",
		Orig: "Windows",
	},
	"4931": {
		Data: "An Active Directory replica destination naming context was modified",
		Orig: "Windows",
	},
	"4932": {
		Data: "Synchronization of a replica of an Active Directory naming context has begun",
		Orig: "Windows",
	},
	"4933": {
		Data: "Synchronization of a replica of an Active Directory naming context has ended",
		Orig: "Windows",
	},
	"4934": {
		Data: "Attributes of an Active Directory object were replicated",
		Orig: "Windows",
	},
	"4935": {
		Data: "Replication failure begins",
		Orig: "Windows",
	},
	"4936": {
		Data: "Replication failure ends",
		Orig: "Windows",
	},
	"4937": {
		Data: "A lingering object was removed from a replica",
		Orig: "Windows",
	},
	"4944": {
		Data: "The following policy was active when the Windows Firewall started",
		Orig: "Windows",
	},
	"4945": {
		Data: "A rule was listed when the Windows Firewall started",
		Orig: "Windows",
	},
	"4946": {
		Data: "A change has been made to Windows Firewall exception list. A rule was added",
		Orig: "Windows",
	},
	"4947": {
		Data: "A change has been made to Windows Firewall exception list. A rule was modified",
		Orig: "Windows",
	},
	"4948": {
		Data: "A change has been made to Windows Firewall exception list. A rule was deleted",
		Orig: "Windows",
	},
	"4949": {
		Data: "Windows Firewall settings were restored to the default values",
		Orig: "Windows",
	},
	"4950": {
		Data: "A Windows Firewall setting has changed",
		Orig: "Windows",
	},
	"4951": {
		Data: "A rule has been ignored because its major version number was not recognized by Windows Firewall",
		Orig: "Windows",
	},
	"4952": {
		Data: "Parts of a rule have been ignored because its minor version number was not recognized by Windows Firewall",
		Orig: "Windows",
	},
	"4953": {
		Data: "A rule has been ignored by Windows Firewall because it could not parse the rule",
		Orig: "Windows",
	},
	"4954": {
		Data: "Windows Firewall Group Policy settings has changed. The new settings have been applied",
		Orig: "Windows",
	},
	"4956": {
		Data: "Windows Firewall has changed the active profile",
		Orig: "Windows",
	},
	"4957": {
		Data: "Windows Firewall did not apply the following rule",
		Orig: "Windows",
	},
	"4958": {
		Data: "Windows Firewall did not apply the following rule because the rule referred to items not configured on this computer",
		Orig: "Windows",
	},
	"4960": {
		Data: "IPsec dropped an inbound packet that failed an integrity check",
		Orig: "Windows",
	},
	"4961": {
		Data: "IPsec dropped an inbound packet that failed a replay check",
		Orig: "Windows",
	},
	"4962": {
		Data: "IPsec dropped an inbound packet that failed a replay check",
		Orig: "Windows",
	},
	"4963": {
		Data: "IPsec dropped an inbound clear text packet that should have been secured",
		Orig: "Windows",
	},
	"4964": {
		Data: "Special groups have been assigned to a new logon",
		Orig: "Windows",
	},
	"4965": {
		Data: "IPsec received a packet from a remote computer with an incorrect Security Parameter Index (SPI).",
		Orig: "Windows",
	},
	"4976": {
		Data: "During Main Mode negotiation, IPsec received an invalid negotiation packet.",
		Orig: "Windows",
	},
	"4977": {
		Data: "During Quick Mode negotiation, IPsec received an invalid negotiation packet.",
		Orig: "Windows",
	},
	"4978": {
		Data: "During Extended Mode negotiation, IPsec received an invalid negotiation packet.",
		Orig: "Windows",
	},
	"4979": {
		Data: "IPsec Main Mode and Extended Mode security associations were established.",
		Orig: "Windows",
	},
	"4980": {
		Data: "IPsec Main Mode and Extended Mode security associations were established",
		Orig: "Windows",
	},
	"4981": {
		Data: "IPsec Main Mode and Extended Mode security associations were established",
		Orig: "Windows",
	},
	"4982": {
		Data: "IPsec Main Mode and Extended Mode security associations were established",
		Orig: "Windows",
	},
	"4983": {
		Data: "An IPsec Extended Mode negotiation failed",
		Orig: "Windows",
	},
	"4984": {
		Data: "An IPsec Extended Mode negotiation failed",
		Orig: "Windows",
	},
	"4985": {
		Data: "The state of a transaction has changed",
		Orig: "Windows",
	},
	"5024": {
		Data: "The Windows Firewall Service has started successfully",
		Orig: "Windows",
	},
	"5025": {
		Data: "The Windows Firewall Service has been stopped",
		Orig: "Windows",
	},
	"5027": {
		Data: "The Windows Firewall Service was unable to retrieve the security policy from the local storage",
		Orig: "Windows",
	},
	"5028": {
		Data: "The Windows Firewall Service was unable to parse the new security policy.",
		Orig: "Windows",
	},
	"5029": {
		Data: "The Windows Firewall Service failed to initialize the driver",
		Orig: "Windows",
	},
	"5030": {
		Data: "The Windows Firewall Service failed to start",
		Orig: "Windows",
	},
	"5031": {
		Data: "The Windows Firewall Service blocked an application from accepting incoming connections on the network.",
		Orig: "Windows",
	},
	"5032": {
		Data: "Windows Firewall was unable to notify the user that it blocked an application from accepting incoming connections on the network",
		Orig: "Windows",
	},
	"5033": {
		Data: "The Windows Firewall Driver has started successfully",
		Orig: "Windows",
	},
	"5034": {
		Data: "The Windows Firewall Driver has been stopped",
		Orig: "Windows",
	},
	"5035": {
		Data: "The Windows Firewall Driver failed to start",
		Orig: "Windows",
	},
	"5037": {
		Data: "The Windows Firewall Driver detected critical runtime error. Terminating",
		Orig: "Windows",
	},
	"5038": {
		Data: "Code integrity determined that the image hash of a file is not valid",
		Orig: "Windows",
	},
	"5039": {
		Data: "A registry key was virtualized.",
		Orig: "Windows",
	},
	"5040": {
		Data: "A change has been made to IPsec settings. An Authentication Set was added.",
		Orig: "Windows",
	},
	"5041": {
		Data: "A change has been made to IPsec settings. An Authentication Set was modified",
		Orig: "Windows",
	},
	"5042": {
		Data: "A change has been made to IPsec settings. An Authentication Set was deleted",
		Orig: "Windows",
	},
	"5043": {
		Data: "A change has been made to IPsec settings. A Connection Security Rule was added",
		Orig: "Windows",
	},
	"5044": {
		Data: "A change has been made to IPsec settings. A Connection Security Rule was modified",
		Orig: "Windows",
	},
	"5045": {
		Data: "A change has been made to IPsec settings. A Connection Security Rule was deleted",
		Orig: "Windows",
	},
	"5046": {
		Data: "A change has been made to IPsec settings. A Crypto Set was added",
		Orig: "Windows",
	},
	"5047": {
		Data: "A change has been made to IPsec settings. A Crypto Set was modified",
		Orig: "Windows",
	},
	"5048": {
		Data: "A change has been made to IPsec settings. A Crypto Set was deleted",
		Orig: "Windows",
	},
	"5049": {
		Data: "An IPsec Security Association was deleted",
		Orig: "Windows",
	},
	"5050": {
		Data: "An attempt to programmatically disable the Windows Firewall using a call to INetFwProfile.FirewallEnabled(FALSE",
		Orig: "Windows",
	},
	"5051": {
		Data: "A file was virtualized",
		Orig: "Windows",
	},
	"5056": {
		Data: "A cryptographic self test was performed",
		Orig: "Windows",
	},
	"5057": {
		Data: "A cryptographic primitive operation failed",
		Orig: "Windows",
	},
	"5058": {
		Data: "Key file operation",
		Orig: "Windows",
	},
	"5059": {
		Data: "Key migration operation",
		Orig: "Windows",
	},
	"5060": {
		Data: "Verification operation failed",
		Orig: "Windows",
	},
	"5061": {
		Data: "Cryptographic operation",
		Orig: "Windows",
	},
	"5062": {
		Data: "A kernel-mode cryptographic self test was performed",
		Orig: "Windows",
	},
	"5063": {
		Data: "A cryptographic provider operation was attempted",
		Orig: "Windows",
	},
	"5064": {
		Data: "A cryptographic context operation was attempted",
		Orig: "Windows",
	},
	"5065": {
		Data: "A cryptographic context modification was attempted",
		Orig: "Windows",
	},
	"5066": {
		Data: "A cryptographic function operation was attempted",
		Orig: "Windows",
	},
	"5067": {
		Data: "A cryptographic function modification was attempted",
		Orig: "Windows",
	},
	"5068": {
		Data: "A cryptographic function provider operation was attempted",
		Orig: "Windows",
	},
	"5069": {
		Data: "A cryptographic function property operation was attempted",
		Orig: "Windows",
	},
	"5070": {
		Data: "A cryptographic function property operation was attempted",
		Orig: "Windows",
	},
	"5071": {
		Data: "Key access denied by Microsoft key distribution service",
		Orig: "Windows",
	},
	"5120": {
		Data: "OCSP Responder Service Started",
		Orig: "Windows",
	},
	"5121": {
		Data: "OCSP Responder Service Stopped",
		Orig: "Windows",
	},
	"5122": {
		Data: "A Configuration entry changed in the OCSP Responder Service",
		Orig: "Windows",
	},
	"5123": {
		Data: "A configuration entry changed in the OCSP Responder Service",
		Orig: "Windows",
	},
	"5124": {
		Data: "A security setting was updated on OCSP Responder Service",
		Orig: "Windows",
	},
	"5125": {
		Data: "A request was submitted to OCSP Responder Service",
		Orig: "Windows",
	},
	"5126": {
		Data: "Signing Certificate was automatically updated by the OCSP Responder Service",
		Orig: "Windows",
	},
	"5127": {
		Data: "The OCSP Revocation Provider successfully updated the revocation information",
		Orig: "Windows",
	},
	"5136": {
		Data: "A directory service object was modified",
		Orig: "Windows",
	},
	"5137": {
		Data: "A directory service object was created",
		Orig: "Windows",
	},
	"5138": {
		Data: "A directory service object was undeleted",
		Orig: "Windows",
	},
	"5139": {
		Data: "A directory service object was moved",
		Orig: "Windows",
	},
	"5140": {
		Data: "A network share object was accessed",
		Orig: "Windows",
	},
	"5141": {
		Data: "A directory service object was deleted",
		Orig: "Windows",
	},
	"5142": {
		Data: "A network share object was added.",
		Orig: "Windows",
	},
	"5143": {
		Data: "A network share object was modified",
		Orig: "Windows",
	},
	"5144": {
		Data: "A network share object was deleted.",
		Orig: "Windows",
	},
	"5145": {
		Data: "A network share object was checked to see whether client can be granted desired access",
		Orig: "Windows",
	},
	"5146": {
		Data: "The Windows Filtering Platform has blocked a packet",
		Orig: "Windows",
	},
	"5147": {
		Data: "A more restrictive Windows Filtering Platform filter has blocked a packet",
		Orig: "Windows",
	},
	"5148": {
		Data: "The Windows Filtering Platform has detected a DoS attack and entered a defensive mode; packets associated with this attack will be discarded.",
		Orig: "Windows",
	},
	"5149": {
		Data: "The DoS attack has subsided and normal processing is being resumed.",
		Orig: "Windows",
	},
	"5150": {
		Data: "The Windows Filtering Platform has blocked a packet.",
		Orig: "Windows",
	},
	"5151": {
		Data: "A more restrictive Windows Filtering Platform filter has blocked a packet.",
		Orig: "Windows",
	},
	"5152": {
		Data: "The Windows Filtering Platform blocked a packet",
		Orig: "Windows",
	},
	"5153": {
		Data: "A more restrictive Windows Filtering Platform filter has blocked a packet",
		Orig: "Windows",
	},
	"5154": {
		Data: "The Windows Filtering Platform has permitted an application or service to listen on a port for incoming connections",
		Orig: "Windows",
	},
	"5155": {
		Data: "The Windows Filtering Platform has blocked an application or service from listening on a port for incoming connections",
		Orig: "Windows",
	},
	"5156": {
		Data: "The Windows Filtering Platform has allowed a connection",
		Orig: "Windows",
	},
	"5157": {
		Data: "The Windows Filtering Platform has blocked a connection",
		Orig: "Windows",
	},
	"5158": {
		Data: "The Windows Filtering Platform has permitted a bind to a local port",
		Orig: "Windows",
	},
	"5159": {
		Data: "The Windows Filtering Platform has blocked a bind to a local port",
		Orig: "Windows",
	},
	"5168": {
		Data: "Spn check for SMB/SMB2 fails.",
		Orig: "Windows",
	},
	"5169": {
		Data: "A directory service object was modified",
		Orig: "Windows",
	},
	"5170": {
		Data: "A directory service object was modified during a background cleanup task",
		Orig: "Windows",
	},
	"5376": {
		Data: "Credential Manager credentials were backed up",
		Orig: "Windows",
	},
	"5377": {
		Data: "Credential Manager credentials were restored from a backup",
		Orig: "Windows",
	},
	"5378": {
		Data: "The requested credentials delegation was disallowed by policy",
		Orig: "Windows",
	},
	"5440": {
		Data: "The following callout was present when the Windows Filtering Platform Base Filtering Engine started",
		Orig: "Windows",
	},
	"5441": {
		Data: "The following filter was present when the Windows Filtering Platform Base Filtering Engine started",
		Orig: "Windows",
	},
	"5442": {
		Data: "The following provider was present when the Windows Filtering Platform Base Filtering Engine started",
		Orig: "Windows",
	},
	"5443": {
		Data: "The following provider context was present when the Windows Filtering Platform Base Filtering Engine started",
		Orig: "Windows",
	},
	"5444": {
		Data: "The following sub-layer was present when the Windows Filtering Platform Base Filtering Engine started",
		Orig: "Windows",
	},
	"5446": {
		Data: "A Windows Filtering Platform callout has been changed",
		Orig: "Windows",
	},
	"5447": {
		Data: "A Windows Filtering Platform filter has been changed",
		Orig: "Windows",
	},
	"5448": {
		Data: "A Windows Filtering Platform provider has been changed",
		Orig: "Windows",
	},
	"5449": {
		Data: "A Windows Filtering Platform provider context has been changed",
		Orig: "Windows",
	},
	"5450": {
		Data: "A Windows Filtering Platform sub-layer has been changed",
		Orig: "Windows",
	},
	"5451": {
		Data: "An IPsec Quick Mode security association was established",
		Orig: "Windows",
	},
	"5452": {
		Data: "An IPsec Quick Mode security association ended",
		Orig: "Windows",
	},
	"5453": {
		Data: "An IPsec negotiation with a remote computer failed because the IKE and AuthIP IPsec Keying Modules (IKEEXT) service is not started",
		Orig: "Windows",
	},
	"5456": {
		Data: "PAStore Engine applied Active Directory storage IPsec policy on the computer",
		Orig: "Windows",
	},
	"5457": {
		Data: "PAStore Engine failed to apply Active Directory storage IPsec policy on the computer",
		Orig: "Windows",
	},
	"5458": {
		Data: "PAStore Engine applied locally cached copy of Active Directory storage IPsec policy on the computer",
		Orig: "Windows",
	},
	"5459": {
		Data: "PAStore Engine failed to apply locally cached copy of Active Directory storage IPsec policy on the computer",
		Orig: "Windows",
	},
	"5460": {
		Data: "PAStore Engine applied local registry storage IPsec policy on the computer",
		Orig: "Windows",
	},
	"5461": {
		Data: "PAStore Engine failed to apply local registry storage IPsec policy on the computer",
		Orig: "Windows",
	},
	"5462": {
		Data: "PAStore Engine failed to apply some rules of the active IPsec policy on the computer",
		Orig: "Windows",
	},
	"5463": {
		Data: "PAStore Engine polled for changes to the active IPsec policy and detected no changes",
		Orig: "Windows",
	},
	"5464": {
		Data: "PAStore Engine polled for changes to the active IPsec policy, detected changes, and applied them to IPsec Services",
		Orig: "Windows",
	},
	"5465": {
		Data: "PAStore Engine received a control for forced reloading of IPsec policy and processed the control successfully",
		Orig: "Windows",
	},
	"5466": {
		Data: "PAStore Engine polled for changes to the Active Directory IPsec policy, determined that Active Directory cannot be reached, and will use the cached copy of the Active Directory IPsec policy instead",
		Orig: "Windows",
	},
	"5467": {
		Data: "PAStore Engine polled for changes to the Active Directory IPsec policy, determined that Active Directory can be reached, and found no changes to the policy",
		Orig: "Windows",
	},
	"5468": {
		Data: "PAStore Engine polled for changes to the Active Directory IPsec policy, determined that Active Directory can be reached, found changes to the policy, and applied those changes",
		Orig: "Windows",
	},
	"5471": {
		Data: "PAStore Engine loaded local storage IPsec policy on the computer",
		Orig: "Windows",
	},
	"5472": {
		Data: "PAStore Engine failed to load local storage IPsec policy on the computer",
		Orig: "Windows",
	},
	"5473": {
		Data: "PAStore Engine loaded directory storage IPsec policy on the computer",
		Orig: "Windows",
	},
	"5474": {
		Data: "PAStore Engine failed to load directory storage IPsec policy on the computer",
		Orig: "Windows",
	},
	"5477": {
		Data: "PAStore Engine failed to add quick mode filter",
		Orig: "Windows",
	},
	"5478": {
		Data: "IPsec Services has started successfully",
		Orig: "Windows",
	},
	"5479": {
		Data: "IPsec Services has been shut down successfully",
		Orig: "Windows",
	},
	"5480": {
		Data: "IPsec Services failed to get the complete list of network interfaces on the computer",
		Orig: "Windows",
	},
	"5483": {
		Data: "IPsec Services failed to initialize RPC server. IPsec Services could not be started",
		Orig: "Windows",
	},
	"5484": {
		Data: "IPsec Services has experienced a critical failure and has been shut down",
		Orig: "Windows",
	},
	"5485": {
		Data: "IPsec Services failed to process some IPsec filters on a plug-and-play event for network interfaces",
		Orig: "Windows",
	},
	"5632": {
		Data: "A request was made to authenticate to a wireless network",
		Orig: "Windows",
	},
	"5633": {
		Data: "A request was made to authenticate to a wired network",
		Orig: "Windows",
	},
	"5712": {
		Data: "A Remote Procedure Call (RPC) was attempted",
		Orig: "Windows",
	},
	"5888": {
		Data: "An object in the COM+ Catalog was modified",
		Orig: "Windows",
	},
	"5889": {
		Data: "An object was deleted from the COM+ Catalog",
		Orig: "Windows",
	},
	"5890": {
		Data: "An object was added to the COM+ Catalog",
		Orig: "Windows",
	},
	"6144": {
		Data: "Security policy in the group policy objects has been applied successfully",
		Orig: "Windows",
	},
	"6145": {
		Data: "One or more errors occured while processing security policy in the group policy objects",
		Orig: "Windows",
	},
	"6272": {
		Data: "Network Policy Server granted access to a user",
		Orig: "Windows",
	},
	"6273": {
		Data: "Network Policy Server denied access to a user",
		Orig: "Windows",
	},
	"6274": {
		Data: "Network Policy Server discarded the request for a user",
		Orig: "Windows",
	},
	"6275": {
		Data: "Network Policy Server discarded the accounting request for a user",
		Orig: "Windows",
	},
	"6276": {
		Data: "Network Policy Server quarantined a user",
		Orig: "Windows",
	},
	"6277": {
		Data: "Network Policy Server granted access to a user but put it on probation because the host did not meet the defined health policy",
		Orig: "Windows",
	},
	"6278": {
		Data: "Network Policy Server granted full access to a user because the host met the defined health policy",
		Orig: "Windows",
	},
	"6279": {
		Data: "Network Policy Server locked the user account due to repeated failed authentication attempts",
		Orig: "Windows",
	},
	"6280": {
		Data: "Network Policy Server unlocked the user account",
		Orig: "Windows",
	},
	"6281": {
		Data: "Code Integrity determined that the page hashes of an image file are not valid...",
		Orig: "Windows",
	},
	"6400": {
		Data: "BranchCache: Received an incorrectly formatted response while discovering availability of content.",
		Orig: "Windows",
	},
	"6401": {
		Data: "BranchCache: Received invalid data from a peer. Data discarded.",
		Orig: "Windows",
	},
	"6402": {
		Data: "BranchCache: The message to the hosted cache offering it data is incorrectly formatted.",
		Orig: "Windows",
	},
	"6403": {
		Data: "BranchCache: The hosted cache sent an incorrectly formatted response to the client's message to offer it data.",
		Orig: "Windows",
	},
	"6404": {
		Data: "BranchCache: Hosted cache could not be authenticated using the provisioned SSL certificate.",
		Orig: "Windows",
	},
	"6405": {
		Data: "BranchCache: %2 instance(s) of event id %1 occurred.",
		Orig: "Windows",
	},
	"6406": {
		Data: "%1 registered to Windows Firewall to control filtering for the following:",
		Orig: "Windows",
	},
	"6407": {
		Data: "%1",
		Orig: "Windows",
	},
	"6408": {
		Data: "Registered product %1 failed and Windows Firewall is now controlling the filtering for %2.",
		Orig: "Windows",
	},
	"6409": {
		Data: "BranchCache: A service connection point object could not be parsed",
		Orig: "Windows",
	},
	"6410": {
		Data: "Code integrity determined that a file does not meet the security requirements to load into a process. This could be due to the use of shared sections or other issues",
		Orig: "Windows",
	},
	"6416": {
		Data: "A new external device was recognized by the system.",
		Orig: "Windows",
	},
	"6417": {
		Data: "The FIPS mode crypto selftests succeeded",
		Orig: "Windows",
	},
	"6418": {
		Data: "The FIPS mode crypto selftests failed",
		Orig: "Windows",
	},
	"6419": {
		Data: "A request was made to disable a device",
		Orig: "Windows",
	},
	"6420": {
		Data: "A device was disabled",
		Orig: "Windows",
	},
	"6421": {
		Data: "A request was made to enable a device",
		Orig: "Windows",
	},
	"6422": {
		Data: "A device was enabled",
		Orig: "Windows",
	},
	"6423": {
		Data: "The installation of this device is forbidden by system policy",
		Orig: "Windows",
	},
	"6424": {
		Data: "The installation of this device was allowed, after having previously been forbidden by policy",
		Orig: "Windows",
	},
	"8191": {
		Data: "Highest System-Defined Audit Message Value",
		Orig: "Windows",
	},
	"11": {
		Data: "Site collection audit policy changed",
		Orig: "SharePoint",
	},
	"12": {
		Data: "Audit policy changed",
		Orig: "SharePoint",
	},
	"13": {
		Data: "Document checked in",
		Orig: "SharePoint",
	},
	"14": {
		Data: "Document checked out",
		Orig: "SharePoint",
	},
	"15": {
		Data: "Child object deleted",
		Orig: "SharePoint",
	},
	"16": {
		Data: "Child object moved",
		Orig: "SharePoint",
	},
	"17": {
		Data: "Object copied",
		Orig: "SharePoint",
	},
	"18": {
		Data: "Custom event",
		Orig: "SharePoint",
	},
	"19": {
		Data: "Object deleted",
		Orig: "SharePoint",
	},
	"20": {
		Data: "SharePoint audit logs deleted",
		Orig: "SharePoint",
	},
	"21": {
		Data: "Object moved",
		Orig: "SharePoint",
	},
	"22": {
		Data: "Object profile changed",
		Orig: "SharePoint",
	},
	"23": {
		Data: "SharePoint object structure changed",
		Orig: "SharePoint",
	},
	"24": {
		Data: "Search performed",
		Orig: "SharePoint",
	},
	"25": {
		Data: "SharePoint group created",
		Orig: "SharePoint",
	},
	"26": {
		Data: "SharePoint group deleted",
		Orig: "SharePoint",
	},
	"27": {
		Data: "SharePoint group member added",
		Orig: "SharePoint",
	},
	"28": {
		Data: "SharePoint group member removed",
		Orig: "SharePoint",
	},
	"29": {
		Data: "Unique permissions created",
		Orig: "SharePoint",
	},
	"30": {
		Data: "Unique permissions removed",
		Orig: "SharePoint",
	},
	"31": {
		Data: "Permissions updated",
		Orig: "SharePoint",
	},
	"32": {
		Data: "Permissions removed",
		Orig: "SharePoint",
	},
	"33": {
		Data: "Unique permission levels created",
		Orig: "SharePoint",
	},
	"34": {
		Data: "Permission level created",
		Orig: "SharePoint",
	},
	"35": {
		Data: "Permission level deleted",
		Orig: "SharePoint",
	},
	"36": {
		Data: "Permission level modified",
		Orig: "SharePoint",
	},
	"37": {
		Data: "SharePoint site collection administrator added",
		Orig: "SharePoint",
	},
	"38": {
		Data: "SharePoint site collection administrator removed",
		Orig: "SharePoint",
	},
	"39": {
		Data: "Object restored",
		Orig: "SharePoint",
	},
	"40": {
		Data: "Site collection updated",
		Orig: "SharePoint",
	},
	"41": {
		Data: "Web updated",
		Orig: "SharePoint",
	},
	"42": {
		Data: "Document library updated",
		Orig: "SharePoint",
	},
	"43": {
		Data: "Document updated",
		Orig: "SharePoint",
	},
	"44": {
		Data: "List updated",
		Orig: "SharePoint",
	},
	"45": {
		Data: "List item updated",
		Orig: "SharePoint",
	},
	"46": {
		Data: "Folder updated",
		Orig: "SharePoint",
	},
	"47": {
		Data: "Document viewed",
		Orig: "SharePoint",
	},
	"48": {
		Data: "Document library viewed",
		Orig: "SharePoint",
	},
	"49": {
		Data: "List viewed",
		Orig: "SharePoint",
	},
	"50": {
		Data: "Object viewed",
		Orig: "SharePoint",
	},
	"51": {
		Data: "Workflow accessed",
		Orig: "SharePoint",
	},
	"52": {
		Data: "Information management policy created",
		Orig: "SharePoint",
	},
	"53": {
		Data: "Information management policy changed",
		Orig: "SharePoint",
	},
	"54": {
		Data: "Site collection information management policy created",
		Orig: "SharePoint",
	},
	"55": {
		Data: "Site collection information management policy changed",
		Orig: "SharePoint",
	},
	"56": {
		Data: "Export of objects started",
		Orig: "SharePoint",
	},
	"57": {
		Data: "Export of objects completed",
		Orig: "SharePoint",
	},
	"58": {
		Data: "Import of objects started",
		Orig: "SharePoint",
	},
	"59": {
		Data: "Import of objects completed",
		Orig: "SharePoint",
	},
	"60": {
		Data: "Possible tampering warning",
		Orig: "SharePoint",
	},
	"61": {
		Data: "Retention policy processed",
		Orig: "SharePoint",
	},
	"62": {
		Data: "Document fragment updated",
		Orig: "SharePoint",
	},
	"63": {
		Data: "Content type imported",
		Orig: "SharePoint",
	},
	"64": {
		Data: "Information management policy deleted",
		Orig: "SharePoint",
	},
	"65": {
		Data: "Item declared as a record",
		Orig: "SharePoint",
	},
	"66": {
		Data: "Item undeclared as a record",
		Orig: "SharePoint",
	},
	"24000": {
		Data: "SQL audit event",
		Orig: "SQL Server",
	},
	"24001": {
		Data: "Login succeeded (action_id LGIS)",
		Orig: "SQL Server",
	},
	"24002": {
		Data: "Logout succeeded (action_id LGO)",
		Orig: "SQL Server",
	},
	"24003": {
		Data: "Login failed (action_id LGIF)",
		Orig: "SQL Server",
	},
	"24004": {
		Data: "Change own password succeeded (action_id PWCS; class_type LX)",
		Orig: "SQL Server",
	},
	"24005": {
		Data: "Change own password failed (action_id PWCS; class_type LX)",
		Orig: "SQL Server",
	},
	"24006": {
		Data: "Change password succeeded (action_id PWC class_type LX)",
		Orig: "SQL Server",
	},
	"24007": {
		Data: "Change password failed (action_id PWC class_type LX)",
		Orig: "SQL Server",
	},
	"24008": {
		Data: "Reset own password succeeded (action_id PWRS; class_type LX)",
		Orig: "SQL Server",
	},
	"24009": {
		Data: "Reset own password failed (action_id PWRS; class_type LX)",
		Orig: "SQL Server",
	},
	"24010": {
		Data: "Reset password succeeded (action_id PWR; class_type LX)",
		Orig: "SQL Server",
	},
	"24011": {
		Data: "Reset password failed (action_id PWR; class_type LX)",
		Orig: "SQL Server",
	},
	"24012": {
		Data: "Must change password (action_id PWMC)",
		Orig: "SQL Server",
	},
	"24013": {
		Data: "Account unlocked (action_id PWU)",
		Orig: "SQL Server",
	},
	"24014": {
		Data: "Change application role password succeeded (action_id PWC; class_type AR)",
		Orig: "SQL Server",
	},
	"24015": {
		Data: "Change application role password failed (action_id PWC class_type AR)",
		Orig: "SQL Server",
	},
	"24016": {
		Data: "Add member to server role succeeded (action_id APRL class_type SG)",
		Orig: "SQL Server",
	},
	"24017": {
		Data: "Add member to server role failed (action_id APRL class_type SG)",
		Orig: "SQL Server",
	},
	"24018": {
		Data: "Remove member from server role succeeded (action_id DPRL class_type SG)",
		Orig: "SQL Server",
	},
	"24019": {
		Data: "Remove member from server role failed (action_id DPRL class_type SG)",
		Orig: "SQL Server",
	},
	"24020": {
		Data: "Add member to database role succeeded (action_id APRL class_type RL)",
		Orig: "SQL Server",
	},
	"24021": {
		Data: "Add member to database role failed (action_id APRL class_type RL)",
		Orig: "SQL Server",
	},
	"24022": {
		Data: "Remove member from database role succeeded (action_id DPRL class_type RL)",
		Orig: "SQL Server",
	},
	"24023": {
		Data: "Remove member from database role failed (action_id DPRL class_type RL)",
		Orig: "SQL Server",
	},
	"24024": {
		Data: "Issued database backup command (action_id BA class_type DB)",
		Orig: "SQL Server",
	},
	"24025": {
		Data: "Issued transaction log backup command (action_id BAL)",
		Orig: "SQL Server",
	},
	"24026": {
		Data: "Issued database restore command (action_id RS class_type DB)",
		Orig: "SQL Server",
	},
	"24027": {
		Data: "Issued transaction log restore command (action_id RS class_type DB)",
		Orig: "SQL Server",
	},
	"24028": {
		Data: "Issued database console command (action_id DBCC)",
		Orig: "SQL Server",
	},
	"24029": {
		Data: "Issued a bulk administration command (action_id ADBO)",
		Orig: "SQL Server",
	},
	"24030": {
		Data: "Issued an alter connection command (action_id ALCN)",
		Orig: "SQL Server",
	},
	"24031": {
		Data: "Issued an alter reOrigs command (action_id ALRS)",
		Orig: "SQL Server",
	},
	"24032": {
		Data: "Issued an alter server state command (action_id ALSS)",
		Orig: "SQL Server",
	},
	"24033": {
		Data: "Issued an alter server settings command (action_id ALST)",
		Orig: "SQL Server",
	},
	"24034": {
		Data: "Issued a view server state command (action_id VSST)",
		Orig: "SQL Server",
	},
	"24035": {
		Data: "Issued an external access assembly command (action_id XA)",
		Orig: "SQL Server",
	},
	"24036": {
		Data: "Issued an unsafe assembly command (action_id XU)",
		Orig: "SQL Server",
	},
	"24037": {
		Data: "Issued an alter reOrig governor command (action_id ALRS class_type RG)",
		Orig: "SQL Server",
	},
	"24038": {
		Data: "Issued a database authenticate command (action_id AUTH)",
		Orig: "SQL Server",
	},
	"24039": {
		Data: "Issued a database checkpoint command (action_id CP)",
		Orig: "SQL Server",
	},
	"24040": {
		Data: "Issued a database show plan command (action_id SPLN)",
		Orig: "SQL Server",
	},
	"24041": {
		Data: "Issued a subscribe to query information command (action_id SUQN)",
		Orig: "SQL Server",
	},
	"24042": {
		Data: "Issued a view database state command (action_id VDST)",
		Orig: "SQL Server",
	},
	"24043": {
		Data: "Issued a change server audit command (action_id AL class_type A)",
		Orig: "SQL Server",
	},
	"24044": {
		Data: "Issued a change server audit specification command (action_id AL class_type SA)",
		Orig: "SQL Server",
	},
	"24045": {
		Data: "Issued a change database audit specification command (action_id AL class_type DA)",
		Orig: "SQL Server",
	},
	"24046": {
		Data: "Issued a create server audit command (action_id CR class_type A)",
		Orig: "SQL Server",
	},
	"24047": {
		Data: "Issued a create server audit specification command (action_id CR class_type SA)",
		Orig: "SQL Server",
	},
	"24048": {
		Data: "Issued a create database audit specification command (action_id CR class_type DA)",
		Orig: "SQL Server",
	},
	"24049": {
		Data: "Issued a delete server audit command (action_id DR class_type A)",
		Orig: "SQL Server",
	},
	"24050": {
		Data: "Issued a delete server audit specification command (action_id DR class_type SA)",
		Orig: "SQL Server",
	},
	"24051": {
		Data: "Issued a delete database audit specification command (action_id DR class_type DA)",
		Orig: "SQL Server",
	},
	"24052": {
		Data: "Audit failure (action_id AUSF)",
		Orig: "SQL Server",
	},
	"24053": {
		Data: "Audit session changed (action_id AUSC)",
		Orig: "SQL Server",
	},
	"24054": {
		Data: "Started SQL server (action_id SVSR)",
		Orig: "SQL Server",
	},
	"24055": {
		Data: "Paused SQL server (action_id SVPD)",
		Orig: "SQL Server",
	},
	"24056": {
		Data: "Resumed SQL server (action_id SVCN)",
		Orig: "SQL Server",
	},
	"24057": {
		Data: "Stopped SQL server (action_id SVSD)",
		Orig: "SQL Server",
	},
	"24058": {
		Data: "Issued a create server object command (action_id CR; class_type AG, EP, SD, SE, T)",
		Orig: "SQL Server",
	},
	"24059": {
		Data: "Issued a change server object command (action_id AL; class_type AG, EP, SD, SE, T)",
		Orig: "SQL Server",
	},
	"24060": {
		Data: "Issued a delete server object command (action_id DR; class_type AG, EP, SD, SE, T)",
		Orig: "SQL Server",
	},
	"24061": {
		Data: "Issued a create server setting command (action_id CR class_type SR)",
		Orig: "SQL Server",
	},
	"24062": {
		Data: "Issued a change server setting command (action_id AL class_type SR)",
		Orig: "SQL Server",
	},
	"24063": {
		Data: "Issued a delete server setting command (action_id DR class_type SR)",
		Orig: "SQL Server",
	},
	"24064": {
		Data: "Issued a create server cryptographic provider command (action_id CR class_type CP)",
		Orig: "SQL Server",
	},
	"24065": {
		Data: "Issued a delete server cryptographic provider command (action_id DR class_type CP)",
		Orig: "SQL Server",
	},
	"24066": {
		Data: "Issued a change server cryptographic provider command (action_id AL class_type CP)",
		Orig: "SQL Server",
	},
	"24067": {
		Data: "Issued a create server credential command (action_id CR class_type CD)",
		Orig: "SQL Server",
	},
	"24068": {
		Data: "Issued a delete server credential command (action_id DR class_type CD)",
		Orig: "SQL Server",
	},
	"24069": {
		Data: "Issued a change server credential command (action_id AL class_type CD)",
		Orig: "SQL Server",
	},
	"24070": {
		Data: "Issued a change server master key command (action_id AL class_type MK)",
		Orig: "SQL Server",
	},
	"24071": {
		Data: "Issued a back up server master key command (action_id BA class_type MK)",
		Orig: "SQL Server",
	},
	"24072": {
		Data: "Issued a restore server master key command (action_id RS class_type MK)",
		Orig: "SQL Server",
	},
	"24073": {
		Data: "Issued a map server credential to login command (action_id CMLG)",
		Orig: "SQL Server",
	},
	"24074": {
		Data: "Issued a remove map between server credential and login command (action_id NMLG)",
		Orig: "SQL Server",
	},
	"24075": {
		Data: "Issued a create server principal command (action_id CR class_type LX, SL)",
		Orig: "SQL Server",
	},
	"24076": {
		Data: "Issued a delete server principal command (action_id DR class_type LX, SL)",
		Orig: "SQL Server",
	},
	"24077": {
		Data: "Issued a change server principal credentials command (action_id CCLG)",
		Orig: "SQL Server",
	},
	"24078": {
		Data: "Issued a disable server principal command (action_id LGDA)",
		Orig: "SQL Server",
	},
	"24079": {
		Data: "Issued a change server principal default database command (action_id LGDB)",
		Orig: "SQL Server",
	},
	"24080": {
		Data: "Issued an enable server principal command (action_id LGEA)",
		Orig: "SQL Server",
	},
	"24081": {
		Data: "Issued a change server principal default language command (action_id LGLG)",
		Orig: "SQL Server",
	},
	"24082": {
		Data: "Issued a change server principal password expiration command (action_id PWEX)",
		Orig: "SQL Server",
	},
	"24083": {
		Data: "Issued a change server principal password policy command (action_id PWPL)",
		Orig: "SQL Server",
	},
	"24084": {
		Data: "Issued a change server principal Data command (action_id LGNM)",
		Orig: "SQL Server",
	},
	"24085": {
		Data: "Issued a create database command (action_id CR class_type DB)",
		Orig: "SQL Server",
	},
	"24086": {
		Data: "Issued a change database command (action_id AL class_type DB)",
		Orig: "SQL Server",
	},
	"24087": {
		Data: "Issued a delete database command (action_id DR class_type DB)",
		Orig: "SQL Server",
	},
	"24088": {
		Data: "Issued a create certificate command (action_id CR class_type CR)",
		Orig: "SQL Server",
	},
	"24089": {
		Data: "Issued a change certificate command (action_id AL class_type CR)",
		Orig: "SQL Server",
	},
	"24090": {
		Data: "Issued a delete certificate command (action_id DR class_type CR)",
		Orig: "SQL Server",
	},
	"24091": {
		Data: "Issued a back up certificate command (action_id BA class_type CR)",
		Orig: "SQL Server",
	},
	"24092": {
		Data: "Issued an access certificate command (action_id AS class_type CR)",
		Orig: "SQL Server",
	},
	"24093": {
		Data: "Issued a create asymmetric key command (action_id CR class_type AK)",
		Orig: "SQL Server",
	},
	"24094": {
		Data: "Issued a change asymmetric key command (action_id AL class_type AK)",
		Orig: "SQL Server",
	},
	"24095": {
		Data: "Issued a delete asymmetric key command (action_id DR class_type AK)",
		Orig: "SQL Server",
	},
	"24096": {
		Data: "Issued an access asymmetric key command (action_id AS class_type AK)",
		Orig: "SQL Server",
	},
	"24097": {
		Data: "Issued a create database master key command (action_id CR class_type MK)",
		Orig: "SQL Server",
	},
	"24098": {
		Data: "Issued a change database master key command (action_id AL class_type MK)",
		Orig: "SQL Server",
	},
	"24099": {
		Data: "Issued a delete database master key command (action_id DR class_type MK)",
		Orig: "SQL Server",
	},
	"24100": {
		Data: "Issued a back up database master key command (action_id BA class_type MK)",
		Orig: "SQL Server",
	},
	"24101": {
		Data: "Issued a restore database master key command (action_id RS class_type MK)",
		Orig: "SQL Server",
	},
	"24102": {
		Data: "Issued an open database master key command (action_id OP class_type MK)",
		Orig: "SQL Server",
	},
	"24103": {
		Data: "Issued a create database symmetric key command (action_id CR class_type SK)",
		Orig: "SQL Server",
	},
	"24104": {
		Data: "Issued a change database symmetric key command (action_id AL class_type SK)",
		Orig: "SQL Server",
	},
	"24105": {
		Data: "Issued a delete database symmetric key command (action_id DR class_type SK)",
		Orig: "SQL Server",
	},
	"24106": {
		Data: "Issued a back up database symmetric key command (action_id BA class_type SK)",
		Orig: "SQL Server",
	},
	"24107": {
		Data: "Issued an open database symmetric key command (action_id OP class_type SK)",
		Orig: "SQL Server",
	},
	"24108": {
		Data: "Issued a create database object command (action_id CR)",
		Orig: "SQL Server",
	},
	"24109": {
		Data: "Issued a change database object command (action_id AL)",
		Orig: "SQL Server",
	},
	"24110": {
		Data: "Issued a delete database object command (action_id DR)",
		Orig: "SQL Server",
	},
	"24111": {
		Data: "Issued an access database object command (action_id AS)",
		Orig: "SQL Server",
	},
	"24112": {
		Data: "Issued a create assembly command (action_id CR class_type AS)",
		Orig: "SQL Server",
	},
	"24113": {
		Data: "Issued a change assembly command (action_id AL class_type AS)",
		Orig: "SQL Server",
	},
	"24114": {
		Data: "Issued a delete assembly command (action_id DR class_type AS)",
		Orig: "SQL Server",
	},
	"24115": {
		Data: "Issued a create schema command (action_id CR class_type SC)",
		Orig: "SQL Server",
	},
	"24116": {
		Data: "Issued a change schema command (action_id AL class_type SC)",
		Orig: "SQL Server",
	},
	"24117": {
		Data: "Issued a delete schema command (action_id DR class_type SC)",
		Orig: "SQL Server",
	},
	"24118": {
		Data: "Issued a create database encryption key command (action_id CR class_type DK)",
		Orig: "SQL Server",
	},
	"24119": {
		Data: "Issued a change database encryption key command (action_id AL class_type DK)",
		Orig: "SQL Server",
	},
	"24120": {
		Data: "Issued a delete database encryption key command (action_id DR class_type DK)",
		Orig: "SQL Server",
	},
	"24121": {
		Data: "Issued a create database user command (action_id CR; class_type US)",
		Orig: "SQL Server",
	},
	"24122": {
		Data: "Issued a change database user command (action_id AL; class_type US)",
		Orig: "SQL Server",
	},
	"24123": {
		Data: "Issued a delete database user command (action_id DR; class_type US)",
		Orig: "SQL Server",
	},
	"24124": {
		Data: "Issued a create database role command (action_id CR class_type RL)",
		Orig: "SQL Server",
	},
	"24125": {
		Data: "Issued a change database role command (action_id AL class_type RL)",
		Orig: "SQL Server",
	},
	"24126": {
		Data: "Issued a delete database role command (action_id DR class_type RL)",
		Orig: "SQL Server",
	},
	"24127": {
		Data: "Issued a create application role command (action_id CR class_type AR)",
		Orig: "SQL Server",
	},
	"24128": {
		Data: "Issued a change application role command (action_id AL class_type AR)",
		Orig: "SQL Server",
	},
	"24129": {
		Data: "Issued a delete application role command (action_id DR class_type AR)",
		Orig: "SQL Server",
	},
	"24130": {
		Data: "Issued a change database user login command (action_id USAF)",
		Orig: "SQL Server",
	},
	"24131": {
		Data: "Issued an auto-change database user login command (action_id USLG)",
		Orig: "SQL Server",
	},
	"24132": {
		Data: "Issued a create schema object command (action_id CR class_type D)",
		Orig: "SQL Server",
	},
	"24133": {
		Data: "Issued a change schema object command (action_id AL class_type D)",
		Orig: "SQL Server",
	},
	"24134": {
		Data: "Issued a delete schema object command (action_id DR class_type D)",
		Orig: "SQL Server",
	},
	"24135": {
		Data: "Issued a transfer schema object command (action_id TRO class_type D)",
		Orig: "SQL Server",
	},
	"24136": {
		Data: "Issued a create schema type command (action_id CR class_type TY)",
		Orig: "SQL Server",
	},
	"24137": {
		Data: "Issued a change schema type command (action_id AL class_type TY)",
		Orig: "SQL Server",
	},
	"24138": {
		Data: "Issued a delete schema type command (action_id DR class_type TY)",
		Orig: "SQL Server",
	},
	"24139": {
		Data: "Issued a transfer schema type command (action_id TRO class_type TY)",
		Orig: "SQL Server",
	},
	"24140": {
		Data: "Issued a create XML schema collection command (action_id CR class_type SX)",
		Orig: "SQL Server",
	},
	"24141": {
		Data: "Issued a change XML schema collection command (action_id AL class_type SX)",
		Orig: "SQL Server",
	},
	"24142": {
		Data: "Issued a delete XML schema collection command (action_id DR class_type SX)",
		Orig: "SQL Server",
	},
	"24143": {
		Data: "Issued a transfer XML schema collection command (action_id TRO class_type SX)",
		Orig: "SQL Server",
	},
	"24144": {
		Data: "Issued an impersonate within server scope command (action_id IMP; class_type LX)",
		Orig: "SQL Server",
	},
	"24145": {
		Data: "Issued an impersonate within database scope command (action_id IMP; class_type US)",
		Orig: "SQL Server",
	},
	"24146": {
		Data: "Issued a change server object owner command (action_id TO class_type SG)",
		Orig: "SQL Server",
	},
	"24147": {
		Data: "Issued a change database owner command (action_id TO class_type DB)",
		Orig: "SQL Server",
	},
	"24148": {
		Data: "Issued a change schema owner command (action_id TO class_type SC)",
		Orig: "SQL Server",
	},
	"24150": {
		Data: "Issued a change role owner command (action_id TO class_type RL)",
		Orig: "SQL Server",
	},
	"24151": {
		Data: "Issued a change database object owner command (action_id TO)",
		Orig: "SQL Server",
	},
	"24152": {
		Data: "Issued a change symmetric key owner command (action_id TO class_type SK)",
		Orig: "SQL Server",
	},
	"24153": {
		Data: "Issued a change certificate owner command (action_id TO class_type CR)",
		Orig: "SQL Server",
	},
	"24154": {
		Data: "Issued a change asymmetric key owner command (action_id TO class_type AK)",
		Orig: "SQL Server",
	},
	"24155": {
		Data: "Issued a change schema object owner command (action_id TO class_type OB)",
		Orig: "SQL Server",
	},
	"24156": {
		Data: "Issued a change schema type owner command (action_id TO class_type TY)",
		Orig: "SQL Server",
	},
	"24157": {
		Data: "Issued a change XML schema collection owner command (action_id TO class_type SX)",
		Orig: "SQL Server",
	},
	"24158": {
		Data: "Grant server permissions succeeded (action_id G class_type SR)",
		Orig: "SQL Server",
	},
	"24159": {
		Data: "Grant server permissions failed (action_id G class_type SR)",
		Orig: "SQL Server",
	},
	"24160": {
		Data: "Grant server permissions with grant succeeded (action_id GWG class_type SR)",
		Orig: "SQL Server",
	},
	"24161": {
		Data: "Grant server permissions with grant failed (action_id GWG class_type SR)",
		Orig: "SQL Server",
	},
	"24162": {
		Data: "Deny server permissions succeeded (action_id D class_type SR)",
		Orig: "SQL Server",
	},
	"24163": {
		Data: "Deny server permissions failed (action_id D class_type SR)",
		Orig: "SQL Server",
	},
	"24164": {
		Data: "Deny server permissions with cascade succeeded (action_id DWC class_type SR)",
		Orig: "SQL Server",
	},
	"24165": {
		Data: "Deny server permissions with cascade failed (action_id DWC class_type SR)",
		Orig: "SQL Server",
	},
	"24166": {
		Data: "Revoke server permissions succeeded (action_id R class_type SR)",
		Orig: "SQL Server",
	},
	"24167": {
		Data: "Revoke server permissions failed (action_id R class_type SR)",
		Orig: "SQL Server",
	},
	"24168": {
		Data: "Revoke server permissions with grant succeeded (action_id RWG class_type SR)",
		Orig: "SQL Server",
	},
	"24169": {
		Data: "Revoke server permissions with grant failed (action_id RWG class_type SR)",
		Orig: "SQL Server",
	},
	"24170": {
		Data: "Revoke server permissions with cascade succeeded (action_id RWC class_type SR)",
		Orig: "SQL Server",
	},
	"24171": {
		Data: "Revoke server permissions with cascade failed (action_id RWC class_type SR)",
		Orig: "SQL Server",
	},
	"24172": {
		Data: "Issued grant server object permissions command (action_id G; class_type LX)",
		Orig: "SQL Server",
	},
	"24173": {
		Data: "Issued grant server object permissions with grant command (action_id GWG; class_type LX)",
		Orig: "SQL Server",
	},
	"24174": {
		Data: "Issued deny server object permissions command (action_id D; class_type LX)",
		Orig: "SQL Server",
	},
	"24175": {
		Data: "Issued deny server object permissions with cascade command (action_id DWC; class_type LX)",
		Orig: "SQL Server",
	},
	"24176": {
		Data: "Issued revoke server object permissions command (action_id R; class_type LX)",
		Orig: "SQL Server",
	},
	"24177": {
		Data: "Issued revoke server object permissions with grant command (action_id; RWG class_type LX)",
		Orig: "SQL Server",
	},
	"24178": {
		Data: "Issued revoke server object permissions with cascade command (action_id RWC; class_type LX)",
		Orig: "SQL Server",
	},
	"24179": {
		Data: "Grant database permissions succeeded (action_id G class_type DB)",
		Orig: "SQL Server",
	},
	"24180": {
		Data: "Grant database permissions failed (action_id G class_type DB)",
		Orig: "SQL Server",
	},
	"24181": {
		Data: "Grant database permissions with grant succeeded (action_id GWG class_type DB)",
		Orig: "SQL Server",
	},
	"24182": {
		Data: "Grant database permissions with grant failed (action_id GWG class_type DB)",
		Orig: "SQL Server",
	},
	"24183": {
		Data: "Deny database permissions succeeded (action_id D class_type DB)",
		Orig: "SQL Server",
	},
	"24184": {
		Data: "Deny database permissions failed (action_id D class_type DB)",
		Orig: "SQL Server",
	},
	"24185": {
		Data: "Deny database permissions with cascade succeeded (action_id DWC class_type DB)",
		Orig: "SQL Server",
	},
	"24186": {
		Data: "Deny database permissions with cascade failed (action_id DWC class_type DB)",
		Orig: "SQL Server",
	},
	"24187": {
		Data: "Revoke database permissions succeeded (action_id R class_type DB)",
		Orig: "SQL Server",
	},
	"24188": {
		Data: "Revoke database permissions failed (action_id R class_type DB)",
		Orig: "SQL Server",
	},
	"24189": {
		Data: "Revoke database permissions with grant succeeded (action_id RWG class_type DB)",
		Orig: "SQL Server",
	},
	"24190": {
		Data: "Revoke database permissions with grant failed (action_id RWG class_type DB)",
		Orig: "SQL Server",
	},
	"24191": {
		Data: "Revoke database permissions with cascade succeeded (action_id RWC class_type DB)",
		Orig: "SQL Server",
	},
	"24192": {
		Data: "Revoke database permissions with cascade failed (action_id RWC class_type DB)",
		Orig: "SQL Server",
	},
	"24193": {
		Data: "Issued grant database object permissions command (action_id G class_type US)",
		Orig: "SQL Server",
	},
	"24194": {
		Data: "Issued grant database object permissions with grant command (action_id GWG; class_type US)",
		Orig: "SQL Server",
	},
	"24195": {
		Data: "Issued deny database object permissions command (action_id D; class_type US)",
		Orig: "SQL Server",
	},
	"24196": {
		Data: "Issued deny database object permissions with cascade command (action_id DWC; class_type US)",
		Orig: "SQL Server",
	},
	"24197": {
		Data: "Issued revoke database object permissions command (action_id R; class_type US)",
		Orig: "SQL Server",
	},
	"24198": {
		Data: "Issued revoke database object permissions with grant command (action_id RWG; class_type US)",
		Orig: "SQL Server",
	},
	"24199": {
		Data: "Issued revoke database object permissions with cascade command (action_id RWC; class_type US)",
		Orig: "SQL Server",
	},
	"24200": {
		Data: "Issued grant schema permissions command (action_id G class_type SC)",
		Orig: "SQL Server",
	},
	"24201": {
		Data: "Issued grant schema permissions with grant command (action_id GWG class_type SC)",
		Orig: "SQL Server",
	},
	"24202": {
		Data: "Issued deny schema permissions command (action_id D class_type SC)",
		Orig: "SQL Server",
	},
	"24203": {
		Data: "Issued deny schema permissions with cascade command (action_id DWC class_type SC)",
		Orig: "SQL Server",
	},
	"24204": {
		Data: "Issued revoke schema permissions command (action_id R class_type SC)",
		Orig: "SQL Server",
	},
	"24205": {
		Data: "Issued revoke schema permissions with grant command (action_id RWG class_type SC)",
		Orig: "SQL Server",
	},
	"24206": {
		Data: "Issued revoke schema permissions with cascade command (action_id RWC class_type SC)",
		Orig: "SQL Server",
	},
	"24207": {
		Data: "Issued grant assembly permissions command (action_id G class_type AS)",
		Orig: "SQL Server",
	},
	"24208": {
		Data: "Issued grant assembly permissions with grant command (action_id GWG class_type AS)",
		Orig: "SQL Server",
	},
	"24209": {
		Data: "Issued deny assembly permissions command (action_id D class_type AS)",
		Orig: "SQL Server",
	},
	"24210": {
		Data: "Issued deny assembly permissions with cascade command (action_id DWC class_type AS)",
		Orig: "SQL Server",
	},
	"24211": {
		Data: "Issued revoke assembly permissions command (action_id R class_type AS)",
		Orig: "SQL Server",
	},
	"24212": {
		Data: "Issued revoke assembly permissions with grant command (action_id RWG class_type AS)",
		Orig: "SQL Server",
	},
	"24213": {
		Data: "Issued revoke assembly permissions with cascade command (action_id RWC class_type AS)",
		Orig: "SQL Server",
	},
	"24214": {
		Data: "Issued grant database role permissions command (action_id G class_type RL)",
		Orig: "SQL Server",
	},
	"24215": {
		Data: "Issued grant database role permissions with grant command (action_id GWG class_type RL)",
		Orig: "SQL Server",
	},
	"24216": {
		Data: "Issued deny database role permissions command (action_id D class_type RL)",
		Orig: "SQL Server",
	},
	"24217": {
		Data: "Issued deny database role permissions with cascade command (action_id DWC class_type RL)",
		Orig: "SQL Server",
	},
	"24218": {
		Data: "Issued revoke database role permissions command (action_id R class_type RL)",
		Orig: "SQL Server",
	},
	"24219": {
		Data: "Issued revoke database role permissions with grant command (action_id RWG class_type RL)",
		Orig: "SQL Server",
	},
	"24220": {
		Data: "Issued revoke database role permissions with cascade command (action_id RWC class_type RL)",
		Orig: "SQL Server",
	},
	"24221": {
		Data: "Issued grant application role permissions command (action_id G class_type AR)",
		Orig: "SQL Server",
	},
	"24222": {
		Data: "Issued grant application role permissions with grant command (action_id GWG class_type AR)",
		Orig: "SQL Server",
	},
	"24223": {
		Data: "Issued deny application role permissions command (action_id D class_type AR)",
		Orig: "SQL Server",
	},
	"24224": {
		Data: "Issued deny application role permissions with cascade command (action_id DWC class_type AR)",
		Orig: "SQL Server",
	},
	"24225": {
		Data: "Issued revoke application role permissions command (action_id R class_type AR)",
		Orig: "SQL Server",
	},
	"24226": {
		Data: "Issued revoke application role permissions with grant command (action_id RWG class_type AR)",
		Orig: "SQL Server",
	},
	"24227": {
		Data: "Issued revoke application role permissions with cascade command (action_id RWC class_type AR)",
		Orig: "SQL Server",
	},
	"24228": {
		Data: "Issued grant symmetric key permissions command (action_id G class_type SK)",
		Orig: "SQL Server",
	},
	"24229": {
		Data: "Issued grant symmetric key permissions with grant command (action_id GWG class_type SK)",
		Orig: "SQL Server",
	},
	"24230": {
		Data: "Issued deny symmetric key permissions command (action_id D class_type SK)",
		Orig: "SQL Server",
	},
	"24231": {
		Data: "Issued deny symmetric key permissions with cascade command (action_id DWC class_type SK)",
		Orig: "SQL Server",
	},
	"24232": {
		Data: "Issued revoke symmetric key permissions command (action_id R class_type SK)",
		Orig: "SQL Server",
	},
	"24233": {
		Data: "Issued revoke symmetric key permissions with grant command (action_id RWG class_type SK)",
		Orig: "SQL Server",
	},
	"24234": {
		Data: "Issued revoke symmetric key permissions with cascade command (action_id RWC class_type SK)",
		Orig: "SQL Server",
	},
	"24235": {
		Data: "Issued grant certificate permissions command (action_id G class_type CR)",
		Orig: "SQL Server",
	},
	"24236": {
		Data: "Issued grant certificate permissions with grant command (action_id GWG class_type CR)",
		Orig: "SQL Server",
	},
	"24237": {
		Data: "Issued deny certificate permissions command (action_id D class_type CR)",
		Orig: "SQL Server",
	},
	"24238": {
		Data: "Issued deny certificate permissions with cascade command (action_id DWC class_type CR)",
		Orig: "SQL Server",
	},
	"24239": {
		Data: "Issued revoke certificate permissions command (action_id R class_type CR)",
		Orig: "SQL Server",
	},
	"24240": {
		Data: "Issued revoke certificate permissions with grant command (action_id RWG class_type CR)",
		Orig: "SQL Server",
	},
	"24241": {
		Data: "Issued revoke certificate permissions with cascade command (action_id RWC class_type CR)",
		Orig: "SQL Server",
	},
	"24242": {
		Data: "Issued grant asymmetric key permissions command (action_id G class_type AK)",
		Orig: "SQL Server",
	},
	"24243": {
		Data: "Issued grant asymmetric key permissions with grant command (action_id GWG class_type AK)",
		Orig: "SQL Server",
	},
	"24244": {
		Data: "Issued deny asymmetric key permissions command (action_id D class_type AK)",
		Orig: "SQL Server",
	},
	"24245": {
		Data: "Issued deny asymmetric key permissions with cascade command (action_id DWC class_type AK)",
		Orig: "SQL Server",
	},
	"24246": {
		Data: "Issued revoke asymmetric key permissions command (action_id R class_type AK)",
		Orig: "SQL Server",
	},
	"24247": {
		Data: "Issued revoke asymmetric key permissions with grant command (action_id RWG class_type AK)",
		Orig: "SQL Server",
	},
	"24248": {
		Data: "Issued revoke asymmetric key permissions with cascade command (action_id RWC class_type AK)",
		Orig: "SQL Server",
	},
	"24249": {
		Data: "Issued grant schema object permissions command (action_id G class_type OB)",
		Orig: "SQL Server",
	},
	"24250": {
		Data: "Issued grant schema object permissions with grant command (action_id GWG class_type OB)",
		Orig: "SQL Server",
	},
	"24251": {
		Data: "Issued deny schema object permissions command (action_id D class_type OB)",
		Orig: "SQL Server",
	},
	"24252": {
		Data: "Issued deny schema object permissions with cascade command (action_id DWC class_type OB)",
		Orig: "SQL Server",
	},
	"24253": {
		Data: "Issued revoke schema object permissions command (action_id R class_type OB)",
		Orig: "SQL Server",
	},
	"24254": {
		Data: "Issued revoke schema object permissions with grant command (action_id RWG class_type OB)",
		Orig: "SQL Server",
	},
	"24255": {
		Data: "Issued revoke schema object permissions with cascade command (action_id RWC class_type OB)",
		Orig: "SQL Server",
	},
	"24256": {
		Data: "Issued grant schema type permissions command (action_id G class_type TY)",
		Orig: "SQL Server",
	},
	"24257": {
		Data: "Issued grant schema type permissions with grant command (action_id GWG class_type TY)",
		Orig: "SQL Server",
	},
	"24258": {
		Data: "Issued deny schema type permissions command (action_id D class_type TY)",
		Orig: "SQL Server",
	},
	"24259": {
		Data: "Issued deny schema type permissions with cascade command (action_id DWC class_type TY)",
		Orig: "SQL Server",
	},
	"24260": {
		Data: "Issued revoke schema type permissions command (action_id R class_type TY)",
		Orig: "SQL Server",
	},
	"24261": {
		Data: "Issued revoke schema type permissions with grant command (action_id RWG class_type TY)",
		Orig: "SQL Server",
	},
	"24262": {
		Data: "Issued revoke schema type permissions with cascade command (action_id RWC class_type TY)",
		Orig: "SQL Server",
	},
	"24263": {
		Data: "Issued grant XML schema collection permissions command (action_id G class_type SX)",
		Orig: "SQL Server",
	},
	"24264": {
		Data: "Issued grant XML schema collection permissions with grant command (action_id GWG class_type SX)",
		Orig: "SQL Server",
	},
	"24265": {
		Data: "Issued deny XML schema collection permissions command (action_id D class_type SX)",
		Orig: "SQL Server",
	},
	"24266": {
		Data: "Issued deny XML schema collection permissions with cascade command (action_id DWC class_type SX)",
		Orig: "SQL Server",
	},
	"24267": {
		Data: "Issued revoke XML schema collection permissions command (action_id R class_type SX)",
		Orig: "SQL Server",
	},
	"24268": {
		Data: "Issued revoke XML schema collection permissions with grant command (action_id RWG class_type SX)",
		Orig: "SQL Server",
	},
	"24269": {
		Data: "Issued revoke XML schema collection permissions with cascade command (action_id RWC class_type SX)",
		Orig: "SQL Server",
	},
	"24270": {
		Data: "Issued reference database object permissions command (action_id RF)",
		Orig: "SQL Server",
	},
	"24271": {
		Data: "Issued send service request command (action_id SN)",
		Orig: "SQL Server",
	},
	"24272": {
		Data: "Issued check permissions with schema command (action_id VWCT)",
		Orig: "SQL Server",
	},
	"24273": {
		Data: "Issued use service broker transport security command (action_id LGB)",
		Orig: "SQL Server",
	},
	"24274": {
		Data: "Issued use database mirroring transport security command (action_id LGM)",
		Orig: "SQL Server",
	},
	"24275": {
		Data: "Issued alter trace command (action_id ALTR)",
		Orig: "SQL Server",
	},
	"24276": {
		Data: "Issued start trace command (action_id TASA)",
		Orig: "SQL Server",
	},
	"24277": {
		Data: "Issued stop trace command (action_id TASP)",
		Orig: "SQL Server",
	},
	"24278": {
		Data: "Issued enable trace C2 audit mode command (action_id C2ON)",
		Orig: "SQL Server",
	},
	"24279": {
		Data: "Issued disable trace C2 audit mode command (action_id C2OF)",
		Orig: "SQL Server",
	},
	"24280": {
		Data: "Issued server full-text command (action_id FT)",
		Orig: "SQL Server",
	},
	"24281": {
		Data: "Issued select command (action_id SL)",
		Orig: "SQL Server",
	},
	"24282": {
		Data: "Issued update command (action_id UP)",
		Orig: "SQL Server",
	},
	"24283": {
		Data: "Issued insert command (action_id IN)",
		Orig: "SQL Server",
	},
	"24284": {
		Data: "Issued delete command (action_id DL)",
		Orig: "SQL Server",
	},
	"24285": {
		Data: "Issued execute command (action_id EX)",
		Orig: "SQL Server",
	},
	"24286": {
		Data: "Issued receive command (action_id RC)",
		Orig: "SQL Server",
	},
	"24287": {
		Data: "Issued check references command (action_id RF)",
		Orig: "SQL Server",
	},
	"24288": {
		Data: "Issued a create user-defined server role command (action_id CR class_type SG)",
		Orig: "SQL Server",
	},
	"24289": {
		Data: "Issued a change user-defined server role command (action_id AL class_type SG)",
		Orig: "SQL Server",
	},
	"24290": {
		Data: "Issued a delete user-defined server role command (action_id DR class_type SG)",
		Orig: "SQL Server",
	},
	"24291": {
		Data: "Issued grant user-defined server role permissions command (action_id G class_type SG)",
		Orig: "SQL Server",
	},
	"24292": {
		Data: "Issued grant user-defined server role permissions with grant command (action_id GWG class_type SG)",
		Orig: "SQL Server",
	},
	"24293": {
		Data: "Issued deny user-defined server role permissions command (action_id D class_type SG)",
		Orig: "SQL Server",
	},
	"24294": {
		Data: "Issued deny user-defined server role permissions with cascade command (action_id DWC class_type SG)",
		Orig: "SQL Server",
	},
	"24295": {
		Data: "Issued revoke user-defined server role permissions command (action_id R class_type SG)",
		Orig: "SQL Server",
	},
	"24296": {
		Data: "Issued revoke user-defined server role permissions with grant command (action_id RWG class_type SG)",
		Orig: "SQL Server",
	},
	"24297": {
		Data: "Issued revoke user-defined server role permissions with cascade command (action_id RWC class_type SG)",
		Orig: "SQL Server",
	},
	"24298": {
		Data: "Database login succeeded (action_id DBAS)",
		Orig: "SQL Server",
	},
	"24299": {
		Data: "Database login failed (action_id DBAF)",
		Orig: "SQL Server",
	},
	"24300": {
		Data: "Database logout successful (action_id DAGL)",
		Orig: "SQL Server",
	},
	"24301": {
		Data: "Change password succeeded (action_id PWC; class_type US)",
		Orig: "SQL Server",
	},
	"24302": {
		Data: "Change password failed (action_id PWC; class_type US)",
		Orig: "SQL Server",
	},
	"24303": {
		Data: "Change own password succeeded (action_id PWCS; class_type US)",
		Orig: "SQL Server",
	},
	"24304": {
		Data: "Change own password failed (action_id PWCS; class_type US)",
		Orig: "SQL Server",
	},
	"24305": {
		Data: "Reset own password succeeded (action_id PWRS; class_type US)",
		Orig: "SQL Server",
	},
	"24306": {
		Data: "Reset own password failed (action_id PWRS; class_type US)",
		Orig: "SQL Server",
	},
	"24307": {
		Data: "Reset password succeeded (action_id PWR; class_type US)",
		Orig: "SQL Server",
	},
	"24308": {
		Data: "Reset password failed (action_id PWR; class_type US)",
		Orig: "SQL Server",
	},
	"24309": {
		Data: "Copy password (action_id USTC)",
		Orig: "SQL Server",
	},
	"24310": {
		Data: "User-defined SQL audit event (action_id UDAU)",
		Orig: "SQL Server",
	},
	"24311": {
		Data: "Issued a change database audit command (action_id AL class_type DU)",
		Orig: "SQL Server",
	},
	"24312": {
		Data: "Issued a create database audit command (action_id CR class_type DU)",
		Orig: "SQL Server",
	},
	"24313": {
		Data: "Issued a delete database audit command (action_id DR class_type DU)",
		Orig: "SQL Server",
	},
	"24314": {
		Data: "Issued a begin transaction command (action_id TXBG)",
		Orig: "SQL Server",
	},
	"24315": {
		Data: "Issued a commit transaction command (action_id TXCM)",
		Orig: "SQL Server",
	},
	"24316": {
		Data: "Issued a rollback transaction command (action_id TXRB)",
		Orig: "SQL Server",
	},
	"24317": {
		Data: "Issued a create column master key command (action_id CR; class_type CM)",
		Orig: "SQL Server",
	},
	"24318": {
		Data: "Issued a delete column master key command (action_id DR; class_type CM)",
		Orig: "SQL Server",
	},
	"24319": {
		Data: "A column master key was viewed (action_id VW; class_type CM)",
		Orig: "SQL Server",
	},
	"24320": {
		Data: "Issued a create column encryption key command (action_id CR; class_type CK)",
		Orig: "SQL Server",
	},
	"24321": {
		Data: "Issued a change column encryption key command (action_id AL; class_type CK)",
		Orig: "SQL Server",
	},
	"24322": {
		Data: "Issued a delete column encryption key command (action_id DR; class_type CK)",
		Orig: "SQL Server",
	},
	"24323": {
		Data: "A column encryption key was viewed (action_id VW; class_type CK)",
		Orig: "SQL Server",
	},
	"24324": {
		Data: "Issued a create database credential command (action_id CR; class_type DC)",
		Orig: "SQL Server",
	},
	"24325": {
		Data: "Issued a change database credential command (action_id AL; class_type DC)",
		Orig: "SQL Server",
	},
	"24326": {
		Data: "Issued a delete database credential command (action_id DR; class_type DC)",
		Orig: "SQL Server",
	},
	"24327": {
		Data: "Issued a change database scoped configuration command (action_id AL; class_type DS)",
		Orig: "SQL Server",
	},
	"24328": {
		Data: "Issued a create external data Orig command (action_id CR; class_type ED)",
		Orig: "SQL Server",
	},
	"24329": {
		Data: "Issued a change external data Orig command (action_id AL; class_type ED)",
		Orig: "SQL Server",
	},
	"24330": {
		Data: "Issued a delete external data Orig command (action_id DR; class_type ED)",
		Orig: "SQL Server",
	},
	"24331": {
		Data: "Issued a create external file format command (action_id CR; class_type EF)",
		Orig: "SQL Server",
	},
	"24332": {
		Data: "Issued a delete external file format command (action_id DR; class_type EF)",
		Orig: "SQL Server",
	},
	"24333": {
		Data: "Issued a create external reOrig pool command (action_id CR; class_type ER)",
		Orig: "SQL Server",
	},
	"24334": {
		Data: "Issued a change external reOrig pool command (action_id AL; class_type ER)",
		Orig: "SQL Server",
	},
	"24335": {
		Data: "Issued a delete external reOrig pool command (action_id DR; class_type ER)",
		Orig: "SQL Server",
	},
	"24337": {
		Data: "Global transaction login (action_id LGG)",
		Orig: "SQL Server",
	},
	"24338": {
		Data: "Grant permissions on a database scoped credential succeeded (action_id G; class_type DC)",
		Orig: "SQL Server",
	},
	"24339": {
		Data: "Grant permissions on a database scoped credential failed (action_id G; class_type DC)",
		Orig: "SQL Server",
	},
	"24340": {
		Data: "Grant permissions on a database scoped credential with grant succeeded (action_id GWG; class_type DC)",
		Orig: "SQL Server",
	},
	"24341": {
		Data: "Grant permissions on a database scoped credential with grant failed (action_id GWG; class_type DC)",
		Orig: "SQL Server",
	},
	"24342": {
		Data: "Deny permissions on a database scoped credential succeeded (action_id D; class_type DC)",
		Orig: "SQL Server",
	},
	"24343": {
		Data: "Deny permissions on a database scoped credential failed (action_id D; class_type DC)",
		Orig: "SQL Server",
	},
	"24344": {
		Data: "Deny permissions on a database scoped credential with cascade succeeded (action_id DWC; class_type DC)",
		Orig: "SQL Server",
	},
	"24345": {
		Data: "Deny permissions on a database scoped credential with cascade failed (action_id DWC; class_type DC)",
		Orig: "SQL Server",
	},
	"24346": {
		Data: "Revoke permissions on a database scoped credential succeeded (action_id R; class_type DC)",
		Orig: "SQL Server",
	},
	"24347": {
		Data: "Revoke permissions on a database scoped credential failed (action_id R; class_type DC)",
		Orig: "SQL Server",
	},
	"24348": {
		Data: "Revoke permissions with cascade on a database scoped credential succeeded (action_id RWC; class_type DC)",
		Orig: "SQL Server",
	},
	"24349": {
		Data: "Issued a change assembly owner command (action_id TO class_type AS)",
		Orig: "SQL Server",
	},
	"24350": {
		Data: "Revoke permissions with cascade on a database scoped credential failed (action_id RWC; class_type DC)",
		Orig: "SQL Server",
	},
	"24351": {
		Data: "Revoke permissions with grant on a database scoped credential succeeded (action_id RWG; class_type DC)",
		Orig: "SQL Server",
	},
	"24352": {
		Data: "Revoke permissions with grant on a database scoped credential failed (action_id RWG; class_type DC)",
		Orig: "SQL Server",
	},
	"24353": {
		Data: "Issued a change database scoped credential owner command (action_id TO; class_type DC)",
		Orig: "SQL Server",
	},
	"24354": {
		Data: "Issued a create external library command (action_id CR; class_type EL)",
		Orig: "SQL Server",
	},
	"24355": {
		Data: "Issued a change external library command (action_id AL; class_type EL)",
		Orig: "SQL Server",
	},
	"24356": {
		Data: "Issued a drop external library command (action_id DR; class_type EL)",
		Orig: "SQL Server",
	},
	"24357": {
		Data: "Grant permissions on an external library succeeded (action_id G; class_type EL)",
		Orig: "SQL Server",
	},
	"24358": {
		Data: "Grant permissions on an external library failed (action_id G; class_type EL)",
		Orig: "SQL Server",
	},
	"24359": {
		Data: "Grant permissions on an external library with grant succeeded (action_id GWG; class_type EL)",
		Orig: "SQL Server",
	},
	"24360": {
		Data: "Grant permissions on an external library with grant failed (action_id GWG; class_type EL)",
		Orig: "SQL Server",
	},
	"24361": {
		Data: "Deny permissions on an external library succeeded (action_id D; class_type EL)",
		Orig: "SQL Server",
	},
	"24362": {
		Data: "Deny permissions on an external library failed (action_id D; class_type EL)",
		Orig: "SQL Server",
	},
	"24363": {
		Data: "Deny permissions on an external library with cascade succeeded (action_id DWC; class_type EL)",
		Orig: "SQL Server",
	},
	"24364": {
		Data: "Deny permissions on an external library with cascade failed (action_id DWC; class_type EL)",
		Orig: "SQL Server",
	},
	"24365": {
		Data: "Revoke permissions on an external library succeeded (action_id R; class_type EL)",
		Orig: "SQL Server",
	},
	"24366": {
		Data: "Revoke permissions on an external library failed (action_id R; class_type EL)",
		Orig: "SQL Server",
	},
	"24367": {
		Data: "Revoke permissions with cascade on an external library succeeded (action_id RWC; class_type EL)",
		Orig: "SQL Server",
	},
	"24368": {
		Data: "Revoke permissions with cascade on an external library failed (action_id RWC; class_type EL)",
		Orig: "SQL Server",
	},
	"24369": {
		Data: "Revoke permissions with grant on an external library succeeded (action_id RWG; class_type EL)",
		Orig: "SQL Server",
	},
	"24370": {
		Data: "Revoke permissions with grant on an external library failed (action_id RWG; class_type EL)",
		Orig: "SQL Server",
	},
	"24371": {
		Data: "Issued a create database scoped reOrig governor command (action_id CR; class_type DR)",
		Orig: "SQL Server",
	},
	"24372": {
		Data: "Issued a change database scoped reOrig governor command (action_id AL; class_type DR)",
		Orig: "SQL Server",
	},
	"24373": {
		Data: "Issued a drop database scoped reOrig governor command (action_id DR; class_type DR)",
		Orig: "SQL Server",
	},
	"24374": {
		Data: "Issued a database bulk administration command (action_id DABO; class_type DB)",
		Orig: "SQL Server",
	},
	"24375": {
		Data: "Command to change permission failed (action_id D, DWC, G, GWG, R, RWC, RWG; class_type DC, EL)",
		Orig: "SQL Server",
	},
	"25000": {
		Data: "Undocumented Exchange mailbox operation",
		Orig: "Exchange",
	},
	"25001": {
		Data: "Operation Copy - Copy item to another Exchange mailbox folder",
		Orig: "Exchange",
	},
	"25002": {
		Data: "Operation Create - Create item in Exchange mailbox",
		Orig: "Exchange",
	},
	"25003": {
		Data: "Operation FolderBind - Access Exchange mailbox folder",
		Orig: "Exchange",
	},
	"25004": {
		Data: "Operation HardDelete - Delete Exchange mailbox item permanently from Recoverable Items folder",
		Orig: "Exchange",
	},
	"25005": {
		Data: "Operation MessageBind - Access Exchange mailbox item",
		Orig: "Exchange",
	},
	"25006": {
		Data: "Operation Move - Move item to another Exchange mailbox folder",
		Orig: "Exchange",
	},
	"25007": {
		Data: "Operation MoveToDeletedItems - Move Exchange mailbox item to Deleted Items folder",
		Orig: "Exchange",
	},
	"25008": {
		Data: "Operation SendAs - Send message using Send As Exchange mailbox permissions",
		Orig: "Exchange",
	},
	"25009": {
		Data: "Operation SendOnBehalf - Send message using Send on Behalf Exchange mailbox permissions",
		Orig: "Exchange",
	},
	"25010": {
		Data: "Operation SoftDelete - Delete Exchange mailbox item from Deleted Items folder",
		Orig: "Exchange",
	},
	"25011": {
		Data: "Operation Update - Update Exchange mailbox item's properties",
		Orig: "Exchange",
	},
	"25012": {
		Data: "Information Event - Mailbox audit policy applied",
		Orig: "Exchange",
	},
	"25100": {
		Data: "Undocumented Exchange admin operation",
		Orig: "Exchange",
	},
	"25101": {
		Data: "Add-ADPermission Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25102": {
		Data: "Add-AvailabilityAddressSpace Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25103": {
		Data: "Add-ContentFilterPhrase Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25104": {
		Data: "Add-DatabaseAvailabilityGroupServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25105": {
		Data: "Add-DistributionGroupMember Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25106": {
		Data: "Add-FederatedDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25107": {
		Data: "Add-IPAllowListEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25108": {
		Data: "Add-IPAllowListProvider Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25109": {
		Data: "Add-IPBlockListEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25110": {
		Data: "Add-IPBlockListProvider Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25111": {
		Data: "Add-MailboxDatabaseCopy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25112": {
		Data: "Add-MailboxFolderPermission Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25113": {
		Data: "Add-MailboxPermission Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25114": {
		Data: "Add-ManagementRoleEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25115": {
		Data: "Add-PublicFolderAdministrativePermission Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25116": {
		Data: "Add-PublicFolderClientPermission Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25117": {
		Data: "Add-RoleGroupMember Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25118": {
		Data: "Clean-MailboxDatabase Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25119": {
		Data: "Clear-ActiveSyncDevice Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25120": {
		Data: "Clear-TextMessagingAccount Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25121": {
		Data: "Compare-TextMessagingVerificationCode Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25122": {
		Data: "Connect-Mailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25123": {
		Data: "Disable-AddressListPaging Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25124": {
		Data: "Disable-CmdletExtensionAgent Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25125": {
		Data: "Disable-DistributionGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25126": {
		Data: "Disable-InboxRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25127": {
		Data: "Disable-JournalRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25128": {
		Data: "Disable-Mailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25129": {
		Data: "Disable-MailContact Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25130": {
		Data: "Disable-MailPublicFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25131": {
		Data: "Disable-MailUser Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25132": {
		Data: "Disable-OutlookAnywhere Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25133": {
		Data: "Disable-OutlookProtectionRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25134": {
		Data: "Disable-RemoteMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25135": {
		Data: "Disable-ServiceEmailChannel Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25136": {
		Data: "Disable-TransportAgent Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25137": {
		Data: "Disable-TransportRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25138": {
		Data: "Disable-UMAutoAttendant Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25139": {
		Data: "Disable-UMIPGateway Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25140": {
		Data: "Disable-UMMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25141": {
		Data: "Disable-UMServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25142": {
		Data: "Dismount-Database Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25143": {
		Data: "Enable-AddressListPaging Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25144": {
		Data: "Enable-AntispamUpdates Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25145": {
		Data: "Enable-CmdletExtensionAgent Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25146": {
		Data: "Enable-DistributionGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25147": {
		Data: "Enable-ExchangeCertificate Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25148": {
		Data: "Enable-InboxRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25149": {
		Data: "Enable-JournalRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25150": {
		Data: "Enable-Mailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25151": {
		Data: "Enable-MailContact Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25152": {
		Data: "Enable-MailPublicFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25153": {
		Data: "Enable-MailUser Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25154": {
		Data: "Enable-OutlookAnywhere Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25155": {
		Data: "Enable-OutlookProtectionRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25156": {
		Data: "Enable-RemoteMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25157": {
		Data: "Enable-ServiceEmailChannel Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25158": {
		Data: "Enable-TransportAgent Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25159": {
		Data: "Enable-TransportRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25160": {
		Data: "Enable-UMAutoAttendant Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25161": {
		Data: "Enable-UMIPGateway Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25162": {
		Data: "Enable-UMMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25163": {
		Data: "Enable-UMServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25164": {
		Data: "Export-ActiveSyncLog Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25165": {
		Data: "Export-AutoDiscoverConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25166": {
		Data: "Export-ExchangeCertificate Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25167": {
		Data: "Export-JournalRuleCollection Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25168": {
		Data: "Export-MailboxDiagnosticLogs Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25169": {
		Data: "Export-Message Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25170": {
		Data: "Export-RecipientDataProperty Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25171": {
		Data: "Export-TransportRuleCollection Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25172": {
		Data: "Export-UMCallDataRecord Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25173": {
		Data: "Export-UMPrompt Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25174": {
		Data: "Import-ExchangeCertificate Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25175": {
		Data: "Import-JournalRuleCollection Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25176": {
		Data: "Import-RecipientDataProperty Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25177": {
		Data: "Import-TransportRuleCollection Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25178": {
		Data: "Import-UMPrompt Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25179": {
		Data: "Install-TransportAgent Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25180": {
		Data: "Mount-Database Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25181": {
		Data: "Move-ActiveMailboxDatabase Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25182": {
		Data: "Move-AddressList Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25183": {
		Data: "Move-DatabasePath Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25184": {
		Data: "Move-OfflineAddressBook Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25185": {
		Data: "New-AcceptedDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25186": {
		Data: "New-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25187": {
		Data: "New-ActiveSyncMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25188": {
		Data: "New-ActiveSyncVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25189": {
		Data: "New-AddressList Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25190": {
		Data: "New-AdminAuditLogSearch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25191": {
		Data: "New-AutodiscoverVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25192": {
		Data: "New-AvailabilityReportOutage Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25193": {
		Data: "New-ClientAccessArray Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25194": {
		Data: "New-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25195": {
		Data: "New-DatabaseAvailabilityGroupNetwork Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25196": {
		Data: "New-DeliveryAgentConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25197": {
		Data: "New-DistributionGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25198": {
		Data: "New-DynamicDistributionGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25199": {
		Data: "New-EcpVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25200": {
		Data: "New-EdgeSubscription Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25201": {
		Data: "New-EdgeSyncServiceConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25202": {
		Data: "New-EmailAddressPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25203": {
		Data: "New-ExchangeCertificate Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25204": {
		Data: "New-FederationTrust Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25205": {
		Data: "New-ForeignConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25206": {
		Data: "New-GlobalAddressList Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25207": {
		Data: "New-InboxRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25208": {
		Data: "New-JournalRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25209": {
		Data: "New-Mailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25210": {
		Data: "New-MailboxAuditLogSearch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25211": {
		Data: "New-MailboxDatabase Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25212": {
		Data: "New-MailboxFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25213": {
		Data: "New-MailboxRepairRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25214": {
		Data: "New-MailboxRestoreRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25215": {
		Data: "New-MailContact Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25216": {
		Data: "New-MailMessage Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25217": {
		Data: "New-MailUser Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25218": {
		Data: "New-ManagedContentSettings Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25219": {
		Data: "New-ManagedFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25220": {
		Data: "New-ManagedFolderMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25221": {
		Data: "New-ManagementRole Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25222": {
		Data: "New-ManagementRoleAssignment Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25223": {
		Data: "New-ManagementScope Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25224": {
		Data: "New-MessageClassification Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25225": {
		Data: "New-MoveRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25226": {
		Data: "New-OabVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25227": {
		Data: "New-OfflineAddressBook Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25228": {
		Data: "New-OrganizationRelationship Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25229": {
		Data: "New-OutlookProtectionRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25230": {
		Data: "New-OutlookProvider Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25231": {
		Data: "New-OwaMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25232": {
		Data: "New-OwaVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25233": {
		Data: "New-PublicFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25234": {
		Data: "New-PublicFolderDatabase Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25235": {
		Data: "New-PublicFolderDatabaseRepairRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25236": {
		Data: "New-ReceiveConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25237": {
		Data: "New-RemoteDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25238": {
		Data: "New-RemoteMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25239": {
		Data: "New-RetentionPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25240": {
		Data: "New-RetentionPolicyTag Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25241": {
		Data: "New-RoleAssignmentPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25242": {
		Data: "New-RoleGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25243": {
		Data: "New-RoutingGroupConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25244": {
		Data: "New-RpcClientAccess Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25245": {
		Data: "New-SendConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25246": {
		Data: "New-SharingPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25247": {
		Data: "New-SystemMessage Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25248": {
		Data: "New-ThrottlingPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25249": {
		Data: "New-TransportRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25250": {
		Data: "New-UMAutoAttendant Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25251": {
		Data: "New-UMDialPlan Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25252": {
		Data: "New-UMHuntGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25253": {
		Data: "New-UMIPGateway Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25254": {
		Data: "New-UMMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25255": {
		Data: "New-WebServicesVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25256": {
		Data: "New-X400AuthoritativeDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25257": {
		Data: "Remove-AcceptedDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25258": {
		Data: "Remove-ActiveSyncDevice Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25259": {
		Data: "Remove-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25260": {
		Data: "Remove-ActiveSyncDeviceClass Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25261": {
		Data: "Remove-ActiveSyncMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25262": {
		Data: "Remove-ActiveSyncVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25263": {
		Data: "Remove-AddressList Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25264": {
		Data: "Remove-ADPermission Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25265": {
		Data: "Remove-AutodiscoverVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25266": {
		Data: "Remove-AvailabilityAddressSpace Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25267": {
		Data: "Remove-AvailabilityReportOutage Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25268": {
		Data: "Remove-ClientAccessArray Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25269": {
		Data: "Remove-ContentFilterPhrase Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25270": {
		Data: "Remove-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25271": {
		Data: "Remove-DatabaseAvailabilityGroupNetwork Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25272": {
		Data: "Remove-DatabaseAvailabilityGroupServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25273": {
		Data: "Remove-DeliveryAgentConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25274": {
		Data: "Remove-DistributionGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25275": {
		Data: "Remove-DistributionGroupMember Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25276": {
		Data: "Remove-DynamicDistributionGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25277": {
		Data: "Remove-EcpVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25278": {
		Data: "Remove-EdgeSubscription Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25279": {
		Data: "Remove-EmailAddressPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25280": {
		Data: "Remove-ExchangeCertificate Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25281": {
		Data: "Remove-FederatedDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25282": {
		Data: "Remove-FederationTrust Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25283": {
		Data: "Remove-ForeignConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25284": {
		Data: "Remove-GlobalAddressList Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25285": {
		Data: "Remove-InboxRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25286": {
		Data: "Remove-IPAllowListEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25287": {
		Data: "Remove-IPAllowListProvider Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25288": {
		Data: "Remove-IPBlockListEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25289": {
		Data: "Remove-IPBlockListProvider Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25290": {
		Data: "Remove-JournalRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25291": {
		Data: "Remove-Mailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25292": {
		Data: "Remove-MailboxDatabase Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25293": {
		Data: "Remove-MailboxDatabaseCopy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25294": {
		Data: "Remove-MailboxFolderPermission Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25295": {
		Data: "Remove-MailboxPermission Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25296": {
		Data: "Remove-MailboxRestoreRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25297": {
		Data: "Remove-MailContact Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25298": {
		Data: "Remove-MailUser Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25299": {
		Data: "Remove-ManagedContentSettings Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25300": {
		Data: "Remove-ManagedFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25301": {
		Data: "Remove-ManagedFolderMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25302": {
		Data: "Remove-ManagementRole Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25303": {
		Data: "Remove-ManagementRoleAssignment Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25304": {
		Data: "Remove-ManagementRoleEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25305": {
		Data: "Remove-ManagementScope Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25306": {
		Data: "Remove-Message Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25307": {
		Data: "Remove-MessageClassification Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25308": {
		Data: "Remove-MoveRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25309": {
		Data: "Remove-OabVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25310": {
		Data: "Remove-OfflineAddressBook Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25311": {
		Data: "Remove-OrganizationRelationship Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25312": {
		Data: "Remove-OutlookProtectionRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25313": {
		Data: "Remove-OutlookProvider Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25314": {
		Data: "Remove-OwaMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25315": {
		Data: "Remove-OwaVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25316": {
		Data: "Remove-PublicFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25317": {
		Data: "Remove-PublicFolderAdministrativePermission Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25318": {
		Data: "Remove-PublicFolderClientPermission Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25319": {
		Data: "Remove-PublicFolderDatabase Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25320": {
		Data: "Remove-ReceiveConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25321": {
		Data: "Remove-RemoteDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25322": {
		Data: "Remove-RemoteMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25323": {
		Data: "Remove-RetentionPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25324": {
		Data: "Remove-RetentionPolicyTag Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25325": {
		Data: "Remove-RoleAssignmentPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25326": {
		Data: "Remove-RoleGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25327": {
		Data: "Remove-RoleGroupMember Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25328": {
		Data: "Remove-RoutingGroupConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25329": {
		Data: "Remove-RpcClientAccess Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25330": {
		Data: "Remove-SendConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25331": {
		Data: "Remove-SharingPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25332": {
		Data: "Remove-StoreMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25333": {
		Data: "Remove-SystemMessage Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25334": {
		Data: "Remove-ThrottlingPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25335": {
		Data: "Remove-TransportRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25336": {
		Data: "Remove-UMAutoAttendant Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25337": {
		Data: "Remove-UMDialPlan Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25338": {
		Data: "Remove-UMHuntGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25339": {
		Data: "Remove-UMIPGateway Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25340": {
		Data: "Remove-UMMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25341": {
		Data: "Remove-WebServicesVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25342": {
		Data: "Remove-X400AuthoritativeDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25343": {
		Data: "Restore-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25344": {
		Data: "Restore-DetailsTemplate Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25345": {
		Data: "Restore-Mailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25346": {
		Data: "Resume-MailboxDatabaseCopy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25347": {
		Data: "Resume-MailboxExportRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25348": {
		Data: "Resume-MailboxRestoreRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25349": {
		Data: "Resume-Message Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25350": {
		Data: "Resume-MoveRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25351": {
		Data: "Resume-PublicFolderReplication Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25352": {
		Data: "Resume-Queue Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25353": {
		Data: "Retry-Queue Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25354": {
		Data: "Send-TextMessagingVerificationCode Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25355": {
		Data: "Set-AcceptedDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25356": {
		Data: "Set-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25357": {
		Data: "Set-ActiveSyncMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25358": {
		Data: "Set-ActiveSyncOrganizationSettings Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25359": {
		Data: "Set-ActiveSyncVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25360": {
		Data: "Set-AddressList Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25361": {
		Data: "Set-AdminAuditLogConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25362": {
		Data: "Set-ADServerSettings Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25363": {
		Data: "Set-ADSite Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25364": {
		Data: "Set-AdSiteLink Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25365": {
		Data: "Set-AutodiscoverVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25366": {
		Data: "Set-AvailabilityConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25367": {
		Data: "Set-AvailabilityReportOutage Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25368": {
		Data: "Set-CalendarNotification Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25369": {
		Data: "Set-CalendarProcessing Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25370": {
		Data: "Set-CASMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25371": {
		Data: "Set-ClientAccessArray Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25372": {
		Data: "Set-ClientAccessServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25373": {
		Data: "Set-CmdletExtensionAgent Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25374": {
		Data: "Set-Contact Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25375": {
		Data: "Set-ContentFilterConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25376": {
		Data: "Set-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25377": {
		Data: "Set-DatabaseAvailabilityGroupNetwork Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25378": {
		Data: "Set-DeliveryAgentConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25379": {
		Data: "Set-DetailsTemplate Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25380": {
		Data: "Set-DistributionGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25381": {
		Data: "Set-DynamicDistributionGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25382": {
		Data: "Set-EcpVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25383": {
		Data: "Set-EdgeSyncServiceConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25384": {
		Data: "Set-EmailAddressPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25385": {
		Data: "Set-EventLogLevel Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25386": {
		Data: "Set-ExchangeAssistanceConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25387": {
		Data: "Set-ExchangeServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25388": {
		Data: "Set-FederatedOrganizationIdentifier Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25389": {
		Data: "Set-FederationTrust Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25390": {
		Data: "Set-ForeignConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25391": {
		Data: "Set-GlobalAddressList Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25392": {
		Data: "Set-Group Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25393": {
		Data: "Set-ImapSettings Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25394": {
		Data: "Set-InboxRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25395": {
		Data: "Set-IPAllowListConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25396": {
		Data: "Set-IPAllowListProvider Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25397": {
		Data: "Set-IPAllowListProvidersConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25398": {
		Data: "Set-IPBlockListConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25399": {
		Data: "Set-IPBlockListProvider Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25400": {
		Data: "Set-IPBlockListProvidersConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25401": {
		Data: "Set-IRMConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25402": {
		Data: "Set-JournalRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25403": {
		Data: "Set-Mailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25404": {
		Data: "Set-MailboxAuditBypassAssociation Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25405": {
		Data: "Set-MailboxAutoReplyConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25406": {
		Data: "Set-MailboxCalendarConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25407": {
		Data: "Set-MailboxCalendarFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25408": {
		Data: "Set-MailboxDatabase Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25409": {
		Data: "Set-MailboxDatabaseCopy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25410": {
		Data: "Set-MailboxFolderPermission Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25411": {
		Data: "Set-MailboxJunkEmailConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25412": {
		Data: "Set-MailboxMessageConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25413": {
		Data: "Set-MailboxRegionalConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25414": {
		Data: "Set-MailboxRestoreRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25415": {
		Data: "Set-MailboxServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25416": {
		Data: "Set-MailboxSpellingConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25417": {
		Data: "Set-MailContact Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25418": {
		Data: "Set-MailPublicFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25419": {
		Data: "Set-MailUser Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25420": {
		Data: "Set-ManagedContentSettings Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25421": {
		Data: "Set-ManagedFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25422": {
		Data: "Set-ManagedFolderMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25423": {
		Data: "Set-ManagementRoleAssignment Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25424": {
		Data: "Set-ManagementRoleEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25425": {
		Data: "Set-ManagementScope Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25426": {
		Data: "Set-MessageClassification Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25427": {
		Data: "Set-MoveRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25428": {
		Data: "Set-OabVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25429": {
		Data: "Set-OfflineAddressBook Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25430": {
		Data: "Set-OrganizationConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25431": {
		Data: "Set-OrganizationRelationship Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25432": {
		Data: "Set-OutlookAnywhere Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25433": {
		Data: "Set-OutlookProtectionRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25434": {
		Data: "Set-OutlookProvider Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25435": {
		Data: "Set-OwaMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25436": {
		Data: "Set-OwaVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25437": {
		Data: "Set-PopSettings Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25438": {
		Data: "Set-PowerShellVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25439": {
		Data: "Set-PublicFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25440": {
		Data: "Set-PublicFolderDatabase Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25441": {
		Data: "Set-ReceiveConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25442": {
		Data: "Set-RecipientFilterConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25443": {
		Data: "Set-RemoteDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25444": {
		Data: "Set-RemoteMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25445": {
		Data: "Set-ReOrigConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25446": {
		Data: "Set-RetentionPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25447": {
		Data: "Set-RetentionPolicyTag Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25448": {
		Data: "Set-RoleAssignmentPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25449": {
		Data: "Set-RoleGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25450": {
		Data: "Set-RoutingGroupConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25451": {
		Data: "Set-RpcClientAccess Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25452": {
		Data: "Set-SendConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25453": {
		Data: "Set-SenderFilterConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25454": {
		Data: "Set-SenderIdConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25455": {
		Data: "Set-SenderReputationConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25456": {
		Data: "Set-SharingPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25457": {
		Data: "Set-SystemMessage Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25458": {
		Data: "Set-TextMessagingAccount Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25459": {
		Data: "Set-ThrottlingPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25460": {
		Data: "Set-ThrottlingPolicyAssociation Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25461": {
		Data: "Set-TransportAgent Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25462": {
		Data: "Set-TransportConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25463": {
		Data: "Set-TransportRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25464": {
		Data: "Set-TransportServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25465": {
		Data: "Set-UMAutoAttendant Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25466": {
		Data: "Set-UMDialPlan Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25467": {
		Data: "Set-UMIPGateway Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25468": {
		Data: "Set-UMMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25469": {
		Data: "Set-UMMailboxPIN Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25470": {
		Data: "Set-UMMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25471": {
		Data: "Set-UmServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25472": {
		Data: "Set-User Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25473": {
		Data: "Set-WebServicesVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25474": {
		Data: "Set-X400AuthoritativeDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25475": {
		Data: "Start-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25476": {
		Data: "Start-EdgeSynchronization Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25477": {
		Data: "Start-ManagedFolderAssistant Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25478": {
		Data: "Start-RetentionAutoTagLearning Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25479": {
		Data: "Stop-DatabaseAvailabilityGroup Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25480": {
		Data: "Stop-ManagedFolderAssistant Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25481": {
		Data: "Suspend-MailboxDatabaseCopy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25482": {
		Data: "Suspend-MailboxRestoreRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25483": {
		Data: "Suspend-Message Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25484": {
		Data: "Suspend-MoveRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25485": {
		Data: "Suspend-PublicFolderReplication Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25486": {
		Data: "Suspend-Queue Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25487": {
		Data: "Test-ActiveSyncConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25488": {
		Data: "Test-AssistantHealth Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25489": {
		Data: "Test-CalendarConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25490": {
		Data: "Test-EcpConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25491": {
		Data: "Test-EdgeSynchronization Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25492": {
		Data: "Test-ExchangeSearch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25493": {
		Data: "Test-FederationTrust Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25494": {
		Data: "Test-FederationTrustCertificate Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25495": {
		Data: "Test-ImapConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25496": {
		Data: "Test-IPAllowListProvider Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25497": {
		Data: "Test-IPBlockListProvider Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25498": {
		Data: "Test-IRMConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25499": {
		Data: "Test-Mailflow Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25500": {
		Data: "Test-MAPIConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25501": {
		Data: "Test-MRSHealth Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25502": {
		Data: "Test-OrganizationRelationship Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25503": {
		Data: "Test-OutlookConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25504": {
		Data: "Test-OutlookWebServices Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25505": {
		Data: "Test-OwaConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25506": {
		Data: "Test-PopConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25507": {
		Data: "Test-PowerShellConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25508": {
		Data: "Test-ReplicationHealth Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25509": {
		Data: "Test-SenderId Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25510": {
		Data: "Test-ServiceHealth Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25511": {
		Data: "Test-SmtpConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25512": {
		Data: "Test-SystemHealth Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25513": {
		Data: "Test-UMConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25514": {
		Data: "Test-WebServicesConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25515": {
		Data: "Uninstall-TransportAgent Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25516": {
		Data: "Update-AddressList Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25517": {
		Data: "Update-DistributionGroupMember Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25518": {
		Data: "Update-EmailAddressPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25519": {
		Data: "Update-FileDistributionService Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25520": {
		Data: "Update-GlobalAddressList Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25521": {
		Data: "Update-MailboxDatabaseCopy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25522": {
		Data: "Update-OfflineAddressBook Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25523": {
		Data: "Update-PublicFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25524": {
		Data: "Update-PublicFolderHierarchy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25525": {
		Data: "Update-Recipient Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25526": {
		Data: "Update-RoleGroupMember Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25527": {
		Data: "Update-SafeList Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25528": {
		Data: "Write-AdminAuditLog Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25529": {
		Data: "Add-GlobalMonitoringOverride Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25530": {
		Data: "Add-ResubmitRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25531": {
		Data: "Add-ServerMonitoringOverride Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25532": {
		Data: "Clear-MobileDevice Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25533": {
		Data: "Complete-MigrationBatch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25534": {
		Data: "Disable-App Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25535": {
		Data: "Disable-MailboxQuarantine Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25536": {
		Data: "Disable-UMCallAnsweringRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25537": {
		Data: "Disable-UMService Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25538": {
		Data: "Dump-ProvisioningCache Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25539": {
		Data: "Enable-App Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25540": {
		Data: "Enable-MailboxQuarantine Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25541": {
		Data: "Enable-UMCallAnsweringRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25542": {
		Data: "Enable-UMService Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25543": {
		Data: "Export-DlpPolicyCollection Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25544": {
		Data: "Export-MigrationReport Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25545": {
		Data: "Import-DlpPolicyCollection Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25546": {
		Data: "Import-DlpPolicyTemplate Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25547": {
		Data: "Invoke-MonitoringProbe Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25548": {
		Data: "New-AddressBookPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25549": {
		Data: "New-App Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25550": {
		Data: "New-AuthServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25551": {
		Data: "New-ClassificationRuleCollection Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25552": {
		Data: "New-DlpPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25553": {
		Data: "New-HybridConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25554": {
		Data: "New-MailboxExportRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25555": {
		Data: "New-MailboxImportRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25556": {
		Data: "New-MailboxSearch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25557": {
		Data: "New-MalwareFilterPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25558": {
		Data: "New-MigrationBatch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25559": {
		Data: "New-MigrationEndpoint Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25560": {
		Data: "New-MobileDeviceMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25561": {
		Data: "New-OnPremisesOrganization Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25562": {
		Data: "New-PartnerApplication Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25563": {
		Data: "New-PolicyTipConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25564": {
		Data: "New-PowerShellVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25565": {
		Data: "New-PublicFolderMigrationRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25566": {
		Data: "New-ReOrigPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25567": {
		Data: "New-SiteMailboxProvisioningPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25568": {
		Data: "New-SyncMailPublicFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25569": {
		Data: "New-UMCallAnsweringRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25570": {
		Data: "New-WorkloadManagementPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25571": {
		Data: "New-WorkloadPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25572": {
		Data: "Redirect-Message Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25573": {
		Data: "Remove-AddressBookPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25574": {
		Data: "Remove-App Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25575": {
		Data: "Remove-AuthServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25576": {
		Data: "Remove-ClassificationRuleCollection Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25577": {
		Data: "Remove-DlpPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25578": {
		Data: "Remove-DlpPolicyTemplate Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25579": {
		Data: "Remove-GlobalMonitoringOverride Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25580": {
		Data: "Remove-HybridConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25581": {
		Data: "Remove-LinkedUser Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25582": {
		Data: "Remove-MailboxExportRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25583": {
		Data: "Remove-MailboxImportRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25584": {
		Data: "Remove-MailboxSearch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25585": {
		Data: "Remove-MalwareFilterPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25586": {
		Data: "Remove-MalwareFilterRecoveryItem Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25587": {
		Data: "Remove-MigrationBatch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25588": {
		Data: "Remove-MigrationEndpoint Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25589": {
		Data: "Remove-MigrationUser Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25590": {
		Data: "Remove-MobileDevice Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25591": {
		Data: "Remove-MobileDeviceMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25592": {
		Data: "Remove-OnPremisesOrganization Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25593": {
		Data: "Remove-PartnerApplication Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25594": {
		Data: "Remove-PolicyTipConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25595": {
		Data: "Remove-PowerShellVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25596": {
		Data: "Remove-PublicFolderMigrationRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25597": {
		Data: "Remove-ReOrigPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25598": {
		Data: "Remove-ResubmitRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25599": {
		Data: "Remove-SiteMailboxProvisioningPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25600": {
		Data: "Remove-UMCallAnsweringRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25601": {
		Data: "Remove-UserPhoto Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25602": {
		Data: "Remove-WorkloadManagementPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25603": {
		Data: "Remove-WorkloadPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25604": {
		Data: "Reset-ProvisioningCache Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25605": {
		Data: "Resume-MailboxImportRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25606": {
		Data: "Resume-MalwareFilterRecoveryItem Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25607": {
		Data: "Resume-PublicFolderMigrationRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25608": {
		Data: "Set-ActiveSyncDeviceAccessRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25609": {
		Data: "Set-AddressBookPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25610": {
		Data: "Set-App Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25611": {
		Data: "Set-AuthConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25612": {
		Data: "Set-AuthServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25613": {
		Data: "Set-ClassificationRuleCollection Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25614": {
		Data: "Set-DlpPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25615": {
		Data: "Set-FrontendTransportService Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25616": {
		Data: "Set-HybridConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25617": {
		Data: "Set-HybridMailflow Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25618": {
		Data: "Set-MailboxExportRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25619": {
		Data: "Set-MailboxImportRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25620": {
		Data: "Set-MailboxSearch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25621": {
		Data: "Set-MailboxTransportService Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25622": {
		Data: "Set-MalwareFilteringServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25623": {
		Data: "Set-MalwareFilterPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25624": {
		Data: "Set-MigrationBatch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25625": {
		Data: "Set-MigrationConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25626": {
		Data: "Set-MigrationEndpoint Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25627": {
		Data: "Set-MobileDeviceMailboxPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25628": {
		Data: "Set-Notification Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25629": {
		Data: "Set-OnPremisesOrganization Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25630": {
		Data: "Set-PartnerApplication Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25631": {
		Data: "Set-PendingFederatedDomain Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25632": {
		Data: "Set-PolicyTipConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25633": {
		Data: "Set-PublicFolderMigrationRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25634": {
		Data: "Set-ReOrigPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25635": {
		Data: "Set-ResubmitRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25636": {
		Data: "Set-RMSTemplate Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25637": {
		Data: "Set-ServerComponentState Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25638": {
		Data: "Set-ServerMonitor Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25639": {
		Data: "Set-SiteMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25640": {
		Data: "Set-SiteMailboxProvisioningPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25641": {
		Data: "Set-TransportService Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25642": {
		Data: "Set-UMCallAnsweringRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25643": {
		Data: "Set-UMCallRouterSettings Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25644": {
		Data: "Set-UMService Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25645": {
		Data: "Set-UserPhoto Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25646": {
		Data: "Set-WorkloadPolicy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25647": {
		Data: "Start-MailboxSearch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25648": {
		Data: "Start-MigrationBatch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25649": {
		Data: "Stop-MailboxSearch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25650": {
		Data: "Stop-MigrationBatch Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25651": {
		Data: "Suspend-MailboxExportRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25652": {
		Data: "Suspend-MailboxImportRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25653": {
		Data: "Suspend-PublicFolderMigrationRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25654": {
		Data: "Test-ArchiveConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25655": {
		Data: "Test-MigrationServerAvailability Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25656": {
		Data: "Test-OAuthConnectivity Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25657": {
		Data: "Test-SiteMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25658": {
		Data: "Update-HybridConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25659": {
		Data: "Update-PublicFolderMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25660": {
		Data: "Update-SiteMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25661": {
		Data: "Add-AttachmentFilterEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25662": {
		Data: "Remove-AttachmentFilterEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25663": {
		Data: "New-AddressRewriteEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25664": {
		Data: "Remove-AddressRewriteEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25665": {
		Data: "Set-AddressRewriteEntry Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25666": {
		Data: "Set-AttachmentFilterListConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25667": {
		Data: "Set-MailboxSentItemsConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25668": {
		Data: "Update-MovedMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25669": {
		Data: "Disable-MalwareFilterRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25670": {
		Data: "Enable-MalwareFilterRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25671": {
		Data: "New-MalwareFilterRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25672": {
		Data: "Remove-MalwareFilterRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25673": {
		Data: "Set-MalwareFilterRule Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25674": {
		Data: "Remove-MailboxRepairRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25675": {
		Data: "Remove-ServerMonitoringOverride Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25676": {
		Data: "Update-ExchangeHelp Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25677": {
		Data: "Update-StoreMailboxState Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25678": {
		Data: "Disable-PushNotificationProxy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25679": {
		Data: "Enable-PushNotificationProxy Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25680": {
		Data: "New-PublicFolderMoveRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25681": {
		Data: "Remove-PublicFolderMoveRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25682": {
		Data: "Resume-PublicFolderMoveRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25683": {
		Data: "Set-PublicFolderMoveRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25684": {
		Data: "Suspend-PublicFolderMoveRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25685": {
		Data: "Update-DatabaseSchema Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25686": {
		Data: "Set-SearchDocumentFormat Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25687": {
		Data: "New-AuthRedirect Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25688": {
		Data: "New-CompliancePolicySyncNotification Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25689": {
		Data: "New-ComplianceServiceVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25690": {
		Data: "New-DatabaseAvailabilityGroupConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25691": {
		Data: "New-DataClassification Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25692": {
		Data: "New-Fingerprint Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25693": {
		Data: "New-IntraOrganizationConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25694": {
		Data: "New-MailboxDeliveryVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25695": {
		Data: "New-MapiVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25696": {
		Data: "New-OutlookServiceVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25697": {
		Data: "New-RestVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25698": {
		Data: "New-SearchDocumentFormat Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25699": {
		Data: "New-SettingOverride Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25700": {
		Data: "New-SiteMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25701": {
		Data: "Remove-AuthRedirect Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25702": {
		Data: "Remove-CompliancePolicySyncNotification Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25703": {
		Data: "Remove-ComplianceServiceVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25704": {
		Data: "Remove-DatabaseAvailabilityGroupConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25705": {
		Data: "Remove-DataClassification Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25706": {
		Data: "Remove-IntraOrganizationConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25707": {
		Data: "Remove-MailboxDeliveryVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25708": {
		Data: "Remove-MapiVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25709": {
		Data: "Remove-OutlookServiceVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25710": {
		Data: "Remove-PublicFolderMailboxMigrationRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25711": {
		Data: "Remove-PushNotificationSubscription Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25712": {
		Data: "Remove-RestVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25713": {
		Data: "Remove-SearchDocumentFormat Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25714": {
		Data: "Remove-SettingOverride Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25715": {
		Data: "Remove-SyncMailPublicFolder Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25716": {
		Data: "Resume-PublicFolderMailboxMigrationRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25717": {
		Data: "Send-MapiSubmitSystemProbe Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25718": {
		Data: "Set-AuthRedirect Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25719": {
		Data: "Set-ClientAccessService Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25720": {
		Data: "Set-Clutter Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25721": {
		Data: "Set-ComplianceServiceVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25722": {
		Data: "Set-ConsumerMailbox Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25723": {
		Data: "Set-DatabaseAvailabilityGroupConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25724": {
		Data: "Set-DataClassification Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25725": {
		Data: "Set-IntraOrganizationConnector Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25726": {
		Data: "Set-LogExportVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25727": {
		Data: "Set-MailboxDeliveryVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25728": {
		Data: "Set-MapiVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25729": {
		Data: "Set-OutlookServiceVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25730": {
		Data: "Set-PublicFolderMailboxMigrationRequest Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25731": {
		Data: "Set-RestVirtualDirectory Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25732": {
		Data: "Set-SettingOverride Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25733": {
		Data: "Set-SmimeConfig Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25734": {
		Data: "Set-SubmissionMalwareFilteringServer Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25735": {
		Data: "Set-UMMailboxConfiguration Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25736": {
		Data: "Set-UnifiedAuditSetting Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25737": {
		Data: "Start-AuditAssistant Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25738": {
		Data: "Start-UMPhoneSession Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25739": {
		Data: "Stop-UMPhoneSession Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25740": {
		Data: "Test-DataClassification Exchange cmdlet issued",
		Orig: "Exchange",
	},
	"25741": {
		Data: "Test-TextExtraction Exchange cmdlet issued",
		Orig: "Exchange",
	},
}

var re = regexp.MustCompile(`"Computer":\s*"(?P<host>.*?)".*"(EventID|Value)":\s*"(?P<data>.*?)".*"SystemTime":\s*"(?P<time>.+?)"`)

func Extract(line string) *extract.Event {
	m := re.FindStringSubmatch(line)

	if len(m) < 4 {
		return nil
	}

	x := re.SubexpIndex("time")
	y := re.SubexpIndex("host")
	z := re.SubexpIndex("data")

	ts, err := time.Parse(time.RFC3339, m[x])

	if err != nil {
		return nil
	}

	ev := &extract.Event{
		Time: ts.UTC(),
		Host: m[y],
		Data: "EventID " + m[z],
	}

	if e, ok := events[m[z]]; ok {
		ev.Data = e.Data
	}

	return ev
}
