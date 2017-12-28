package parsers

import "testing"

var exmoData = []byte(`{"BTC_USD":{"buy_price":"16009.999999","sell_price":"16011","last_trade":"16011","high":"17550","low":"16000","avg":"16651.35882512","vol":"1504.75732379","vol_curr":"24091164.75238141","updated":1513864435},"BTC_EUR":{"buy_price":"14230.8382863","sell_price":"14359.204005","last_trade":"14359.204005","high":"15699.999999","low":"14280.6462203","avg":"14830.90361973","vol":"320.95945611","vol_curr":"4608722.30773252","updated":1513864433},"BTC_RUB":{"buy_price":"966000","sell_price":"969625.65","last_trade":"969625.65","high":"1094900","low":"965000.01","avg":"1018412.02565903","vol":"818.41322241","vol_curr":"793554452.75520563","updated":1513864436},"BTC_UAH":{"buy_price":"468000","sell_price":"468956.357947","last_trade":"468956.357947","high":"513000","low":"468000","avg":"492605.69497631","vol":"87.55983235","vol_curr":"41061740.08487725","updated":1513864435},"BTC_PLN":{"buy_price":"59600.08000007","sell_price":"64007","last_trade":"62146","high":"65798.99","low":"58524.01739999","avg":"61812.4239634","vol":"60.16690714","vol_curr":"3739132.61148545","updated":1513864413},"BCH_BTC":{"buy_price":"0.20822064","sell_price":"0.209","last_trade":"0.209","high":"0.258","low":"0.18324207","avg":"0.21926358","vol":"3652.21478814","vol_curr":"763.31289072","updated":1513864361},"BCH_USD":{"buy_price":"3350.01","sell_price":"3362.74","last_trade":"3355.01","high":"4199","low":"3200","avg":"3680.69313037","vol":"3780.98989434","vol_curr":"12685258.90543927","updated":1513864433},"BCH_RUB":{"buy_price":"201930.33803","sell_price":"203041.156115","last_trade":"203041.156115","high":"265000","low":"194000","avg":"224978.52924452","vol":"443.94348711","vol_curr":"90138798.87262104","updated":1513864433},"BCH_ETH":{"buy_price":"4.04053495","sell_price":"4.14882995","last_trade":"4.11515586","high":"5.98","low":"3.82800011","avg":"4.55071386","vol":"290.52417184","vol_curr":"1195.55224822","updated":1513864417},"DASH_BTC":{"buy_price":"0.09032599","sell_price":"0.09112301","last_trade":"0.09112302","high":"0.09571","low":"0.08317498","avg":"0.08809606","vol":"2819.78334442","vol_curr":"256.94717408","updated":1513864436},"DASH_USD":{"buy_price":"1459.53","sell_price":"1460","last_trade":"1460","high":"1540","low":"1400.290322","avg":"1473.6825586","vol":"2084.34593553","vol_curr":"3043145.06587962","updated":1513864436},"DASH_RUB":{"buy_price":"88020","sell_price":"89742.2156701","last_trade":"89429.2134231","high":"94939.7094624","low":"82817","avg":"90129.30928915","vol":"665.94813415","vol_curr":"59555217.81810127","updated":1513864436},"ETH_BTC":{"buy_price":"0.05025329","sell_price":"0.0504539","last_trade":"0.0504539","high":"0.05046403","low":"0.04631271","avg":"0.04846906","vol":"9235.65126492","vol_curr":"465.97462535","updated":1513864436},"ETH_LTC":{"buy_price":"2.61465148","sell_price":"2.62700918","last_trade":"2.6155","high":"2.6926824","low":"2.41933938","avg":"2.55736605","vol":"533.54119345","vol_curr":"1395.47699148","updated":1513864422},"ETH_USD":{"buy_price":"810","sell_price":"816.18474199","last_trade":"810.00000001","high":"838","low":"775","avg":"810.8034186","vol":"7011.5142591","vol_curr":"5679326.54987651","updated":1513864436},"ETH_EUR":{"buy_price":"721.9061","sell_price":"726.97305976","last_trade":"721.9051","high":"739.78445476","low":"695.73630545","avg":"718.8332394","vol":"1387.41862359","vol_curr":"1001584.5802082","updated":1513864432},"ETH_RUB":{"buy_price":"48500.01","sell_price":"48913","last_trade":"48913","high":"51000","low":"47000","avg":"49507.10979477","vol":"1556.17484458","vol_curr":"76117180.17313899","updated":1513864434},"ETH_UAH":{"buy_price":"23490","sell_price":"23650","last_trade":"23486.0059589","high":"24500","low":"22900","avg":"23872.25455099","vol":"1067.25895958","vol_curr":"25065650.28448047","updated":1513864431},"ETH_PLN":{"buy_price":"2954.6","sell_price":"3149","last_trade":"3052.4","high":"3143.3","low":"2855.01","avg":"3020.45094827","vol":"310.85815492","vol_curr":"948863.43210756","updated":1513864418},"ETC_BTC":{"buy_price":"0.00245792","sell_price":"0.002473","last_trade":"0.002473","high":"0.00247594","low":"0.002162","avg":"0.0023695","vol":"22174.88929351","vol_curr":"54.83850122","updated":1513864435},"ETC_USD":{"buy_price":"39.51000019","sell_price":"39.9","last_trade":"39.901411","high":"41.2","low":"37.5","avg":"39.58074706","vol":"48463.53906726","vol_curr":"1933763.59083752","updated":1513843723},"ETC_RUB":{"buy_price":"2400.00000001","sell_price":"2426.82999959","last_trade":"2400.00000001","high":"2500","low":"2312.5","avg":"2417.81644426","vol":"7284.89330687","vol_curr":"17483743.93657275","updated":1513864433},"LTC_BTC":{"buy_price":"0.0191166","sell_price":"0.01926395","last_trade":"0.01911657","high":"0.01928899","low":"0.01841219","avg":"0.01890606","vol":"10509.22017498","vol_curr":"200.90024312","updated":1513864435},"LTC_USD":{"buy_price":"307.5","sell_price":"309.99999995","last_trade":"307.5","high":"332.9","low":"301.37783325","avg":"316.07545712","vol":"9630.6516649","vol_curr":"2961425.38695678","updated":1513864435},"LTC_EUR":{"buy_price":"275","sell_price":"276.61259048","last_trade":"275","high":"296.63330778","low":"272.23110272","avg":"281.20363135","vol":"1184.59963481","vol_curr":"325764.89957531","updated":1513864433},"LTC_RUB":{"buy_price":"18551.2499773","sell_price":"18718.5868899","last_trade":"18616.1793522","high":"20284.8932553","low":"18500","avg":"19313.11266163","vol":"1974.3597734","vol_curr":"36755035.64743922","updated":1513864422},"ZEC_BTC":{"buy_price":"0.04461113","sell_price":"0.0448499","last_trade":"0.0448499","high":"0.04485","low":"0.033667","avg":"0.03970776","vol":"2804.62638802","vol_curr":"125.78721304","updated":1513864436},"ZEC_USD":{"buy_price":"715.9","sell_price":"717.98","last_trade":"715.9","high":"723.62","low":"616.000001","avg":"663.79226417","vol":"5034.45239599","vol_curr":"3604164.47029408","updated":1513864436},"ZEC_EUR":{"buy_price":"631.93954325","sell_price":"637.61","last_trade":"633.412","high":"640","low":"555.30572845","avg":"587.34984933","vol":"712.73366035","vol_curr":"451454.05327044","updated":1513864433},"ZEC_RUB":{"buy_price":"42886.30754126","sell_price":"43236.61994277","last_trade":"43159.02989697","high":"43524","low":"36296.1","avg":"40364.84891183","vol":"1228.63664365","vol_curr":"53026765.63606512","updated":1513864434},"XRP_BTC":{"buy_price":"0.00006123","sell_price":"0.00006148","last_trade":"0.00006122","high":"0.00006276","low":"0.00004052","avg":"0.00005087","vol":"8509605.72950585","vol_curr":"523.17056025","updated":1513864436},"XRP_USD":{"buy_price":"0.98","sell_price":"0.98191","last_trade":"0.978","high":"0.99999999","low":"0.69","avg":"0.85010706","vol":"15218325.67876651","vol_curr":"14913959.16519118","updated":1513864436},"XRP_RUB":{"buy_price":"59.2862","sell_price":"59.39999999","last_trade":"59.2862","high":"60","low":"42.94082057","avg":"51.49661516","vol":"3590138.80895397","vol_curr":"212845687.45540703","updated":1513864420},"XMR_BTC":{"buy_price":"0.0275001","sell_price":"0.02759849","last_trade":"0.02759849","high":"0.02844658","low":"0.02411079","avg":"0.02645747","vol":"2681.06703247","vol_curr":"73.99340168","updated":1513864436},"XMR_USD":{"buy_price":"442.91281732","sell_price":"443","last_trade":"442.91281732","high":"470","low":"407.287616","avg":"442.16025517","vol":"3901.02863048","vol_curr":"1727815.58117397","updated":1513864435},"XMR_EUR":{"buy_price":"393.19582235","sell_price":"396.70849061","last_trade":"394.9","high":"418","low":"366.65816318","avg":"392.619798","vol":"541.25062158","vol_curr":"213739.87046516","updated":1513864301},"BTC_USDT":{"buy_price":"16150","sell_price":"16168.4368818","last_trade":"16150","high":"17896","low":"16000","avg":"16592.03982567","vol":"57.9308338","vol_curr":"935582.96594325","updated":1513864436},"ETH_USDT":{"buy_price":"812.74305036","sell_price":"813.964581","last_trade":"813.96459134","high":"835.78058649","low":"775.10803211","avg":"805.5457075","vol":"134.05675332","vol_curr":"109117.4504342","updated":1513864421},"USDT_USD":{"buy_price":"0.9941","sell_price":"0.9994","last_trade":"0.9941","high":"1.03597607","low":"0.98501","avg":"1.00242746","vol":"238459.60090684","vol_curr":"237052.68926149","updated":1513864435},"USDT_RUB":{"buy_price":"60","sell_price":"60.46581573","last_trade":"60.01221986","high":"62.38","low":"59.6","avg":"61.1488763","vol":"116659.40985138","vol_curr":"7000990.15273945","updated":1513864421},"USD_RUB":{"buy_price":"60.4","sell_price":"60.42","last_trade":"60.4","high":"62.2","low":"60","avg":"61.12921497","vol":"1398046.76357948","vol_curr":"84442024.52020061","updated":1513864432},"DOGE_BTC":{"buy_price":"0.00000056","sell_price":"0.00000057","last_trade":"0.00000059","high":"0.00000063","low":"0.00000031","avg":"0.00000043","vol":"221508402.25377761","vol_curr":"130.68995732","updated":1513864436},"WAVES_BTC":{"buy_price":"0.001","sell_price":"0.00100997","last_trade":"0.001","high":"0.00102959","low":"0.000861","avg":"0.00093394","vol":"201953.02869806","vol_curr":"201.96110681","updated":1513864436},"WAVES_RUB":{"buy_price":"981","sell_price":"988.88011098","last_trade":"980.35898792","high":"992.999934","low":"880","avg":"946.97208546","vol":"43037.52344892","vol_curr":"42192222.93096794","updated":1513841783},"KICK_BTC":{"buy_price":"0.00000365","sell_price":"0.00000369","last_trade":"0.00000369","high":"0.00000372","low":"0.000003","avg":"0.00000342","vol":"6557338.76198714","vol_curr":"24.19658003","updated":1513854857},"KICK_ETH":{"buy_price":"0.00007307","sell_price":"0.00007398","last_trade":"0.00007368","high":"0.00007489","low":"0.00006274","avg":"0.00007046","vol":"3726383.05169024","vol_curr":"274.55990324","updated":1513864397}}`)

func BenchmarkParseExmoData(b *testing.B) {
	b.SetBytes(int64(len(exmoData)))

	for i := 0; i < b.N; i++ {
		if _, err := ParseExmoData(exmoData); nil != err {
			b.Error(err)
		}
	}
}