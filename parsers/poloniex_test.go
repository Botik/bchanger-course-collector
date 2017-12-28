package parsers

import "testing"

var poloniexData = []byte(`{"BTC_BCN":{"id":7,"last":"0.00000044","lowestAsk":"0.00000043","highestBid":"0.00000042","percentChange":"0.62962962","baseVolume":"543.23019195","quoteVolume":"1535903340.68901491","isFrozen":"0","high24hr":"0.00000048","low24hr":"0.00000026"},"BTC_BELA":{"id":8,"last":"0.00001899","lowestAsk":"0.00001899","highestBid":"0.00001898","percentChange":"0.38310269","baseVolume":"78.70502295","quoteVolume":"4513906.79213982","isFrozen":"0","high24hr":"0.00002140","low24hr":"0.00001320"},"BTC_BLK":{"id":10,"last":"0.00003754","lowestAsk":"0.00003755","highestBid":"0.00003754","percentChange":"0.14277016","baseVolume":"40.02061766","quoteVolume":"1107777.66914239","isFrozen":"0","high24hr":"0.00003900","low24hr":"0.00003286"},"BTC_BTCD":{"id":12,"last":"0.01700000","lowestAsk":"0.01700000","highestBid":"0.01698898","percentChange":"0.38077009","baseVolume":"43.54501098","quoteVolume":"2849.79092662","isFrozen":"0","high24hr":"0.01765238","low24hr":"0.01170005"},"BTC_BTM":{"id":13,"last":"0.00014102","lowestAsk":"0.00014133","highestBid":"0.00014102","percentChange":"0.11267161","baseVolume":"21.81128800","quoteVolume":"164220.86890843","isFrozen":"0","high24hr":"0.00014300","low24hr":"0.00012018"},"BTC_BTS":{"id":14,"last":"0.00003717","lowestAsk":"0.00003717","highestBid":"0.00003716","percentChange":"0.25151515","baseVolume":"1874.69750525","quoteVolume":"52543621.79303084","isFrozen":"0","high24hr":"0.00004000","low24hr":"0.00002903"},"BTC_BURST":{"id":15,"last":"0.00000275","lowestAsk":"0.00000275","highestBid":"0.00000273","percentChange":"0.77419354","baseVolume":"277.64208826","quoteVolume":"134455860.65617618","isFrozen":"0","high24hr":"0.00000279","low24hr":"0.00000150"},"BTC_CLAM":{"id":20,"last":"0.00071809","lowestAsk":"0.00071800","highestBid":"0.00070914","percentChange":"0.15820967","baseVolume":"90.36668383","quoteVolume":"136225.88679721","isFrozen":"0","high24hr":"0.00072000","low24hr":"0.00062000"},"BTC_DASH":{"id":24,"last":"0.09026988","lowestAsk":"0.09026988","highestBid":"0.09026987","percentChange":"0.08371717","baseVolume":"3438.53337556","quoteVolume":"39106.86642720","isFrozen":"0","high24hr":"0.09199600","low24hr":"0.08298588"},"BTC_DGB":{"id":25,"last":"0.00000327","lowestAsk":"0.00000327","highestBid":"0.00000326","percentChange":"0.18478260","baseVolume":"2260.14328858","quoteVolume":"718803886.00195932","isFrozen":"0","high24hr":"0.00000364","low24hr":"0.00000269"},"BTC_DOGE":{"id":27,"last":"0.00000051","lowestAsk":"0.00000051","highestBid":"0.00000050","percentChange":"0.54545454","baseVolume":"2724.52805164","quoteVolume":"6388170244.57754230","isFrozen":"0","high24hr":"0.00000052","low24hr":"0.00000032"},"BTC_EMC2":{"id":28,"last":"0.00008410","lowestAsk":"0.00008412","highestBid":"0.00008410","percentChange":"-0.06514006","baseVolume":"989.75797831","quoteVolume":"11328845.26623350","isFrozen":"0","high24hr":"0.00009568","low24hr":"0.00008000"},"BTC_FLDC":{"id":31,"last":"0.00000395","lowestAsk":"0.00000394","highestBid":"0.00000391","percentChange":"0.89903846","baseVolume":"327.69703334","quoteVolume":"96827392.08357923","isFrozen":"0","high24hr":"0.00000488","low24hr":"0.00000203"},"BTC_FLO":{"id":32,"last":"0.00001720","lowestAsk":"0.00001721","highestBid":"0.00001720","percentChange":"0.22857142","baseVolume":"64.94667791","quoteVolume":"3866684.78892577","isFrozen":"0","high24hr":"0.00001921","low24hr":"0.00001383"},"BTC_GAME":{"id":38,"last":"0.00022406","lowestAsk":"0.00022550","highestBid":"0.00022406","percentChange":"0.16274001","baseVolume":"73.56160317","quoteVolume":"353717.08825754","isFrozen":"0","high24hr":"0.00022550","low24hr":"0.00018688"},"BTC_GRC":{"id":40,"last":"0.00000675","lowestAsk":"0.00000675","highestBid":"0.00000671","percentChange":"0.17187500","baseVolume":"48.98940955","quoteVolume":"7822112.28306983","isFrozen":"0","high24hr":"0.00000720","low24hr":"0.00000571"},"BTC_HUC":{"id":43,"last":"0.00001737","lowestAsk":"0.00001738","highestBid":"0.00001737","percentChange":"0.12209302","baseVolume":"7.75578073","quoteVolume":"472704.12730885","isFrozen":"0","high24hr":"0.00001738","low24hr":"0.00001512"},"BTC_LTC":{"id":50,"last":"0.01894800","lowestAsk":"0.01895000","highestBid":"0.01894800","percentChange":"-0.00795811","baseVolume":"4701.23889985","quoteVolume":"250680.12819073","isFrozen":"0","high24hr":"0.01917876","low24hr":"0.01831878"},"BTC_MAID":{"id":51,"last":"0.00005213","lowestAsk":"0.00005240","highestBid":"0.00005204","percentChange":"0.14445664","baseVolume":"313.71980816","quoteVolume":"6517757.62370088","isFrozen":"0","high24hr":"0.00005353","low24hr":"0.00004384"},"BTC_OMNI":{"id":58,"last":"0.00561112","lowestAsk":"0.00561113","highestBid":"0.00561112","percentChange":"0.24332373","baseVolume":"67.17476833","quoteVolume":"14171.66862597","isFrozen":"0","high24hr":"0.00568199","low24hr":"0.00419002"},"BTC_NAV":{"id":61,"last":"0.00017848","lowestAsk":"0.00017848","highestBid":"0.00017833","percentChange":"0.09035371","baseVolume":"230.81435132","quoteVolume":"1298071.39954394","isFrozen":"0","high24hr":"0.00019988","low24hr":"0.00016173"},"BTC_NEOS":{"id":63,"last":"0.00042000","lowestAsk":"0.00041972","highestBid":"0.00041950","percentChange":"0.11326105","baseVolume":"25.55784185","quoteVolume":"64249.72800564","isFrozen":"0","high24hr":"0.00042710","low24hr":"0.00037318"},"BTC_NMC":{"id":64,"last":"0.00027500","lowestAsk":"0.00027553","highestBid":"0.00027500","percentChange":"0.11214461","baseVolume":"16.66962441","quoteVolume":"63838.76188566","isFrozen":"0","high24hr":"0.00027820","low24hr":"0.00023999"},"BTC_NXT":{"id":69,"last":"0.00007804","lowestAsk":"0.00007804","highestBid":"0.00007801","percentChange":"0.46609055","baseVolume":"2638.68506675","quoteVolume":"38399770.08350569","isFrozen":"0","high24hr":"0.00008300","low24hr":"0.00005218"},"BTC_PINK":{"id":73,"last":"0.00000362","lowestAsk":"0.00000364","highestBid":"0.00000362","percentChange":"0.52100840","baseVolume":"135.25688662","quoteVolume":"39898436.58983550","isFrozen":"0","high24hr":"0.00000418","low24hr":"0.00000231"},"BTC_POT":{"id":74,"last":"0.00002147","lowestAsk":"0.00002146","highestBid":"0.00002145","percentChange":"0.11997913","baseVolume":"119.57073232","quoteVolume":"5806940.05933000","isFrozen":"0","high24hr":"0.00002260","low24hr":"0.00001894"},"BTC_PPC":{"id":75,"last":"0.00037171","lowestAsk":"0.00037165","highestBid":"0.00036601","percentChange":"0.08080367","baseVolume":"35.75727949","quoteVolume":"103016.98293173","isFrozen":"0","high24hr":"0.00037220","low24hr":"0.00033001"},"BTC_RIC":{"id":83,"last":"0.00001636","lowestAsk":"0.00001640","highestBid":"0.00001636","percentChange":"0.15049226","baseVolume":"38.46634658","quoteVolume":"2517387.17505207","isFrozen":"0","high24hr":"0.00001697","low24hr":"0.00001322"},"BTC_STR":{"id":89,"last":"0.00001630","lowestAsk":"0.00001632","highestBid":"0.00001630","percentChange":"0.16845878","baseVolume":"3687.57249161","quoteVolume":"243056794.71035343","isFrozen":"0","high24hr":"0.00001658","low24hr":"0.00001327"},"BTC_SYS":{"id":92,"last":"0.00004184","lowestAsk":"0.00004182","highestBid":"0.00004164","percentChange":"0.24264924","baseVolume":"219.86915626","quoteVolume":"5805725.04715249","isFrozen":"0","high24hr":"0.00004391","low24hr":"0.00003231"},"BTC_VIA":{"id":97,"last":"0.00024545","lowestAsk":"0.00024545","highestBid":"0.00024445","percentChange":"0.24203015","baseVolume":"66.21117867","quoteVolume":"286195.38588828","isFrozen":"0","high24hr":"0.00025000","low24hr":"0.00019761"},"BTC_XVC":{"id":98,"last":"0.00006110","lowestAsk":"0.00006119","highestBid":"0.00006110","percentChange":"0.26658374","baseVolume":"24.96648424","quoteVolume":"476484.06768930","isFrozen":"0","high24hr":"0.00006115","low24hr":"0.00004603"},"BTC_VRC":{"id":99,"last":"0.00006416","lowestAsk":"0.00006430","highestBid":"0.00006416","percentChange":"0.10735243","baseVolume":"43.98064318","quoteVolume":"744247.81026404","isFrozen":"0","high24hr":"0.00006430","low24hr":"0.00005422"},"BTC_VTC":{"id":100,"last":"0.00058018","lowestAsk":"0.00058695","highestBid":"0.00058026","percentChange":"0.16124254","baseVolume":"370.14924417","quoteVolume":"685748.79250507","isFrozen":"0","high24hr":"0.00060192","low24hr":"0.00048545"},"BTC_XBC":{"id":104,"last":"0.00860939","lowestAsk":"0.00865264","highestBid":"0.00860939","percentChange":"0.25071765","baseVolume":"34.45931036","quoteVolume":"4546.31758170","isFrozen":"0","high24hr":"0.00865265","low24hr":"0.00665487"},"BTC_XCP":{"id":108,"last":"0.00232981","lowestAsk":"0.00232981","highestBid":"0.00231471","percentChange":"0.11020519","baseVolume":"53.09226817","quoteVolume":"24991.46627955","isFrozen":"0","high24hr":"0.00236010","low24hr":"0.00195246"},"BTC_XEM":{"id":112,"last":"0.00006006","lowestAsk":"0.00006020","highestBid":"0.00006007","percentChange":"0.08784640","baseVolume":"1700.54023119","quoteVolume":"29469631.16641163","isFrozen":"0","high24hr":"0.00006309","low24hr":"0.00005300"},"BTC_XMR":{"id":114,"last":"0.02773500","lowestAsk":"0.02785900","highestBid":"0.02772510","percentChange":"0.12652315","baseVolume":"4385.81995566","quoteVolume":"163233.23852836","isFrozen":"0","high24hr":"0.02890000","low24hr":"0.02426007"},"BTC_XPM":{"id":116,"last":"0.00003189","lowestAsk":"0.00003197","highestBid":"0.00003189","percentChange":"0.39745836","baseVolume":"34.12543829","quoteVolume":"1134318.27295924","isFrozen":"0","high24hr":"0.00003755","low24hr":"0.00002282"},"BTC_XRP":{"id":117,"last":"0.00006166","lowestAsk":"0.00006166","highestBid":"0.00006162","percentChange":"0.47476680","baseVolume":"14690.23201361","quoteVolume":"285081925.29438621","isFrozen":"0","high24hr":"0.00006275","low24hr":"0.00003913"},"USDT_BTC":{"id":121,"last":"16177.69999983","lowestAsk":"16177.69999983","highestBid":"16150.38200022","percentChange":"-0.05504088","baseVolume":"164207401.71013987","quoteVolume":"9888.66346576","isFrozen":"0","high24hr":"17375.98798753","low24hr":"15765.78546988"},"USDT_DASH":{"id":122,"last":"1469.99999990","lowestAsk":"1469.99999991","highestBid":"1469.99999990","percentChange":"0.02663953","baseVolume":"8590956.74427352","quoteVolume":"5864.54823265","isFrozen":"0","high24hr":"1535.63159991","low24hr":"1386.42020224"},"USDT_LTC":{"id":123,"last":"305.99999997","lowestAsk":"305.99999996","highestBid":"305.99000000","percentChange":"-0.05956113","baseVolume":"22698974.96397791","quoteVolume":"72870.67986988","isFrozen":"0","high24hr":"328.53600000","low24hr":"293.50000000"},"USDT_NXT":{"id":124,"last":"1.25710000","lowestAsk":"1.25800000","highestBid":"1.25710000","percentChange":"0.38294829","baseVolume":"18717987.30928298","quoteVolume":"16370224.37687810","isFrozen":"0","high24hr":"1.42600000","low24hr":"0.84648398"},"USDT_STR":{"id":125,"last":"0.26370000","lowestAsk":"0.26399999","highestBid":"0.26370000","percentChange":"0.10478812","baseVolume":"12724956.34832313","quoteVolume":"49970009.75921644","isFrozen":"0","high24hr":"0.28222000","low24hr":"0.22845786"},"USDT_XMR":{"id":126,"last":"448.29164900","lowestAsk":"448.29164900","highestBid":"447.33475126","percentChange":"0.05607116","baseVolume":"7018775.18304096","quoteVolume":"15770.60399307","isFrozen":"0","high24hr":"478.00000000","low24hr":"400.00000000"},"USDT_XRP":{"id":127,"last":"0.99500000","lowestAsk":"0.99500000","highestBid":"0.99200000","percentChange":"0.39160880","baseVolume":"54286026.08176750","quoteVolume":"63602347.14584787","isFrozen":"0","high24hr":"1.02000019","low24hr":"0.68000000"},"XMR_BCN":{"id":129,"last":"0.00001600","lowestAsk":"0.00001644","highestBid":"0.00001501","percentChange":"0.43112701","baseVolume":"46.17861923","quoteVolume":"3473353.02267603","isFrozen":"0","high24hr":"0.00001730","low24hr":"0.00001060"},"XMR_BLK":{"id":130,"last":"0.00117000","lowestAsk":"0.00137638","highestBid":"0.00116114","percentChange":"-0.15219848","baseVolume":"47.90172452","quoteVolume":"31809.80726987","isFrozen":"0","high24hr":"0.00195000","low24hr":"0.00117000"},"XMR_BTCD":{"id":131,"last":"0.61451579","lowestAsk":"0.62207953","highestBid":"0.59119801","percentChange":"0.25571856","baseVolume":"41.45651177","quoteVolume":"70.77044938","isFrozen":"0","high24hr":"0.67093769","low24hr":"0.48448008"},"XMR_DASH":{"id":132,"last":"3.25094092","lowestAsk":"3.29265865","highestBid":"3.25094092","percentChange":"-0.03082491","baseVolume":"153.17978757","quoteVolume":"45.64828073","isFrozen":"0","high24hr":"3.62853115","low24hr":"3.01702400"},"XMR_LTC":{"id":137,"last":"0.69746490","lowestAsk":"0.69718514","highestBid":"0.67844082","percentChange":"-0.10551894","baseVolume":"2908.93060635","quoteVolume":"4068.92860303","isFrozen":"0","high24hr":"0.79999997","low24hr":"0.64037953"},"XMR_MAID":{"id":138,"last":"0.00188794","lowestAsk":"0.00193604","highestBid":"0.00185746","percentChange":"0.01394759","baseVolume":"33.81049948","quoteVolume":"19280.86268167","isFrozen":"0","high24hr":"0.00198405","low24hr":"0.00152339"},"XMR_NXT":{"id":140,"last":"0.00278510","lowestAsk":"0.00285261","highestBid":"0.00279879","percentChange":"0.29173043","baseVolume":"203.51506637","quoteVolume":"80784.72150440","isFrozen":"0","high24hr":"0.00308521","low24hr":"0.00191365"},"BTC_ETH":{"id":148,"last":"0.05019493","lowestAsk":"0.05019493","highestBid":"0.05018001","percentChange":"0.07239619","baseVolume":"9045.01975596","quoteVolume":"186232.41467265","isFrozen":"0","high24hr":"0.05047987","low24hr":"0.04640017"},"USDT_ETH":{"id":149,"last":"813.68999930","lowestAsk":"813.68999930","highestBid":"810.68999989","percentChange":"0.01331257","baseVolume":"31927196.40388707","quoteVolume":"39504.32731487","isFrozen":"0","high24hr":"842.00000000","low24hr":"765.09309495"},"BTC_SC":{"id":150,"last":"0.00000153","lowestAsk":"0.00000154","highestBid":"0.00000153","percentChange":"0.47115384","baseVolume":"2207.17253784","quoteVolume":"1699661635.67689490","isFrozen":"0","high24hr":"0.00000168","low24hr":"0.00000100"},"BTC_BCY":{"id":151,"last":"0.00004000","lowestAsk":"0.00004047","highestBid":"0.00004000","percentChange":"0.11111111","baseVolume":"43.69734978","quoteVolume":"1046989.38874786","isFrozen":"0","high24hr":"0.00005148","low24hr":"0.00003464"},"BTC_EXP":{"id":153,"last":"0.00025856","lowestAsk":"0.00026012","highestBid":"0.00025856","percentChange":"0.20625145","baseVolume":"86.31078865","quoteVolume":"354421.72908786","isFrozen":"0","high24hr":"0.00027793","low24hr":"0.00021222"},"BTC_FCT":{"id":155,"last":"0.00278501","lowestAsk":"0.00280592","highestBid":"0.00278501","percentChange":"0.16434078","baseVolume":"712.23953770","quoteVolume":"272541.38753203","isFrozen":"0","high24hr":"0.00284324","low24hr":"0.00232913"},"BTC_RADS":{"id":158,"last":"0.00065761","lowestAsk":"0.00065761","highestBid":"0.00065760","percentChange":"0.06371518","baseVolume":"51.45737112","quoteVolume":"81291.66053416","isFrozen":"0","high24hr":"0.00073002","low24hr":"0.00058100"},"BTC_AMP":{"id":160,"last":"0.00003115","lowestAsk":"0.00003115","highestBid":"0.00003102","percentChange":"0.22300745","baseVolume":"65.57152453","quoteVolume":"2291924.93947091","isFrozen":"0","high24hr":"0.00003249","low24hr":"0.00002483"},"BTC_DCR":{"id":162,"last":"0.00629900","lowestAsk":"0.00629999","highestBid":"0.00622766","percentChange":"0.14550182","baseVolume":"148.59823671","quoteVolume":"25572.47329720","isFrozen":"0","high24hr":"0.00643931","low24hr":"0.00540430"},"BTC_LSK":{"id":163,"last":"0.00150002","lowestAsk":"0.00150004","highestBid":"0.00150002","percentChange":"0.31148142","baseVolume":"2139.83794369","quoteVolume":"1563369.90420573","isFrozen":"0","high24hr":"0.00170000","low24hr":"0.00111429"},"ETH_LSK":{"id":166,"last":"0.03004539","lowestAsk":"0.03044989","highestBid":"0.02999853","percentChange":"0.22834389","baseVolume":"1277.13152221","quoteVolume":"44658.48919712","isFrozen":"0","high24hr":"0.03400000","low24hr":"0.02346866"},"BTC_LBC":{"id":167,"last":"0.00004871","lowestAsk":"0.00004870","highestBid":"0.00004864","percentChange":"0.07172717","baseVolume":"174.88679988","quoteVolume":"3945032.20382236","isFrozen":"0","high24hr":"0.00004995","low24hr":"0.00003901"},"BTC_STEEM":{"id":168,"last":"0.00024429","lowestAsk":"0.00024427","highestBid":"0.00024150","percentChange":"0.46141421","baseVolume":"574.85900065","quoteVolume":"2689897.57503483","isFrozen":"0","high24hr":"0.00024996","low24hr":"0.00016701"},"ETH_STEEM":{"id":169,"last":"0.00486137","lowestAsk":"0.00499946","highestBid":"0.00499806","percentChange":"0.34403372","baseVolume":"350.65118841","quoteVolume":"78155.32467955","isFrozen":"0","high24hr":"0.00500000","low24hr":"0.00354518"},"BTC_SBD":{"id":170,"last":"0.00084674","lowestAsk":"0.00085428","highestBid":"0.00084673","percentChange":"0.16801390","baseVolume":"90.12227959","quoteVolume":"117510.30155185","isFrozen":"0","high24hr":"0.00089100","low24hr":"0.00064384"},"BTC_ETC":{"id":171,"last":"0.00247155","lowestAsk":"0.00247155","highestBid":"0.00247000","percentChange":"0.08250335","baseVolume":"1355.82621120","quoteVolume":"566991.76025183","isFrozen":"0","high24hr":"0.00249999","low24hr":"0.00226236"},"ETH_ETC":{"id":172,"last":"0.04905352","lowestAsk":"0.04966514","highestBid":"0.04905353","percentChange":"0.00949578","baseVolume":"1028.70372822","quoteVolume":"20871.36482773","isFrozen":"0","high24hr":"0.05100000","low24hr":"0.04739093"},"USDT_ETC":{"id":173,"last":"40.08600000","lowestAsk":"40.08600000","highestBid":"40.07000000","percentChange":"0.02521738","baseVolume":"11938140.28209119","quoteVolume":"299255.07740240","isFrozen":"0","high24hr":"42.55000000","low24hr":"37.50000000"},"BTC_REP":{"id":174,"last":"0.00550660","lowestAsk":"0.00550660","highestBid":"0.00550650","percentChange":"0.04126602","baseVolume":"573.95867284","quoteVolume":"104621.43511800","isFrozen":"0","high24hr":"0.00625965","low24hr":"0.00502335"},"USDT_REP":{"id":175,"last":"89.99999982","lowestAsk":"89.99499992","highestBid":"89.10000100","percentChange":"0.00000000","baseVolume":"3635421.51285649","quoteVolume":"39697.24813601","isFrozen":"0","high24hr":"104.00000000","low24hr":"82.17900000"},"ETH_REP":{"id":176,"last":"0.10962679","lowestAsk":"0.10962679","highestBid":"0.10908138","percentChange":"-0.02896330","baseVolume":"1861.69809500","quoteVolume":"15890.32221614","isFrozen":"0","high24hr":"0.13003825","low24hr":"0.10543058"},"BTC_ARDR":{"id":177,"last":"0.00007743","lowestAsk":"0.00007854","highestBid":"0.00007763","percentChange":"0.28599900","baseVolume":"401.11644739","quoteVolume":"5691643.48424611","isFrozen":"0","high24hr":"0.00008000","low24hr":"0.00005953"},"BTC_ZEC":{"id":178,"last":"0.04478162","lowestAsk":"0.04478162","highestBid":"0.04470004","percentChange":"0.23458159","baseVolume":"1407.30899657","quoteVolume":"35295.92404348","isFrozen":"0","high24hr":"0.04500000","low24hr":"0.03596966"},"ETH_ZEC":{"id":179,"last":"0.89200000","lowestAsk":"0.89262862","highestBid":"0.89113499","percentChange":"0.15675772","baseVolume":"1411.48894168","quoteVolume":"1702.20090752","isFrozen":"0","high24hr":"0.89200000","low24hr":"0.76012736"},"USDT_ZEC":{"id":180,"last":"725.00000000","lowestAsk":"726.00000000","highestBid":"725.00000000","percentChange":"0.16296557","baseVolume":"8742719.53414302","quoteVolume":"13099.14136392","isFrozen":"0","high24hr":"734.69060001","low24hr":"612.48026867"},"XMR_ZEC":{"id":181,"last":"1.61121663","lowestAsk":"1.62849087","highestBid":"1.56448069","percentChange":"0.07760972","baseVolume":"263.44276339","quoteVolume":"177.00479603","isFrozen":"0","high24hr":"1.61121663","low24hr":"1.37000009"},"BTC_STRAT":{"id":182,"last":"0.00082634","lowestAsk":"0.00082443","highestBid":"0.00082108","percentChange":"0.27135098","baseVolume":"1249.65545209","quoteVolume":"1684907.29696500","isFrozen":"0","high24hr":"0.00085328","low24hr":"0.00063300"},"BTC_NXC":{"id":183,"last":"0.00002236","lowestAsk":"0.00002242","highestBid":"0.00002236","percentChange":"0.23195592","baseVolume":"34.79455338","quoteVolume":"1693481.55824669","isFrozen":"0","high24hr":"0.00002279","low24hr":"0.00001786"},"BTC_PASC":{"id":184,"last":"0.00009407","lowestAsk":"0.00009424","highestBid":"0.00009320","percentChange":"0.15565110","baseVolume":"29.51912169","quoteVolume":"343116.34795877","isFrozen":"0","high24hr":"0.00009420","low24hr":"0.00008098"},"BTC_GNT":{"id":185,"last":"0.00004328","lowestAsk":"0.00004370","highestBid":"0.00004328","percentChange":"0.31470230","baseVolume":"448.20036311","quoteVolume":"11314202.81810377","isFrozen":"0","high24hr":"0.00004526","low24hr":"0.00003290"},"ETH_GNT":{"id":186,"last":"0.00086999","lowestAsk":"0.00086999","highestBid":"0.00086043","percentChange":"0.24449625","baseVolume":"1820.36619738","quoteVolume":"2176648.09893840","isFrozen":"0","high24hr":"0.00094000","low24hr":"0.00069907"},"BTC_GNO":{"id":187,"last":"0.01325338","lowestAsk":"0.01325330","highestBid":"0.01325000","percentChange":"0.08195273","baseVolume":"71.74995630","quoteVolume":"5622.44154787","isFrozen":"0","high24hr":"0.01356787","low24hr":"0.01181288"},"ETH_GNO":{"id":188,"last":"0.26499943","lowestAsk":"0.26499995","highestBid":"0.25717162","percentChange":"0.00340556","baseVolume":"174.43119102","quoteVolume":"662.53332874","isFrozen":"0","high24hr":"0.27999998","low24hr":"0.24800003"},"BTC_BCH":{"id":189,"last":"0.20886994","lowestAsk":"0.20887954","highestBid":"0.20886994","percentChange":"0.01000127","baseVolume":"22310.61758265","quoteVolume":"100979.21790905","isFrozen":"0","high24hr":"0.25700000","low24hr":"0.18900000"},"ETH_BCH":{"id":190,"last":"4.11648930","lowestAsk":"4.18557927","highestBid":"4.11648940","percentChange":"-0.06003714","baseVolume":"16495.99964885","quoteVolume":"3543.16856193","isFrozen":"0","high24hr":"5.31822147","low24hr":"3.88516404"},"USDT_BCH":{"id":191,"last":"3380.00000006","lowestAsk":"3380.00000006","highestBid":"3380.00000005","percentChange":"-0.03756402","baseVolume":"98924759.01870164","quoteVolume":"26920.70012331","isFrozen":"0","high24hr":"4211.99999995","low24hr":"3227.00000001"},"BTC_ZRX":{"id":192,"last":"0.00003860","lowestAsk":"0.00003860","highestBid":"0.00003846","percentChange":"0.25691957","baseVolume":"200.99792280","quoteVolume":"5829184.76006624","isFrozen":"0","high24hr":"0.00003899","low24hr":"0.00003046"},"ETH_ZRX":{"id":193,"last":"0.00076719","lowestAsk":"0.00077632","highestBid":"0.00076573","percentChange":"0.17090703","baseVolume":"737.98198452","quoteVolume":"1038418.25605133","isFrozen":"0","high24hr":"0.00077607","low24hr":"0.00065193"},"BTC_CVC":{"id":194,"last":"0.00004578","lowestAsk":"0.00004575","highestBid":"0.00004551","percentChange":"0.32772621","baseVolume":"222.26570392","quoteVolume":"5271350.60481972","isFrozen":"0","high24hr":"0.00004778","low24hr":"0.00003431"},"ETH_CVC":{"id":195,"last":"0.00094423","lowestAsk":"0.00094421","highestBid":"0.00090125","percentChange":"0.29169630","baseVolume":"495.80058437","quoteVolume":"577888.11817774","isFrozen":"0","high24hr":"0.00098050","low24hr":"0.00073100"},"BTC_OMG":{"id":196,"last":"0.00111300","lowestAsk":"0.00111300","highestBid":"0.00111005","percentChange":"0.06065659","baseVolume":"501.45943151","quoteVolume":"469211.03616555","isFrozen":"0","high24hr":"0.00114300","low24hr":"0.00100000"},"ETH_OMG":{"id":197,"last":"0.02232617","lowestAsk":"0.02265118","highestBid":"0.02230260","percentChange":"-0.01260657","baseVolume":"946.13758336","quoteVolume":"42983.27576155","isFrozen":"0","high24hr":"0.02306244","low24hr":"0.02110001"},"BTC_GAS":{"id":198,"last":"0.00208901","lowestAsk":"0.00210986","highestBid":"0.00208901","percentChange":"0.10212404","baseVolume":"99.43409604","quoteVolume":"49059.66525670","isFrozen":"0","high24hr":"0.00219998","low24hr":"0.00183559"},"ETH_GAS":{"id":199,"last":"0.04178195","lowestAsk":"0.04278189","highestBid":"0.04178194","percentChange":"0.02969209","baseVolume":"143.77661265","quoteVolume":"3434.98447748","isFrozen":"0","high24hr":"0.04465112","low24hr":"0.03875406"},"BTC_STORJ":{"id":200,"last":"0.00010583","lowestAsk":"0.00010584","highestBid":"0.00010500","percentChange":"0.14547028","baseVolume":"142.51217975","quoteVolume":"1423206.07107793","isFrozen":"0","high24hr":"0.00011190","low24hr":"0.00008850"}}`)

func BenchmarkParsePoloniexData(b *testing.B) {
	b.SetBytes(int64(len(poloniexData)))

	for i := 0; i < b.N; i++ {
		if _, err := ParsePoloniexData(poloniexData); nil != err {
			b.Error(err)
		}
	}
}
