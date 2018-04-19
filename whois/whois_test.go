package whois

import "testing"

var afiliasData = `Domain Name: OGUZHAN.INFO
Registry Domain ID: D2644416-LRMS
Registrar WHOIS Server: whois.godaddy.com
Registrar URL: http://www.godaddy.com
Updated Date: 2018-04-11T07:06:49Z
Creation Date: 2003-04-19T14:13:48Z
Registry Expiry Date: 2019-04-19T14:13:48Z
Registrar Registration Expiration Date:
Registrar: GoDaddy.com, LLC
Registrar IANA ID: 146
Registrar Abuse Contact Email: abuse@godaddy.com
Registrar Abuse Contact Phone: +1.4806242505
Reseller:
Domain Status: clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited
Domain Status: clientRenewProhibited https://icann.org/epp#clientRenewProhibited
Domain Status: clientTransferProhibited https://icann.org/epp#clientTransferProhibited
Domain Status: clientUpdateProhibited https://icann.org/epp#clientUpdateProhibited
Registry Registrant ID: C71042308-LRMS
Registrant Name: Oguzhan YILMAZ
Registrant Organization: MaestroPanel
Registrant Street: Istanbul Universitesi Avcilar Kampusu
Registrant Street: Teknokent Binasi  Z19
Registrant City: istanbul
Registrant State/Province:
Registrant Postal Code: 34524
Registrant Country: TR
Registrant Phone: +90.8508850004
Registrant Phone Ext:
Registrant Fax:
Registrant Fax Ext:
Registrant Email: aspsrc@gmail.com
Registry Admin ID: C71042314-LRMS
Admin Name: Oguzhan YILMAZ
Admin Organization: MaestroPanel
Admin Street: Istanbul Universitesi Avcilar Kampusu
Admin Street: Teknokent Binasi  Z19
Admin City: istanbul
Admin State/Province:
Admin Postal Code: 34524
Admin Country: TR
Admin Phone: +90.8508850004
Admin Phone Ext:
Admin Fax:
Admin Fax Ext:
Admin Email: aspsrc@gmail.com
Registry Tech ID: C71042311-LRMS
Tech Name: Oguzhan YILMAZ
Tech Organization: MaestroPanel
Tech Street: Istanbul Universitesi Avcilar Kampusu
Tech Street: Teknokent Binasi  Z19
Tech City: istanbul
Tech State/Province:
Tech Postal Code: 34524
Tech Country: TR
Tech Phone: +90.8508850004
Tech Phone Ext:
Tech Fax:
Tech Fax Ext:
Tech Email: aspsrc@gmail.com
Registry Billing ID: C71042317-LRMS
Billing Name: Oguzhan YILMAZ
Billing Organization: MaestroPanel
Billing Street: Istanbul Universitesi Avcilar Kampusu
Billing Street: Teknokent Binasi  Z19
Billing City: istanbul
Billing State/Province:
Billing Postal Code: 34524
Billing Country: TR
Billing Phone: +90.8508850004
Billing Phone Ext:
Billing Fax:
Billing Fax Ext:
Billing Email: aspsrc@gmail.com
Name Server: NS22.DOMAINCONTROL.COM
Name Server: NS21.DOMAINCONTROL.COM
DNSSEC: unsigned
URL of the ICANN Whois Inaccuracy Complaint Form: https://www.icann.org/wicf/
>>> Last update of WHOIS database: 2018-04-19T21:49:48Z <<<

For more information on Whois status codes, please visit https://icann.org/epp

Access to AFILIAS WHOIS information is provided to assist persons in determining the contents of a domain name registration record in the Afilias registry database. Thedata in this record is provided by Afilias Limited for informational purposes only, and Afilias does not guarantee its accuracy.  This service is intended only for query-based access. You agree that you will use this data only for lawful purposes and that, under no circumstances will you use this data to(a) allow, enable, or otherwise support the transmission by e-mail, telephone, or facsimile of mass unsolicited, commercial advertising or solicitations to entities other than the data recipient's own existing customers; or (b) enable high volume, automated, electronic processes that send queries or data to the systems of Registry Operator, a Registrar, or Afilias except as reasonably necessary to register domain names or modify existing registrations. All rights reserved. Afilias reserves the right to modify these terms at any time. By submitting this query, you agree to abide by this policy.`

var verisignData = `    Domain Name: MAESTROPANEL.COM                   Registry Domain ID: 1624403870_DOMAIN_COM-VRSN                   Registrar WHOIS Server: whois.godaddy.com                   Registrar URL: http://www.godaddy.com                   Updated Date: 2017-10-24T15:01:16Z
Creation Date: 2010-11-08T10:29:25Z
Registry Expiry Date: 2018-11-08T10:29:25Z
Registrar: GoDaddy.com, LLC
Registrar IANA ID: 146
Registrar Abuse Contact Email: abuse@godaddy.com
Registrar Abuse Contact Phone: 480-624-2505
Domain Status: clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited
Domain Status: clientRenewProhibited https://icann.org/epp#clientRenewProhibited
Domain Status: clientTransferProhibited https://icann.org/epp#clientTransferProhibited
Domain Status: clientUpdateProhibited https://icann.org/epp#clientUpdateProhibited
Name Server: EVA.NS.CLOUDFLARE.COM
Name Server: WILL.NS.CLOUDFLARE.COM
DNSSEC: unsigned
URL of the ICANN Whois Inaccuracy Complaint Form: https://www.icann.org/wicf/
>>> Last update of whois database: 2018-04-19T22:36:36Z <<<

For more information on Whois status codes, please visit https://icann.org/epp

NOTICE: The expiration date displayed in this record is the date the
registrar's sponsorship of the domain name registration in the registry is
currently set to expire. This date does not necessarily reflect the expiration
date of the domain name registrant's agreement with the sponsoring
registrar.  Users may consult the sponsoring registrar's Whois database to
view the registrar's reported date of expiration for this registration.

TERMS OF USE: You are not authorized to access or query our Whois
database through the use of electronic processes that are high-volume and
automated except as reasonably necessary to register domain names or
modify existing registrations; the Data in VeriSign Global Registry
Services' ("VeriSign") Whois database is provided by VeriSign for
information purposes only, and to assist persons in obtaining information
about or related to a domain name registration record. VeriSign does not
guarantee its accuracy. By submitting a Whois query, you agree to abide
by the following terms of use: You agree that you may use this Data only
for lawful purposes and that under no circumstances will you use this Data
to: (1) allow, enable, or otherwise support the transmission of mass
unsolicited, commercial advertising or solicitations via e-mail, telephone,
or facsimile; or (2) enable high volume, automated, electronic processes
that apply to VeriSign (or its computer systems). The compilation,
repackaging, dissemination or other use of this Data is expressly
prohibited without the prior written consent of VeriSign. You agree not to
use electronic processes that are automated and high-volume to access or
query the Whois database except as reasonably necessary to register
domain names or modify existing registrations. VeriSign reserves the right
to restrict your access to the Whois database in its sole discretion to ensure
operational stability.  VeriSign may restrict or terminate your access to the
Whois database for failure to abide by these terms of use. VeriSign
reserves the right to modify these terms at any time.

The Registry database contains ONLY .COM, .NET, .EDU domains and
Registrars.`

func tTestWhoisConn(t *testing.T) {
	data, err := whoisConn("whois.verisign-grs.com", "maestropanel.com")

	if err != nil {
		t.Error(err)
	}

	t.Log(data)
}

func TestParserNameServer(t *testing.T) {

	list := parseNameServers(afiliasData)

	if len(list) != 2 {
		t.Error("invalid name server list")
	}
}

func TestParserRegistrarWhoisServer(t *testing.T) {

	server := parserRegistatWhoisServer(afiliasData)

	if server == "" {
		t.Error("registrar whois server not found")
	}

	if server != "whois.godaddy.com" {
		t.Error("not expected whois server:", server)
	}
}

func TestParseDomainExtension(t *testing.T) {

	var domains = []struct {
		domain    string
		extension string
	}{
		{"domain.com.tr", ".com.tr"},
		{"domain.com", ".com"},
		{"domain.info", ".info"},
		{"domain.istanbul", ".istanbul"},
		{"domain.co.uk", ".co.uk"},
	}

	for _, v := range domains {

		ext := parseDomainNameExtension(v.domain)

		if ext != v.extension {
			t.Error("invalid extension capture:", ext, "=", v.extension)
		}

	}
}
