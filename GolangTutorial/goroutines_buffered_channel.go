/*
Order ascending by timestamp for the given JSON object and create a scheduler to read 10 records 
from the ordered data on each 10 seconds interval using 10 goroutines and 
print the first name and email id on the execution of each routine.
*/
package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

//Record struct
type Record struct {
	ID         int
	First_Name string
	Last_Name  string
	Email      string
	Gender     string
	IP_Address string
	Timestamp  string
}

//Records struct
type Records struct {
	Items []Record
}

type recSort []Record

func (r recSort) Len() int           { return len(r) }
func (r recSort) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r recSort) Less(i, j int) bool { return r[i].Timestamp < r[j].Timestamp }

var recs []Record
var ids int

func main() {
	recordBody := []byte(`[{"id":1,"first_name":"Salvatore","last_name":"Loos","email":"sloos0@surveymonkey.com","gender":"Male","ip_address":"171.233.241.55","timestamp":"1554052388"},
{"id":2,"first_name":"Claudianus","last_name":"Starcks","email":"cstarcks1@cyberchimps.com","gender":"Male","ip_address":"213.171.12.8","timestamp":"1538623871"},
{"id":3,"first_name":"Garvin","last_name":"Hofer","email":"ghofer2@cnn.com","gender":"Male","ip_address":"12.243.145.185","timestamp":"1534630782"},
{"id":4,"first_name":"Sharleen","last_name":"Vivash","email":"svivash3@parallels.com","gender":"Female","ip_address":"39.45.112.148","timestamp":"1548197591"},
{"id":5,"first_name":"Aristotle","last_name":"Grellier","email":"agrellier4@washington.edu","gender":"Male","ip_address":"232.78.140.16","timestamp":"1553744957"},
{"id":6,"first_name":"Barbara","last_name":"Van der Beek","email":"bvanderbeek5@wsj.com","gender":"Female","ip_address":"193.194.92.195","timestamp":"1526550575"},
{"id":7,"first_name":"Abel","last_name":"Strutley","email":"astrutley6@wp.com","gender":"Male","ip_address":"70.115.148.51","timestamp":"1541483164"},
{"id":8,"first_name":"Claudie","last_name":"Howcroft","email":"chowcroft7@studiopress.com","gender":"Female","ip_address":"171.50.229.175","timestamp":"1525079979"},
{"id":9,"first_name":"Chloette","last_name":"Halloway","email":"challoway8@pen.io","gender":"Female","ip_address":"201.170.69.209","timestamp":"1552849005"},
{"id":10,"first_name":"Abby","last_name":"Champerlen","email":"achamperlen9@bloglovin.com","gender":"Female","ip_address":"234.171.254.198","timestamp":"1539284234"},
{"id":11,"first_name":"Ninette","last_name":"Randlesome","email":"nrandlesomea@netvibes.com","gender":"Female","ip_address":"79.42.61.3","timestamp":"1549347124"},
{"id":12,"first_name":"Merilee","last_name":"Granger","email":"mgrangerb@netvibes.com","gender":"Female","ip_address":"153.165.53.241","timestamp":"1542189741"},
{"id":13,"first_name":"Chery","last_name":"Trussell","email":"ctrussellc@cyberchimps.com","gender":"Female","ip_address":"78.33.133.172","timestamp":"1530408288"},
{"id":14,"first_name":"Jermaine","last_name":"Capeling","email":"jcapelingd@google.nl","gender":"Male","ip_address":"187.66.72.11","timestamp":"1540544849"},
{"id":15,"first_name":"Lane","last_name":"Eager","email":"leagere@imdb.com","gender":"Male","ip_address":"243.11.168.218","timestamp":"1549600926"},
{"id":16,"first_name":"Cullan","last_name":"Steuhlmeyer","email":"csteuhlmeyerf@vimeo.com","gender":"Male","ip_address":"232.175.110.140","timestamp":"1552149696"},
{"id":17,"first_name":"Kirbee","last_name":"Sterland","email":"ksterlandg@ox.ac.uk","gender":"Female","ip_address":"98.221.77.90","timestamp":"1555195407"},
{"id":18,"first_name":"Othello","last_name":"Yushachkov","email":"oyushachkovh@google.co.uk","gender":"Male","ip_address":"105.106.66.72","timestamp":"1531393058"},
{"id":19,"first_name":"Angelle","last_name":"Faithfull","email":"afaithfulli@imageshack.us","gender":"Female","ip_address":"244.234.233.107","timestamp":"1545861267"},
{"id":20,"first_name":"Greg","last_name":"Ayers","email":"gayersj@go.com","gender":"Male","ip_address":"255.2.113.5","timestamp":"1549473970"},
{"id":21,"first_name":"Fay","last_name":"De Mars","email":"fdemarsk@about.me","gender":"Female","ip_address":"66.53.36.49","timestamp":"1542589943"},
{"id":22,"first_name":"Isidora","last_name":"Fedorski","email":"ifedorskil@theatlantic.com","gender":"Female","ip_address":"205.70.215.52","timestamp":"1546131629"},
{"id":23,"first_name":"Leontine","last_name":"Vreede","email":"lvreedem@google.co.jp","gender":"Female","ip_address":"241.124.127.178","timestamp":"1539300028"},
{"id":24,"first_name":"Lionello","last_name":"Sheers","email":"lsheersn@is.gd","gender":"Male","ip_address":"125.112.124.6","timestamp":"1543616830"},
{"id":25,"first_name":"Myrilla","last_name":"Vasyutochkin","email":"mvasyutochkino@hp.com","gender":"Female","ip_address":"87.236.86.57","timestamp":"1544083778"},
{"id":26,"first_name":"Redd","last_name":"Filippello","email":"rfilippellop@nationalgeographic.com","gender":"Male","ip_address":"249.92.238.20","timestamp":"1543854660"},
{"id":27,"first_name":"Gregor","last_name":"Neild","email":"gneildq@toplist.cz","gender":"Male","ip_address":"248.102.74.241","timestamp":"1533400053"},
{"id":28,"first_name":"Charles","last_name":"Reubel","email":"creubelr@washington.edu","gender":"Male","ip_address":"234.54.141.137","timestamp":"1553820850"},
{"id":29,"first_name":"Anton","last_name":"Roft","email":"arofts@omniture.com","gender":"Male","ip_address":"227.29.79.236","timestamp":"1534661878"},
{"id":30,"first_name":"Halsy","last_name":"Leyshon","email":"hleyshont@cbc.ca","gender":"Male","ip_address":"66.142.49.212","timestamp":"1544089117"},
{"id":31,"first_name":"Jacinthe","last_name":"Jellman","email":"jjellmanu@salon.com","gender":"Female","ip_address":"151.60.68.16","timestamp":"1549498396"},
{"id":32,"first_name":"Dyanne","last_name":"Petrillo","email":"dpetrillov@admin.ch","gender":"Female","ip_address":"62.219.45.204","timestamp":"1546768262"},
{"id":33,"first_name":"Florencia","last_name":"Spillman","email":"fspillmanw@dailymail.co.uk","gender":"Female","ip_address":"138.215.207.187","timestamp":"1543840961"},
{"id":34,"first_name":"Hyacinth","last_name":"Redding","email":"hreddingx@moonfruit.com","gender":"Female","ip_address":"91.180.219.35","timestamp":"1524692438"},
{"id":35,"first_name":"Domenico","last_name":"MacAree","email":"dmacareey@php.net","gender":"Male","ip_address":"132.196.116.117","timestamp":"1526541834"},
{"id":36,"first_name":"Alasteir","last_name":"Doubrava","email":"adoubravaz@nhs.uk","gender":"Male","ip_address":"236.123.136.1","timestamp":"1554138039"},
{"id":37,"first_name":"Glen","last_name":"Thompsett","email":"gthompsett10@walmart.com","gender":"Male","ip_address":"71.53.153.44","timestamp":"1532835277"},
{"id":38,"first_name":"Alexio","last_name":"Feacham","email":"afeacham11@youtu.be","gender":"Male","ip_address":"31.2.51.58","timestamp":"1540671225"},
{"id":39,"first_name":"Shayla","last_name":"Guidera","email":"sguidera12@dedecms.com","gender":"Female","ip_address":"186.218.139.73","timestamp":"1552534056"},
{"id":40,"first_name":"Joletta","last_name":"Dammarell","email":"jdammarell13@wsj.com","gender":"Female","ip_address":"85.237.117.92","timestamp":"1546677452"},
{"id":41,"first_name":"Haskell","last_name":"Gogan","email":"hgogan14@infoseek.co.jp","gender":"Male","ip_address":"109.109.67.144","timestamp":"1533953652"},
{"id":42,"first_name":"Tandy","last_name":"Matuszak","email":"tmatuszak15@virginia.edu","gender":"Female","ip_address":"32.160.224.153","timestamp":"1548103457"},
{"id":43,"first_name":"Errick","last_name":"Bushell","email":"ebushell16@admin.ch","gender":"Male","ip_address":"190.25.245.111","timestamp":"1532912493"},
{"id":44,"first_name":"Rodd","last_name":"Menendes","email":"rmenendes17@reference.com","gender":"Male","ip_address":"128.189.48.50","timestamp":"1548890961"},
{"id":45,"first_name":"Johnny","last_name":"Evered","email":"jevered18@msn.com","gender":"Male","ip_address":"253.175.252.52","timestamp":"1525174291"},
{"id":46,"first_name":"Edeline","last_name":"Brisco","email":"ebrisco19@over-blog.com","gender":"Female","ip_address":"170.96.167.238","timestamp":"1532364493"},
{"id":47,"first_name":"Arel","last_name":"Kurten","email":"akurten1a@mit.edu","gender":"Male","ip_address":"252.206.178.102","timestamp":"1548464920"},
{"id":48,"first_name":"Leisha","last_name":"Withrington","email":"lwithrington1b@mashable.com","gender":"Female","ip_address":"132.220.17.184","timestamp":"1528427616"},
{"id":49,"first_name":"Mitch","last_name":"Spadeck","email":"mspadeck1c@npr.org","gender":"Male","ip_address":"72.209.240.178","timestamp":"1548398922"},
{"id":50,"first_name":"Car","last_name":"Gindghill","email":"cgindghill1d@mayoclinic.com","gender":"Male","ip_address":"64.157.217.175","timestamp":"1525240787"},
{"id":51,"first_name":"Regan","last_name":"Maxwale","email":"rmaxwale1e@yellowpages.com","gender":"Female","ip_address":"48.25.54.108","timestamp":"1536942698"},
{"id":52,"first_name":"Hiram","last_name":"Luno","email":"hluno1f@list-manage.com","gender":"Male","ip_address":"120.153.7.232","timestamp":"1547103458"},
{"id":53,"first_name":"Cinderella","last_name":"Jiranek","email":"cjiranek1g@mashable.com","gender":"Female","ip_address":"34.112.52.116","timestamp":"1536628968"},
{"id":54,"first_name":"Adria","last_name":"Shewan","email":"ashewan1h@army.mil","gender":"Female","ip_address":"187.245.28.209","timestamp":"1550112462"},
{"id":55,"first_name":"Tobie","last_name":"Walas","email":"twalas1i@fastcompany.com","gender":"Male","ip_address":"244.139.82.235","timestamp":"1555239273"},
{"id":56,"first_name":"Cornell","last_name":"Colecrough","email":"ccolecrough1j@de.vu","gender":"Male","ip_address":"150.101.80.157","timestamp":"1528479989"},
{"id":57,"first_name":"Vittoria","last_name":"Oliva","email":"voliva1k@smugmug.com","gender":"Female","ip_address":"16.189.26.169","timestamp":"1546952921"},
{"id":58,"first_name":"Boot","last_name":"Eddow","email":"beddow1l@csmonitor.com","gender":"Male","ip_address":"153.226.247.8","timestamp":"1555095429"},
{"id":59,"first_name":"Breena","last_name":"Ceeley","email":"bceeley1m@surveymonkey.com","gender":"Female","ip_address":"109.252.61.90","timestamp":"1544579275"},
{"id":60,"first_name":"Kilian","last_name":"Chipps","email":"kchipps1n@fastcompany.com","gender":"Male","ip_address":"137.107.44.93","timestamp":"1544215740"},
{"id":61,"first_name":"Chevy","last_name":"Haking","email":"chaking1o@state.tx.us","gender":"Male","ip_address":"173.57.104.39","timestamp":"1541464297"},
{"id":62,"first_name":"Shela","last_name":"Pyner","email":"spyner1p@i2i.jp","gender":"Female","ip_address":"183.52.122.16","timestamp":"1544599193"},
{"id":63,"first_name":"Jethro","last_name":"Vela","email":"jvela1q@mashable.com","gender":"Male","ip_address":"36.162.16.154","timestamp":"1548515426"},
{"id":64,"first_name":"Camala","last_name":"Pryell","email":"cpryell1r@bizjournals.com","gender":"Female","ip_address":"23.83.235.48","timestamp":"1544160025"},
{"id":65,"first_name":"Hugues","last_name":"Wernher","email":"hwernher1s@unc.edu","gender":"Male","ip_address":"242.170.26.171","timestamp":"1546010462"},
{"id":66,"first_name":"Shamus","last_name":"Spikeings","email":"sspikeings1t@sina.com.cn","gender":"Male","ip_address":"169.82.225.144","timestamp":"1549188612"},
{"id":67,"first_name":"Letizia","last_name":"Hymas","email":"lhymas1u@chron.com","gender":"Female","ip_address":"10.209.10.224","timestamp":"1549895002"},
{"id":68,"first_name":"Cyndy","last_name":"Wardel","email":"cwardel1v@wordpress.org","gender":"Female","ip_address":"178.24.22.91","timestamp":"1537466340"},
{"id":69,"first_name":"Daniella","last_name":"Apps","email":"dapps1w@yellowpages.com","gender":"Female","ip_address":"107.98.23.115","timestamp":"1545041320"},
{"id":70,"first_name":"Sigrid","last_name":"Waterson","email":"swaterson1x@ucla.edu","gender":"Female","ip_address":"166.203.106.167","timestamp":"1533403168"},
{"id":71,"first_name":"Siffre","last_name":"Belden","email":"sbelden1y@yahoo.com","gender":"Male","ip_address":"153.65.1.40","timestamp":"1528966316"},
{"id":72,"first_name":"Garret","last_name":"McEnhill","email":"gmcenhill1z@businessinsider.com","gender":"Male","ip_address":"118.7.63.177","timestamp":"1545103120"},
{"id":73,"first_name":"Ronda","last_name":"Lissandrini","email":"rlissandrini20@elpais.com","gender":"Female","ip_address":"233.53.185.111","timestamp":"1540917030"},
{"id":74,"first_name":"Damiano","last_name":"Linguard","email":"dlinguard21@theatlantic.com","gender":"Male","ip_address":"155.48.67.113","timestamp":"1523873043"},
{"id":75,"first_name":"Christophe","last_name":"Ayris","email":"cayris22@cbslocal.com","gender":"Male","ip_address":"107.36.23.75","timestamp":"1535067606"},
{"id":76,"first_name":"Antonin","last_name":"O'Scanlon","email":"aoscanlon23@rambler.ru","gender":"Male","ip_address":"1.244.132.206","timestamp":"1543212444"},
{"id":77,"first_name":"Donica","last_name":"Crocker","email":"dcrocker24@vistaprint.com","gender":"Female","ip_address":"215.36.220.26","timestamp":"1540520929"},
{"id":78,"first_name":"Faith","last_name":"Chorley","email":"fchorley25@aol.com","gender":"Female","ip_address":"30.220.192.53","timestamp":"1550170859"},
{"id":79,"first_name":"Britt","last_name":"Dudmarsh","email":"bdudmarsh26@rambler.ru","gender":"Male","ip_address":"112.168.113.228","timestamp":"1527560945"},
{"id":80,"first_name":"Cirillo","last_name":"Reagan","email":"creagan27@sfgate.com","gender":"Male","ip_address":"180.112.15.136","timestamp":"1530944693"},
{"id":81,"first_name":"Daniella","last_name":"Breslauer","email":"dbreslauer28@stanford.edu","gender":"Female","ip_address":"119.241.231.140","timestamp":"1524691145"},
{"id":82,"first_name":"Engelbert","last_name":"Eymor","email":"eeymor29@google.com.br","gender":"Male","ip_address":"116.186.171.164","timestamp":"1526472882"},
{"id":83,"first_name":"Maximilien","last_name":"Byre","email":"mbyre2a@usda.gov","gender":"Male","ip_address":"60.227.105.204","timestamp":"1544657487"},
{"id":84,"first_name":"Debbie","last_name":"Gaythor","email":"dgaythor2b@twitter.com","gender":"Female","ip_address":"22.67.127.18","timestamp":"1534613240"},
{"id":85,"first_name":"Lucie","last_name":"Murrie","email":"lmurrie2c@drupal.org","gender":"Female","ip_address":"232.81.120.142","timestamp":"1523913485"},
{"id":86,"first_name":"William","last_name":"Dongles","email":"wdongles2d@noaa.gov","gender":"Male","ip_address":"219.163.138.168","timestamp":"1526684786"},
{"id":87,"first_name":"Ignatius","last_name":"Trembley","email":"itrembley2e@newsvine.com","gender":"Male","ip_address":"250.33.217.180","timestamp":"1554927221"},
{"id":88,"first_name":"Jesus","last_name":"McGuckin","email":"jmcguckin2f@list-manage.com","gender":"Male","ip_address":"124.7.108.210","timestamp":"1553043042"},
{"id":89,"first_name":"Morna","last_name":"Sich","email":"msich2g@ucla.edu","gender":"Female","ip_address":"86.242.54.114","timestamp":"1550064869"},
{"id":90,"first_name":"Vasilis","last_name":"Dragoe","email":"vdragoe2h@dot.gov","gender":"Male","ip_address":"11.169.37.186","timestamp":"1535964700"},
{"id":91,"first_name":"Trace","last_name":"Newling","email":"tnewling2i@xing.com","gender":"Male","ip_address":"242.31.71.26","timestamp":"1524276585"},
{"id":92,"first_name":"Erin","last_name":"Eustes","email":"eeustes2j@flavors.me","gender":"Male","ip_address":"244.216.217.226","timestamp":"1526510573"},
{"id":93,"first_name":"Hanny","last_name":"Cruddace","email":"hcruddace2k@over-blog.com","gender":"Female","ip_address":"211.235.68.238","timestamp":"1535370392"},
{"id":94,"first_name":"Randall","last_name":"Lenox","email":"rlenox2l@arstechnica.com","gender":"Male","ip_address":"227.100.249.186","timestamp":"1530804435"},
{"id":95,"first_name":"Jami","last_name":"Rennles","email":"jrennles2m@toplist.cz","gender":"Female","ip_address":"2.105.126.132","timestamp":"1544880455"},
{"id":96,"first_name":"Sol","last_name":"Mapston","email":"smapston2n@tinypic.com","gender":"Male","ip_address":"207.172.133.161","timestamp":"1526626926"},
{"id":97,"first_name":"Erhart","last_name":"Lambin","email":"elambin2o@furl.net","gender":"Male","ip_address":"5.164.204.208","timestamp":"1547865546"},
{"id":98,"first_name":"Tye","last_name":"Chichgar","email":"tchichgar2p@ebay.com","gender":"Male","ip_address":"231.137.17.58","timestamp":"1550145940"},
{"id":99,"first_name":"Bradan","last_name":"Dougherty","email":"bdougherty2q@51.la","gender":"Male","ip_address":"232.129.1.91","timestamp":"1552531509"},
{"id":100,"first_name":"Angelita","last_name":"Fussey","email":"afussey2r@dion.ne.jp","gender":"Female","ip_address":"131.155.20.213","timestamp":"1535249637"},
{"id":101,"first_name":"Deborah","last_name":"Koubu","email":"dkoubu2s@acquirethisname.com","gender":"Female","ip_address":"158.5.35.127","timestamp":"1542337528"},
{"id":102,"first_name":"Sidnee","last_name":"Greengrass","email":"sgreengrass2t@amazon.co.jp","gender":"Male","ip_address":"21.20.78.223","timestamp":"1549216769"},
{"id":103,"first_name":"Moina","last_name":"Tankin","email":"mtankin2u@uol.com.br","gender":"Female","ip_address":"49.13.87.121","timestamp":"1537260054"},
{"id":104,"first_name":"Helen-elizabeth","last_name":"Pawlata","email":"hpawlata2v@php.net","gender":"Female","ip_address":"1.109.225.84","timestamp":"1525391688"},
{"id":105,"first_name":"Carissa","last_name":"Pembery","email":"cpembery2w@sciencedirect.com","gender":"Female","ip_address":"250.118.41.187","timestamp":"1543407554"},
{"id":106,"first_name":"Natal","last_name":"Fibbitts","email":"nfibbitts2x@cdbaby.com","gender":"Male","ip_address":"227.6.248.233","timestamp":"1552742987"},
{"id":107,"first_name":"Tadeas","last_name":"Flory","email":"tflory2y@npr.org","gender":"Male","ip_address":"68.197.198.1","timestamp":"1537015462"},
{"id":108,"first_name":"Nils","last_name":"Gamblin","email":"ngamblin2z@si.edu","gender":"Male","ip_address":"11.223.45.206","timestamp":"1548653294"},
{"id":109,"first_name":"Siffre","last_name":"Brookesbie","email":"sbrookesbie30@timesonline.co.uk","gender":"Male","ip_address":"57.228.45.205","timestamp":"1544367181"},
{"id":110,"first_name":"Shepard","last_name":"Crilly","email":"scrilly31@bigcartel.com","gender":"Male","ip_address":"105.206.108.225","timestamp":"1549079955"},
{"id":111,"first_name":"Odessa","last_name":"Parcall","email":"oparcall32@canalblog.com","gender":"Female","ip_address":"29.128.238.47","timestamp":"1539001161"},
{"id":112,"first_name":"Paulo","last_name":"Cammiemile","email":"pcammiemile33@ask.com","gender":"Male","ip_address":"243.77.57.31","timestamp":"1543110562"},
{"id":113,"first_name":"Vanessa","last_name":"Beller","email":"vbeller34@bloglovin.com","gender":"Female","ip_address":"2.79.13.126","timestamp":"1550342989"},
{"id":114,"first_name":"Paola","last_name":"Wartnaby","email":"pwartnaby35@illinois.edu","gender":"Female","ip_address":"216.24.203.96","timestamp":"1531005894"},
{"id":115,"first_name":"Alane","last_name":"Cornfield","email":"acornfield36@weebly.com","gender":"Female","ip_address":"21.172.221.226","timestamp":"1525243314"},
{"id":116,"first_name":"Berna","last_name":"Sakins","email":"bsakins37@scribd.com","gender":"Female","ip_address":"152.146.240.115","timestamp":"1545551063"},
{"id":117,"first_name":"Noami","last_name":"Sirey","email":"nsirey38@ebay.co.uk","gender":"Female","ip_address":"220.191.52.237","timestamp":"1528520875"},
{"id":118,"first_name":"Britt","last_name":"Dahlback","email":"bdahlback39@wsj.com","gender":"Female","ip_address":"131.8.236.8","timestamp":"1544242862"},
{"id":119,"first_name":"Darell","last_name":"Foli","email":"dfoli3a@g.co","gender":"Female","ip_address":"210.216.28.129","timestamp":"1528533636"},
{"id":120,"first_name":"Hakeem","last_name":"Baughan","email":"hbaughan3b@bloglovin.com","gender":"Male","ip_address":"168.172.152.181","timestamp":"1549333632"},
{"id":121,"first_name":"Estel","last_name":"Geekin","email":"egeekin3c@yolasite.com","gender":"Female","ip_address":"196.120.85.77","timestamp":"1544590616"},
{"id":122,"first_name":"Fair","last_name":"Wythill","email":"fwythill3d@elegantthemes.com","gender":"Male","ip_address":"101.157.159.59","timestamp":"1551351992"},
{"id":123,"first_name":"Hallsy","last_name":"Spelwood","email":"hspelwood3e@dailymail.co.uk","gender":"Male","ip_address":"85.93.79.147","timestamp":"1554456045"},
{"id":124,"first_name":"Antoni","last_name":"Lambirth","email":"alambirth3f@1688.com","gender":"Male","ip_address":"0.67.70.187","timestamp":"1539778884"},
{"id":125,"first_name":"Nathaniel","last_name":"Brettoner","email":"nbrettoner3g@myspace.com","gender":"Male","ip_address":"162.33.176.195","timestamp":"1537325383"},
{"id":126,"first_name":"Francois","last_name":"Durnill","email":"fdurnill3h@cnbc.com","gender":"Male","ip_address":"185.248.250.30","timestamp":"1527102972"},
{"id":127,"first_name":"Neal","last_name":"Macklam","email":"nmacklam3i@sun.com","gender":"Male","ip_address":"187.63.252.139","timestamp":"1554519199"},
{"id":128,"first_name":"Rab","last_name":"De la Barre","email":"rdelabarre3j@chicagotribune.com","gender":"Male","ip_address":"207.225.178.187","timestamp":"1534024188"},
{"id":129,"first_name":"Thoma","last_name":"Terney","email":"tterney3k@youtube.com","gender":"Male","ip_address":"108.77.163.190","timestamp":"1547861050"},
{"id":130,"first_name":"Nanette","last_name":"Ashton","email":"nashton3l@sourceforge.net","gender":"Female","ip_address":"40.124.35.168","timestamp":"1554558898"},
{"id":131,"first_name":"Dalila","last_name":"Colquitt","email":"dcolquitt3m@purevolume.com","gender":"Female","ip_address":"25.34.252.83","timestamp":"1542207783"},
{"id":132,"first_name":"Philbert","last_name":"McIlvenna","email":"pmcilvenna3n@upenn.edu","gender":"Male","ip_address":"49.21.219.22","timestamp":"1548401786"},
{"id":133,"first_name":"Sophronia","last_name":"Grint","email":"sgrint3o@about.com","gender":"Female","ip_address":"168.87.162.131","timestamp":"1551354882"},
{"id":134,"first_name":"Kerianne","last_name":"Backes","email":"kbackes3p@nps.gov","gender":"Female","ip_address":"0.66.31.97","timestamp":"1554293115"},
{"id":135,"first_name":"Wake","last_name":"Littrell","email":"wlittrell3q@intel.com","gender":"Male","ip_address":"113.119.98.129","timestamp":"1528468673"},
{"id":136,"first_name":"Demott","last_name":"Kemp","email":"dkemp3r@ustream.tv","gender":"Male","ip_address":"252.238.224.6","timestamp":"1527578367"},
{"id":137,"first_name":"Conan","last_name":"Donn","email":"cdonn3s@ftc.gov","gender":"Male","ip_address":"255.213.130.56","timestamp":"1547997030"},
{"id":138,"first_name":"Lou","last_name":"Pounds","email":"lpounds3t@epa.gov","gender":"Male","ip_address":"5.115.159.126","timestamp":"1527873599"},
{"id":139,"first_name":"Cord","last_name":"Kondratovich","email":"ckondratovich3u@sfgate.com","gender":"Male","ip_address":"37.133.212.219","timestamp":"1554051574"},
{"id":140,"first_name":"Corine","last_name":"Ballingal","email":"cballingal3v@omniture.com","gender":"Female","ip_address":"219.22.68.35","timestamp":"1548443423"},
{"id":141,"first_name":"Nicky","last_name":"Venour","email":"nvenour3w@nps.gov","gender":"Male","ip_address":"56.0.100.71","timestamp":"1532412604"},
{"id":142,"first_name":"Kerrin","last_name":"Ziems","email":"kziems3x@shareasale.com","gender":"Female","ip_address":"110.47.90.32","timestamp":"1544565354"},
{"id":143,"first_name":"Penn","last_name":"Calleja","email":"pcalleja3y@state.tx.us","gender":"Male","ip_address":"60.240.150.184","timestamp":"1527018356"},
{"id":144,"first_name":"Mack","last_name":"Harbisher","email":"mharbisher3z@amazon.co.uk","gender":"Male","ip_address":"17.84.69.137","timestamp":"1535717379"},
{"id":145,"first_name":"Minda","last_name":"Gyford","email":"mgyford40@list-manage.com","gender":"Female","ip_address":"160.251.239.28","timestamp":"1541083052"},
{"id":146,"first_name":"Chrissie","last_name":"Sorley","email":"csorley41@stumbleupon.com","gender":"Male","ip_address":"46.246.199.36","timestamp":"1544140603"},
{"id":147,"first_name":"Gaultiero","last_name":"Daft","email":"gdaft42@bing.com","gender":"Male","ip_address":"156.156.164.163","timestamp":"1542224950"},
{"id":148,"first_name":"Moishe","last_name":"Solano","email":"msolano43@stumbleupon.com","gender":"Male","ip_address":"203.64.115.8","timestamp":"1543030015"},
{"id":149,"first_name":"Neddy","last_name":"Malbon","email":"nmalbon44@geocities.com","gender":"Male","ip_address":"238.228.124.153","timestamp":"1529802556"},
{"id":150,"first_name":"Morey","last_name":"Andriulis","email":"mandriulis45@phoca.cz","gender":"Male","ip_address":"110.242.65.39","timestamp":"1544575256"},
{"id":151,"first_name":"Terence","last_name":"Howen","email":"thowen46@hhs.gov","gender":"Male","ip_address":"233.220.69.62","timestamp":"1544018520"},
{"id":152,"first_name":"Debee","last_name":"Caesar","email":"dcaesar47@un.org","gender":"Female","ip_address":"155.156.247.95","timestamp":"1542526709"},
{"id":153,"first_name":"Spenser","last_name":"Hillett","email":"shillett48@pinterest.com","gender":"Male","ip_address":"124.220.103.200","timestamp":"1535130748"},
{"id":154,"first_name":"Netty","last_name":"Griffin","email":"ngriffin49@plala.or.jp","gender":"Female","ip_address":"148.196.191.175","timestamp":"1526704755"},
{"id":155,"first_name":"Rosemonde","last_name":"Sigward","email":"rsigward4a@answers.com","gender":"Female","ip_address":"89.169.228.63","timestamp":"1538137853"},
{"id":156,"first_name":"Ann-marie","last_name":"Beldan","email":"abeldan4b@elegantthemes.com","gender":"Female","ip_address":"63.237.56.51","timestamp":"1551301730"},
{"id":157,"first_name":"Mahalia","last_name":"Fosdike","email":"mfosdike4c@comcast.net","gender":"Female","ip_address":"129.16.21.236","timestamp":"1549503172"},
{"id":158,"first_name":"Ardra","last_name":"Quince","email":"aquince4d@multiply.com","gender":"Female","ip_address":"187.22.111.169","timestamp":"1526608507"},
{"id":159,"first_name":"Wadsworth","last_name":"De-Ville","email":"wdeville4e@dailymail.co.uk","gender":"Male","ip_address":"160.146.160.182","timestamp":"1531395171"},
{"id":160,"first_name":"Reuven","last_name":"Gisburne","email":"rgisburne4f@newsvine.com","gender":"Male","ip_address":"243.117.135.138","timestamp":"1523757196"},
{"id":161,"first_name":"Ranna","last_name":"MacKay","email":"rmackay4g@comcast.net","gender":"Female","ip_address":"209.89.101.221","timestamp":"1540417984"},
{"id":162,"first_name":"Adria","last_name":"Jose","email":"ajose4h@yahoo.co.jp","gender":"Female","ip_address":"168.124.194.55","timestamp":"1539968979"},
{"id":163,"first_name":"Beatrisa","last_name":"Delany","email":"bdelany4i@examiner.com","gender":"Female","ip_address":"90.77.141.22","timestamp":"1527605763"},
{"id":164,"first_name":"Frannie","last_name":"Ditty","email":"fditty4j@wsj.com","gender":"Female","ip_address":"231.33.50.125","timestamp":"1527981429"},
{"id":165,"first_name":"Harriet","last_name":"Bliss","email":"hbliss4k@vk.com","gender":"Female","ip_address":"231.127.199.198","timestamp":"1543025543"},
{"id":166,"first_name":"Shayne","last_name":"Cranna","email":"scranna4l@economist.com","gender":"Male","ip_address":"103.178.157.253","timestamp":"1555208637"},
{"id":167,"first_name":"Brena","last_name":"Pellitt","email":"bpellitt4m@umn.edu","gender":"Female","ip_address":"10.222.121.168","timestamp":"1545544527"},
{"id":168,"first_name":"Elinor","last_name":"Giveen","email":"egiveen4n@usgs.gov","gender":"Female","ip_address":"90.226.95.100","timestamp":"1529000449"},
{"id":169,"first_name":"Sharron","last_name":"Creggan","email":"screggan4o@google.pl","gender":"Female","ip_address":"238.64.231.108","timestamp":"1535735665"},
{"id":170,"first_name":"Andros","last_name":"Beltzner","email":"abeltzner4p@github.com","gender":"Male","ip_address":"143.243.5.207","timestamp":"1535454836"},
{"id":171,"first_name":"Kelsi","last_name":"Rentelll","email":"krentelll4q@blog.com","gender":"Female","ip_address":"170.129.132.211","timestamp":"1543467840"},
{"id":172,"first_name":"Clayson","last_name":"Atkirk","email":"catkirk4r@purevolume.com","gender":"Male","ip_address":"200.85.23.228","timestamp":"1532622084"},
{"id":173,"first_name":"Kalle","last_name":"Reyes","email":"kreyes4s@berkeley.edu","gender":"Male","ip_address":"254.182.255.147","timestamp":"1524716487"},
{"id":174,"first_name":"Darbee","last_name":"Bixley","email":"dbixley4t@foxnews.com","gender":"Male","ip_address":"83.26.218.237","timestamp":"1529884371"},
{"id":175,"first_name":"Caterina","last_name":"Aymer","email":"caymer4u@nyu.edu","gender":"Female","ip_address":"34.46.50.162","timestamp":"1547363418"},
{"id":176,"first_name":"Marta","last_name":"Geely","email":"mgeely4v@dion.ne.jp","gender":"Female","ip_address":"170.73.76.5","timestamp":"1551071786"},
{"id":177,"first_name":"Keelia","last_name":"Joselovitch","email":"kjoselovitch4w@boston.com","gender":"Female","ip_address":"20.97.229.155","timestamp":"1534103140"},
{"id":178,"first_name":"Michal","last_name":"Simeonov","email":"msimeonov4x@tripadvisor.com","gender":"Female","ip_address":"42.244.215.118","timestamp":"1537662478"},
{"id":179,"first_name":"Barthel","last_name":"Sandeland","email":"bsandeland4y@meetup.com","gender":"Male","ip_address":"161.197.247.87","timestamp":"1535258086"},
{"id":180,"first_name":"Page","last_name":"Marris","email":"pmarris4z@spiegel.de","gender":"Female","ip_address":"179.164.20.174","timestamp":"1533086805"},
{"id":181,"first_name":"Cletus","last_name":"Roizin","email":"croizin50@1und1.de","gender":"Male","ip_address":"142.83.198.249","timestamp":"1545411388"},
{"id":182,"first_name":"Garrick","last_name":"Girardey","email":"ggirardey51@joomla.org","gender":"Male","ip_address":"51.54.224.50","timestamp":"1540603175"},
{"id":183,"first_name":"Jarvis","last_name":"Bradshaw","email":"jbradshaw52@ameblo.jp","gender":"Male","ip_address":"227.57.197.39","timestamp":"1526531141"},
{"id":184,"first_name":"Reed","last_name":"Larrad","email":"rlarrad53@google.pl","gender":"Male","ip_address":"165.227.255.42","timestamp":"1525655222"},
{"id":185,"first_name":"Dilan","last_name":"Bantham","email":"dbantham54@chicagotribune.com","gender":"Male","ip_address":"162.86.89.170","timestamp":"1538439853"},
{"id":186,"first_name":"Ivy","last_name":"Majury","email":"imajury55@storify.com","gender":"Female","ip_address":"242.217.5.10","timestamp":"1547057961"},
{"id":187,"first_name":"Marga","last_name":"Rennenbach","email":"mrennenbach56@infoseek.co.jp","gender":"Female","ip_address":"140.116.110.79","timestamp":"1535518319"},
{"id":188,"first_name":"Rutherford","last_name":"Logan","email":"rlogan57@adobe.com","gender":"Male","ip_address":"195.121.158.169","timestamp":"1534090345"},
{"id":189,"first_name":"Chen","last_name":"Durdle","email":"cdurdle58@rakuten.co.jp","gender":"Male","ip_address":"88.54.207.135","timestamp":"1542773706"},
{"id":190,"first_name":"Ginevra","last_name":"Balle","email":"gballe59@symantec.com","gender":"Female","ip_address":"18.96.201.166","timestamp":"1532935243"},
{"id":191,"first_name":"Wilt","last_name":"Ruggieri","email":"wruggieri5a@sohu.com","gender":"Male","ip_address":"104.127.9.36","timestamp":"1537876780"},
{"id":192,"first_name":"Gabe","last_name":"Tydd","email":"gtydd5b@google.com.br","gender":"Male","ip_address":"35.207.81.231","timestamp":"1530408924"},
{"id":193,"first_name":"Tamra","last_name":"Ealden","email":"tealden5c@constantcontact.com","gender":"Female","ip_address":"32.19.39.40","timestamp":"1546810700"},
{"id":194,"first_name":"Greer","last_name":"Skittrall","email":"gskittrall5d@blogtalkradio.com","gender":"Female","ip_address":"207.144.255.198","timestamp":"1543434924"},
{"id":195,"first_name":"Fee","last_name":"Fonte","email":"ffonte5e@quantcast.com","gender":"Male","ip_address":"141.102.154.243","timestamp":"1536484001"},
{"id":196,"first_name":"Alyda","last_name":"Milbourne","email":"amilbourne5f@desdev.cn","gender":"Female","ip_address":"25.174.185.40","timestamp":"1542416556"},
{"id":197,"first_name":"Laurel","last_name":"Garlant","email":"lgarlant5g@npr.org","gender":"Female","ip_address":"138.72.242.239","timestamp":"1542676533"},
{"id":198,"first_name":"Marcelline","last_name":"Stait","email":"mstait5h@yelp.com","gender":"Female","ip_address":"17.76.152.188","timestamp":"1539976428"},
{"id":199,"first_name":"Hasheem","last_name":"Ghost","email":"hghost5i@blog.com","gender":"Male","ip_address":"169.101.214.119","timestamp":"1550259042"},
{"id":200,"first_name":"Pryce","last_name":"Yoslowitz","email":"pyoslowitz5j@chicagotribune.com","gender":"Male","ip_address":"141.192.83.226","timestamp":"1553137621"}]`)

	recs = make([]Record, 0)
	json.Unmarshal(recordBody, &recs)

	sort.Sort(recSort(recs))

	go mainSchedule()
	time.Sleep(time.Second * time.Duration(len(recs)))
}

func mainSchedule() {
	//Every 10 secs End 10 records will be printed
	for range time.Tick(time.Second * 10) {
		ch := make(chan int)
		go Schedule(ch, 9)
		for num := range ch {
			fmt.Println(num, "\t First Name:", recs[num].First_Name, "\t\t Email:", recs[num].Email)
		}
	}
}

//Schedule will do Buffered channel
func Schedule(ch chan int, count int) {

	for count >= 0 {
		ch <- ids
		count--
		ids++
	}
	close(ch)
}
