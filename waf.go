package cloudflare

import (
	"encoding/json"
)

func (api *API) ListWAFPackages(zid string) ([]WAFPackage, error) {
	var p WAFPackagesResponse
	var packages []WAFPackage
	var res []byte
	var err error
	uri := "/zones/" + zid + "/firewall/waf/packages"
	res, err = api.makeRequest("GET", uri, nil)
	if err != nil {
		return []WAFPackage{}, err
	}
	err = json.Unmarshal(res, &p)
	if err != nil {
		return []WAFPackage{}, err
	}
	if !p.Success {
		return []WAFPackage{}, err
	}
	for pi, _ := range p.Result {
		packages = append(packages, p.Result[pi])
	}
	return packages, err
}

func (api *API) ListWAFRules(zid, pid string) ([]WAFRule, error) {
	var r WAFRulesResponse
	var rules []WAFRule
	var res []byte
	var err error
	uri := "/zones/" + zid + "/firewall/waf/packages/" + pid + "/rules"
	res, err = api.makeRequest("GET", uri, nil)
	if err != nil {
		return []WAFRule{}, err
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []WAFRule{}, err
	}
	if !r.Success {
		return []WAFRule{}, err
	}
	for ri, _ := range r.Result {
		rules = append(rules, r.Result[ri])
	}
	return rules, err
}
