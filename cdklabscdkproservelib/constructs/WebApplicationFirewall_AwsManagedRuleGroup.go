package constructs


// WAF Managed Rule Groups.
// Experimental.
type WebApplicationFirewall_AwsManagedRuleGroup string

const (
	// Contains rules that are generally applicable to web applications.
	//
	// This provides protection against exploitation of a wide range of vulnerabilities, including those described in OWASP publications.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_COMMON_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "COMMON_RULE_SET"
	// Contains rules that allow you to block external access to exposed admin pages.
	//
	// This may be useful if you are running third-party software or would like to reduce the risk of a malicious actor gaining administrative access to your application.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_ADMIN_PROTECTION_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "ADMIN_PROTECTION_RULE_SET"
	// Contains rules that allow you to block request patterns that are known to be invalid and are associated with exploitation or discovery of vulnerabilities.
	//
	// This can help reduce the risk of a malicious actor discovering a vulnerable application.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_KNOWN_BAD_INPUTS_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "KNOWN_BAD_INPUTS_RULE_SET"
	// Contains rules that allow you to block request patterns associated with exploitation of SQL databases, like SQL injection attacks.
	//
	// This can help prevent remote injection of unauthorized queries.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_SQL_DATABASE_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "SQL_DATABASE_RULE_SET"
	// Contains rules that block request patterns associated with exploitation of vulnerabilities specific to Linux, including LFI attacks.
	//
	// This can help prevent attacks that expose file contents or execute code for which the attacker should not have had access.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_LINUX_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "LINUX_RULE_SET"
	// Contains rules that block request patterns associated with exploiting vulnerabilities specific to POSIX/POSIX-like OS, including LFI attacks.
	//
	// This can help prevent attacks that expose file contents or execute code for which access should not been allowed.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_UNIX_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "UNIX_RULE_SET"
	// Contains rules that block request patterns associated with exploiting vulnerabilities specific to Windows, (e.g., PowerShell commands). This can help prevent exploits that allow attacker to run unauthorized commands or execute malicious code.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_WINDOWS_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "WINDOWS_RULE_SET"
	// Contains rules that block request patterns associated with exploiting vulnerabilities specific to the use of the PHP, including injection of unsafe PHP functions.
	//
	// This can help prevent exploits that allow an attacker to remotely execute code or commands.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_PHP_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "PHP_RULE_SET"
	// The WordPress Applications group contains rules that block request patterns associated with the exploitation of vulnerabilities specific to WordPress sites.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_WORD_PRESS_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "WORD_PRESS_RULE_SET"
	// This group contains rules that are based on Amazon threat intelligence.
	//
	// This is useful if you would like to block sources associated with bots or other threats.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_AMAZON_IP_REPUTATION_LIST WebApplicationFirewall_AwsManagedRuleGroup = "AMAZON_IP_REPUTATION_LIST"
	// This group contains rules that allow you to block requests from services that allow obfuscation of viewer identity.
	//
	// This can include request originating from VPN, proxies, Tor nodes, and hosting providers. This is useful if you want to filter out viewers that may be trying to hide their identity from your application.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_ANONYMOUS_IP_LIST WebApplicationFirewall_AwsManagedRuleGroup = "ANONYMOUS_IP_LIST"
	// Provides protection against automated bots that can consume excess resources, skew business metrics, cause downtime, or perform malicious activities.
	//
	// Bot Control provides additional visibility through Amazon CloudWatch and generates labels that you can use to control bot traffic to your applications.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_BOT_CONTROL_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "BOT_CONTROL_RULE_SET"
	// Provides protection for your login page against stolen credentials, credential stuffing attacks, brute force login attempts, and other anomalous login activities.
	//
	// With account takeover prevention, you can prevent unauthorized access that may lead to fraudulent activities, or inform legitimate users to take a preventive action.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_ATP_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "ATP_RULE_SET"
	// Provides protection against the creation of fraudulent accounts on your site.
	//
	// Fraudulent accounts can be used for activities such as obtaining sign-up bonuses and impersonating legitimate users.
	// Experimental.
	WebApplicationFirewall_AwsManagedRuleGroup_ACFP_RULE_SET WebApplicationFirewall_AwsManagedRuleGroup = "ACFP_RULE_SET"
)

