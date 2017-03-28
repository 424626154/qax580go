// JavaScript Document
var shownum = 0,//第几个文字列表
	allnum = 0,//通关数
	allnumshow = 0,
	time_num = 600, //翻页延迟
	clock = null,
	sec = 10, //倒计时
	show = [{"info":{"index":42,"j":"\u5e01"},"list":[{"index":42,"f":"\u5e63"},{"index":43,"f":"\u958b"},{"index":44,"f":"\u61b6"},{"index":45,"f":"\u6236"},{"index":46,"f":"\u7d2e"},{"index":47,"f":"\u9b25"}]},{"info":{"index":260,"j":"\u95eb"},"list":[{"index":260,"f":"\u9586"},{"index":261,"f":"\u9589"},{"index":262,"f":"\u554f"},{"index":263,"f":"\u95d6"},{"index":264,"f":"\u967d"},{"index":265,"f":"\u9670"}]},{"info":{"index":658,"j":"\u866e"},"list":[{"index":658,"f":"\u87e3"},{"index":659,"f":"\u896f"},{"index":660,"f":"\u898f"},{"index":661,"f":"\u8993"},{"index":662,"f":"\u8996"},{"index":663,"f":"\u8a86"}]},{"info":{"index":847,"j":"\u6b8b"},"list":[{"index":847,"f":"\u6b98"},{"index":848,"f":"\u6c08"},{"index":849,"f":"\u6c2b"},{"index":850,"f":"\u6fa9"},{"index":851,"f":"\u6f54"},{"index":852,"f":"\u7051"}]},{"info":{"index":1003,"j":"\u949f"},"list":[{"index":1003,"f":"\u937e"},{"index":1004,"f":"\u9209"},{"index":1005,"f":"\u92c7"},{"index":1006,"f":"\u92fc"},{"index":1007,"f":"\u9211"},{"index":1008,"f":"\u9210"}]},{"info":{"index":1360,"j":"\u63b7"},"list":[{"index":1360,"f":"\u64f2"},{"index":1361,"f":"\u64a3"},{"index":1362,"f":"\u647b"},{"index":1363,"f":"\u645c"},{"index":1364,"f":"\u6582"},{"index":1365,"f":"\u65b7"}]},{"info":{"index":1522,"j":"\u94f1"},"list":[{"index":1522,"f":"\u92a5"},{"index":1523,"f":"\u93df"},{"index":1524,"f":"\u9283"},{"index":1525,"f":"\u940b"},{"index":1526,"f":"\u92a8"},{"index":1527,"f":"\u9280"}]},{"info":{"index":1780,"j":"\u5ad4"},"list":[{"index":1780,"f":"\u5b2a"},{"index":1781,"f":"\u5be2"},{"index":1782,"f":"\u5c37"},{"index":1783,"f":"\u810a"},{"index":1784,"f":"\u6e63"},{"index":1785,"f":"\u61fe"}]},{"info":{"index":1997,"j":"\u8c31"},"list":[{"index":1997,"f":"\u8b5c"},{"index":1998,"f":"\u8b4e"},{"index":1999,"f":"\u8d05"},{"index":2000,"f":"\u8cfb"},{"index":2001,"f":"\u8cfa"},{"index":2002,"f":"\u8cfd"}]},{"info":{"index":2195,"j":"\u9cad"},"list":[{"index":2195,"f":"\u9bd6"},{"index":2196,"f":"\u9bea"},{"index":2197,"f":"\u9beb"},{"index":2198,"f":"\u9be1"},{"index":2199,"f":"\u9be4"},{"index":2200,"f":"\u9be7"}]},{"info":{"index":205,"j":"\u626b"},"list":[{"index":205,"f":"\u6383"},{"index":206,"f":"\u63da"},{"index":207,"f":"\u6a38"},{"index":208,"f":"\u6a5f"},{"index":209,"f":"\u6bba"},{"index":210,"f":"\u96dc"}]},{"info":{"index":323,"j":"\u5760"},"list":[{"index":323,"f":"\u589c"},{"index":324,"f":"\u8072"},{"index":325,"f":"\u6bbc"},{"index":326,"f":"\u5969"},{"index":327,"f":"\u5950"},{"index":328,"f":"\u5af5"}]},{"info":{"index":650,"j":"\u82f9"},"list":[{"index":650,"f":"\u860b"},{"index":651,"f":"\u7bc4"},{"index":652,"f":"\u8396"},{"index":653,"f":"\u8622"},{"index":654,"f":"\u8526"},{"index":655,"f":"\u584b"}]},{"info":{"index":745,"j":"\u9e22"},"list":[{"index":745,"f":"\u9cf6"},{"index":746,"f":"\u9cf4"},{"index":747,"f":"\u9efd"},{"index":748,"f":"\u9f52"},{"index":749,"f":"\u81e8"},{"index":750,"f":"\u8209"}]},{"info":{"index":983,"j":"\u8f73"},"list":[{"index":983,"f":"\u8f64"},{"index":984,"f":"\u8ef8"},{"index":985,"f":"\u8ef9"},{"index":986,"f":"\u8efc"},{"index":987,"f":"\u8ee4"},{"index":988,"f":"\u8eeb"}]},{"info":{"index":1385,"j":"\u7315"},"list":[{"index":1385,"f":"\u737c"},{"index":1386,"f":"\u7380"},{"index":1387,"f":"\u8c6c"},{"index":1388,"f":"\u8c93"},{"index":1389,"f":"\u7489"},{"index":1390,"f":"\u7463"}]},{"info":{"index":1593,"j":"\u6401"},"list":[{"index":1593,"f":"\u64f1"},{"index":1594,"f":"\u645f"},{"index":1595,"f":"\u652a"},{"index":1596,"f":"\u66ab"},{"index":1597,"f":"\u69e8"},{"index":1598,"f":"\u6add"}]},{"info":{"index":1632,"j":"\u7b5b"},"list":[{"index":1632,"f":"\u7be9"},{"index":1633,"f":"\u7b8f"},{"index":1634,"f":"\u7cb5"},{"index":1635,"f":"\u7cde"},{"index":1636,"f":"\u7e36"},{"index":1637,"f":"\u7dd9"}]},{"info":{"index":2013,"j":"\u9534"},"list":[{"index":2013,"f":"\u9347"},{"index":2014,"f":"\u93d8"},{"index":2015,"f":"\u9376"},{"index":2016,"f":"\u9354"},{"index":2017,"f":"\u9364"},{"index":2018,"f":"\u936c"}]},{"info":{"index":2160,"j":"\u736d"},"list":[{"index":2160,"f":"\u737a"},{"index":2161,"f":"\u766e"},{"index":2162,"f":"\u766d"},{"index":2163,"f":"\u7a61"},{"index":2164,"f":"\u7c43"},{"index":2165,"f":"\u7c6c"}]},{"info":{"index":214,"j":"\u6c61"},"list":[{"index":214,"f":"\u6c59"},{"index":215,"f":"\u6e6f"},{"index":216,"f":"\u71c8"},{"index":217,"f":"\u723a"},{"index":218,"f":"\u7377"},{"index":219,"f":"\u7341"}]},{"info":{"index":319,"j":"\u575c"},"list":[{"index":319,"f":"\u58e2"},{"index":320,"f":"\u58e9"},{"index":321,"f":"\u5862"},{"index":322,"f":"\u58b3"},{"index":323,"f":"\u589c"},{"index":324,"f":"\u8072"}]},{"info":{"index":674,"j":"\u8bde"},"list":[{"index":674,"f":"\u8a95"},{"index":675,"f":"\u8a6c"},{"index":676,"f":"\u8a6e"},{"index":677,"f":"\u8a6d"},{"index":678,"f":"\u8a62"},{"index":679,"f":"\u8a63"}]},{"info":{"index":803,"j":"\u5ce4"},"list":[{"index":803,"f":"\u5da0"},{"index":804,"f":"\u5d22"},{"index":805,"f":"\u5dd2"},{"index":806,"f":"\u5e36"},{"index":807,"f":"\u5e40"},{"index":808,"f":"\u5e6b"}]},{"info":{"index":1045,"j":"\u9a84"},"list":[{"index":1045,"f":"\u9a55"},{"index":1046,"f":"\u9a4a"},{"index":1047,"f":"\u99f1"},{"index":1048,"f":"\u99ed"},{"index":1049,"f":"\u99e2"},{"index":1050,"f":"\u9dd7"}]},{"info":{"index":1239,"j":"\u8d44"},"list":[{"index":1239,"f":"\u8cc7"},{"index":1240,"f":"\u8cc5"},{"index":1241,"f":"\u8d10"},{"index":1242,"f":"\u8d95"},{"index":1243,"f":"\u8e89"},{"index":1244,"f":"\u8efe"}]},{"info":{"index":1529,"j":"\u9608"},"list":[{"index":1529,"f":"\u95be"},{"index":1530,"f":"\u95b9"},{"index":1531,"f":"\u95b6"},{"index":1532,"f":"\u9b29"},{"index":1533,"f":"\u95bf"},{"index":1534,"f":"\u95bd"}]},{"info":{"index":1671,"j":"\u86f1"},"list":[{"index":1671,"f":"\u86fa"},{"index":1672,"f":"\u87ef"},{"index":1673,"f":"\u8784"},{"index":1674,"f":"\u8810"},{"index":1675,"f":"\u88dd"},{"index":1676,"f":"\u8933"}]},{"info":{"index":1932,"j":"\u9e51"},"list":[{"index":1932,"f":"\u9d89"},{"index":1933,"f":"\u9f5f"},{"index":1934,"f":"\u9f61"},{"index":1935,"f":"\u9f59"},{"index":1936,"f":"\u9f60"},{"index":1937,"f":"\u5edd"}]},{"info":{"index":2148,"j":"\u61b7"},"list":[{"index":2148,"f":"\u6035"},{"index":2149,"f":"\u61f6"},{"index":2150,"f":"\u61cd"},{"index":2151,"f":"\u64fb"},{"index":2152,"f":"\u6595"},{"index":2153,"f":"\u6ae5"}]},{"info":{"index":66,"j":"\u4e1b"},"list":[{"index":66,"f":"\u53e2"},{"index":67,"f":"\u6771"},{"index":68,"f":"\u7d72"},{"index":69,"f":"\u6a02"},{"index":70,"f":"\u5100"},{"index":71,"f":"\u5011"}]},{"info":{"index":404,"j":"\u7eb2"},"list":[{"index":404,"f":"\u7db1"},{"index":405,"f":"\u7d0d"},{"index":406,"f":"\u7e31"},{"index":407,"f":"\u7db8"},{"index":408,"f":"\u7d1b"},{"index":409,"f":"\u7d19"}]},{"info":{"index":649,"j":"\u82d8"},"list":[{"index":649,"f":"\u6abe"},{"index":650,"f":"\u860b"},{"index":651,"f":"\u7bc4"},{"index":652,"f":"\u8396"},{"index":653,"f":"\u8622"},{"index":654,"f":"\u8526"}]},{"info":{"index":891,"j":"\u781c"},"list":[{"index":891,"f":"\u78b8"},{"index":892,"f":"\u79b0"},{"index":893,"f":"\u7a2e"},{"index":894,"f":"\u7aca"},{"index":895,"f":"\u8c4e"},{"index":896,"f":"\u7be4"}]},{"info":{"index":995,"j":"\u900a"},"list":[{"index":995,"f":"\u905c"},{"index":996,"f":"\u9148"},{"index":997,"f":"\u9116"},{"index":998,"f":"\u9223"},{"index":999,"f":"\u9208"},{"index":1000,"f":"\u9226"}]},{"info":{"index":1355,"j":"\u60ef"},"list":[{"index":1355,"f":"\u6163"},{"index":1356,"f":"\u64da"},{"index":1357,"f":"\u649a"},{"index":1358,"f":"\u64c4"},{"index":1359,"f":"\u6451"},{"index":1360,"f":"\u64f2"}]},{"info":{"index":1529,"j":"\u9608"},"list":[{"index":1529,"f":"\u95be"},{"index":1530,"f":"\u95b9"},{"index":1531,"f":"\u95b6"},{"index":1532,"f":"\u9b29"},{"index":1533,"f":"\u95bf"},{"index":1534,"f":"\u95bd"}]},{"info":{"index":1700,"j":"\u8d8b"},"list":[{"index":1700,"f":"\u8da8"},{"index":1701,"f":"\u8e60"},{"index":1702,"f":"\u8e92"},{"index":1703,"f":"\u8e10"},{"index":1704,"f":"\u8f26"},{"index":1705,"f":"\u8f29"}]},{"info":{"index":1966,"j":"\u7ba7"},"list":[{"index":1966,"f":"\u7bcb"},{"index":1967,"f":"\u7c5c"},{"index":1968,"f":"\u7c6e"},{"index":1969,"f":"\u7c1e"},{"index":1970,"f":"\u7c2b"},{"index":1971,"f":"\u7cdd"}]},{"info":{"index":2274,"j":"\u56af"},"list":[{"index":2274,"f":"\u8b14"},{"index":2275,"f":"\u5dd4"},{"index":2276,"f":"\u6522"},{"index":2277,"f":"\u766c"},{"index":2278,"f":"\u7c5f"},{"index":2279,"f":"\u7e98"}]},{"info":{"index":30,"j":"\u4ed1"},"list":[{"index":30,"f":"\u4f96"},{"index":31,"f":"\u5009"},{"index":32,"f":"\u5167"},{"index":33,"f":"\u5ca1"},{"index":34,"f":"\u9cf3"},{"index":35,"f":"\u52f8"}]},{"info":{"index":464,"j":"\u948a"},"list":[{"index":464,"f":"\u91d7"},{"index":465,"f":"\u91d9"},{"index":466,"f":"\u91d5"},{"index":467,"f":"\u958f"},{"index":468,"f":"\u95c8"},{"index":469,"f":"\u9591"}]},{"info":{"index":641,"j":"\u8083"},"list":[{"index":641,"f":"\u8085"},{"index":642,"f":"\u819a"},{"index":643,"f":"\u8181"},{"index":644,"f":"\u814e"},{"index":645,"f":"\u816b"},{"index":646,"f":"\u8139"}]},{"info":{"index":710,"j":"\u90d1"},"list":[{"index":710,"f":"\u912d"},{"index":711,"f":"\u9106"},{"index":712,"f":"\u91f7"},{"index":713,"f":"\u91fa"},{"index":714,"f":"\u91e7"},{"index":715,"f":"\u91e4"}]},{"info":{"index":1101,"j":"\u6653"},"list":[{"index":1101,"f":"\u66c9"},{"index":1102,"f":"\u66c4"},{"index":1103,"f":"\u6688"},{"index":1104,"f":"\u6689"},{"index":1105,"f":"\u68f2"},{"index":1106,"f":"\u6a23"}]},{"info":{"index":1222,"j":"\u8c01"},"list":[{"index":1222,"f":"\u8ab0"},{"index":1223,"f":"\u8ad7"},{"index":1224,"f":"\u8abf"},{"index":1225,"f":"\u8ac2"},{"index":1226,"f":"\u8ad2"},{"index":1227,"f":"\u8ac4"}]},{"info":{"index":1540,"j":"\u9885"},"list":[{"index":1540,"f":"\u9871"},{"index":1541,"f":"\u9818"},{"index":1542,"f":"\u9817"},{"index":1543,"f":"\u9838"},{"index":1544,"f":"\u991b"},{"index":1545,"f":"\u9921"}]},{"info":{"index":1751,"j":"\u988c"},"list":[{"index":1751,"f":"\u981c"},{"index":1752,"f":"\u6f41"},{"index":1753,"f":"\u9826"},{"index":1754,"f":"\u98b6"},{"index":1755,"f":"\u9957"},{"index":1756,"f":"\u9937"}]},{"info":{"index":1971,"j":"\u7cc1"},"list":[{"index":1971,"f":"\u7cdd"},{"index":1972,"f":"\u7e39"},{"index":1973,"f":"\u7e35"},{"index":1974,"f":"\u7e32"},{"index":1975,"f":"\u7e93"},{"index":1976,"f":"\u7e2e"}]},{"info":{"index":2228,"j":"\u9561"},"list":[{"index":2228,"f":"\u9414"},{"index":2229,"f":"\u9481"},{"index":2230,"f":"\u9410"},{"index":2231,"f":"\u93f7"},{"index":2232,"f":"\u9465"},{"index":2233,"f":"\u9413"}]},{"info":{"index":199,"j":"\u5fcf"},"list":[{"index":199,"f":"\u61fa"},{"index":200,"f":"\u6232"},{"index":201,"f":"\u6261"},{"index":202,"f":"\u57f7"},{"index":203,"f":"\u64f4"},{"index":204,"f":"\u636b"}]},{"info":{"index":383,"j":"\u6ca9"},"list":[{"index":383,"f":"\u6e88"},{"index":384,"f":"\u6eec"},{"index":385,"f":"\u9748"},{"index":386,"f":"\u707d"},{"index":387,"f":"\u71e6"},{"index":388,"f":"\u716c"}]},{"info":{"index":615,"j":"\u75a1"},"list":[{"index":615,"f":"\u760d"},{"index":616,"f":"\u792c"},{"index":617,"f":"\u7926"},{"index":618,"f":"\u78ad"},{"index":619,"f":"\u78bc"},{"index":620,"f":"\u7a08"}]},{"info":{"index":708,"j":"\u90cf"},"list":[{"index":708,"f":"\u90df"},{"index":709,"f":"\u9136"},{"index":710,"f":"\u912d"},{"index":711,"f":"\u9106"},{"index":712,"f":"\u91f7"},{"index":713,"f":"\u91fa"}]},{"info":{"index":1003,"j":"\u949f"},"list":[{"index":1003,"f":"\u937e"},{"index":1004,"f":"\u9209"},{"index":1005,"f":"\u92c7"},{"index":1006,"f":"\u92fc"},{"index":1007,"f":"\u9211"},{"index":1008,"f":"\u9210"}]},{"info":{"index":1331,"j":"\u556c"},"list":[{"index":1331,"f":"\u55c7"},{"index":1332,"f":"\u56c0"},{"index":1333,"f":"\u9f67"},{"index":1334,"f":"\u562f"},{"index":1335,"f":"\u588a"},{"index":1336,"f":"\u57b5"}]},{"info":{"index":1419,"j":"\u7eed"},"list":[{"index":1419,"f":"\u7e8c"},{"index":1420,"f":"\u7dba"},{"index":1421,"f":"\u7dcb"},{"index":1422,"f":"\u7dbd"},{"index":1423,"f":"\u7dd4"},{"index":1424,"f":"\u7dc4"}]},{"info":{"index":1695,"j":"\u8d4f"},"list":[{"index":1695,"f":"\u8cde"},{"index":1696,"f":"\u8cdc"},{"index":1697,"f":"\u8ce1"},{"index":1698,"f":"\u8ce0"},{"index":1699,"f":"\u8ce7"},{"index":1700,"f":"\u8da8"}]},{"info":{"index":2003,"j":"\u8e0a"},"list":[{"index":2003,"f":"\u8e34"},{"index":2004,"f":"\u8e8a"},{"index":2005,"f":"\u8f45"},{"index":2006,"f":"\u8f44"},{"index":2007,"f":"\u8f3e"},{"index":2008,"f":"\u91c5"}]},{"info":{"index":2150,"j":"\u61d4"},"list":[{"index":2150,"f":"\u61cd"},{"index":2151,"f":"\u64fb"},{"index":2152,"f":"\u6595"},{"index":2153,"f":"\u6ae5"},{"index":2154,"f":"\u6ad3"},{"index":2155,"f":"\u6ade"}]},{"info":{"index":123,"j":"\u9965"},"list":[{"index":123,"f":"\u9951"},{"index":124,"f":"\u99ad"},{"index":125,"f":"\u9ce5"},{"index":126,"f":"\u9f8d"},{"index":127,"f":"\u4e1f"},{"index":128,"f":"\u55ac"}]},{"info":{"index":397,"j":"\u7a77"},"list":[{"index":397,"f":"\u7aae"},{"index":398,"f":"\u4fc2"},{"index":399,"f":"\u7def"},{"index":400,"f":"\u7d1c"},{"index":401,"f":"\u7d14"},{"index":402,"f":"\u7d15"}]},{"info":{"index":514,"j":"\u52bf"},"list":[{"index":514,"f":"\u52e2"},{"index":515,"f":"\u532d"},{"index":516,"f":"\u55ae"},{"index":517,"f":"\u8ce3"},{"index":518,"f":"\u81e5"},{"index":519,"f":"\u5df9"}]},{"info":{"index":835,"j":"\u6800"},"list":[{"index":835,"f":"\u6894"},{"index":836,"f":"\u67f5"},{"index":837,"f":"\u6a19"},{"index":838,"f":"\u68e7"},{"index":839,"f":"\u6adb"},{"index":840,"f":"\u6af3"}]},{"info":{"index":1104,"j":"\u6656"},"list":[{"index":1104,"f":"\u6689"},{"index":1105,"f":"\u68f2"},{"index":1106,"f":"\u6a23"},{"index":1107,"f":"\u6b12"},{"index":1108,"f":"\u68ec"},{"index":1109,"f":"\u690f"}]},{"info":{"index":1168,"j":"\u7b0b"},"list":[{"index":1168,"f":"\u7b4d"},{"index":1169,"f":"\u7b46"},{"index":1170,"f":"\u7b67"},{"index":1171,"f":"\u7dca"},{"index":1172,"f":"\u7d86"},{"index":1173,"f":"\u7d83"}]},{"info":{"index":1494,"j":"\u915d"},"list":[{"index":1494,"f":"\u919e"},{"index":1495,"f":"\u92ac"},{"index":1496,"f":"\u92a0"},{"index":1497,"f":"\u927a"},{"index":1498,"f":"\u92aa"},{"index":1499,"f":"\u92cf"}]},{"info":{"index":1719,"j":"\u94ff"},"list":[{"index":1719,"f":"\u93d7"},{"index":1720,"f":"\u92b7"},{"index":1721,"f":"\u9396"},{"index":1722,"f":"\u92f0"},{"index":1723,"f":"\u92e5"},{"index":1724,"f":"\u92e4"}]},{"info":{"index":1890,"j":"\u952b"},"list":[{"index":1890,"f":"\u9307"},{"index":1891,"f":"\u931f"},{"index":1892,"f":"\u9320"},{"index":1893,"f":"\u9375"},{"index":1894,"f":"\u92f8"},{"index":1895,"f":"\u9333"}]},{"info":{"index":2252,"j":"\u9e6c"},"list":[{"index":2252,"f":"\u9df8"},{"index":2253,"f":"\u9f72"},{"index":2254,"f":"\u9f77"},{"index":2255,"f":"\u56c5"},{"index":2256,"f":"\u56c2"},{"index":2257,"f":"\u7669"}]},{"info":{"index":147,"j":"\u4f2a"},"list":[{"index":147,"f":"\u507d"},{"index":148,"f":"\u4f47"},{"index":149,"f":"\u95dc"},{"index":150,"f":"\u8208"},{"index":151,"f":"\u8ecd"},{"index":152,"f":"\u8fb2"}]},{"info":{"index":434,"j":"\u8bc8"},"list":[{"index":434,"f":"\u8a50"},{"index":435,"f":"\u8a34"},{"index":436,"f":"\u8a3a"},{"index":437,"f":"\u8a46"},{"index":438,"f":"\u8b05"},{"index":439,"f":"\u8a5e"}]},{"info":{"index":531,"j":"\u5785"},"list":[{"index":531,"f":"\u58df"},{"index":532,"f":"\u58da"},{"index":533,"f":"\u5099"},{"index":534,"f":"\u596e"},{"index":535,"f":"\u59cd"},{"index":536,"f":"\u5b78"}]},{"info":{"index":791,"j":"\u59dc"},"list":[{"index":791,"f":"\u8591"},{"index":792,"f":"\u5a41"},{"index":793,"f":"\u5a6d"},{"index":794,"f":"\u5b08"},{"index":795,"f":"\u5b0c"},{"index":796,"f":"\u5b4c"}]},{"info":{"index":971,"j":"\u8d33"},"list":[{"index":971,"f":"\u8cb0"},{"index":972,"f":"\u8cbc"},{"index":973,"f":"\u8cb4"},{"index":974,"f":"\u8cba"},{"index":975,"f":"\u8cb8"},{"index":976,"f":"\u8cbf"}]},{"info":{"index":1302,"j":"\u9a8a"},"list":[{"index":1302,"f":"\u9a6a"},{"index":1303,"f":"\u9a01"},{"index":1304,"f":"\u9a57"},{"index":1305,"f":"\u99ff"},{"index":1306,"f":"\u9d23"},{"index":1307,"f":"\u9d87"}]},{"info":{"index":1586,"j":"\u6120"},"list":[{"index":1586,"f":"\u614d"},{"index":1587,"f":"\u61a4"},{"index":1588,"f":"\u6192"},{"index":1589,"f":"\u6463"},{"index":1590,"f":"\u652c"},{"index":1591,"f":"\u64b3"}]},{"info":{"index":1729,"j":"\u9509"},"list":[{"index":1729,"f":"\u92bc"},{"index":1730,"f":"\u92dd"},{"index":1731,"f":"\u92d2"},{"index":1732,"f":"\u92c5"},{"index":1733,"f":"\u92f6"},{"index":1734,"f":"\u9426"}]},{"info":{"index":1885,"j":"\u9524"},"list":[{"index":1885,"f":"\u9318"},{"index":1886,"f":"\u9310"},{"index":1887,"f":"\u9326"},{"index":1888,"f":"\u9341"},{"index":1889,"f":"\u9308"},{"index":1890,"f":"\u9307"}]},{"info":{"index":2276,"j":"\u6512"},"list":[{"index":2276,"f":"\u6522"},{"index":2277,"f":"\u766c"},{"index":2278,"f":"\u7c5f"},{"index":2279,"f":"\u7e98"},{"index":2280,"f":"\u8b96"},{"index":2281,"f":"\u8e95"}]},{"info":{"index":88,"j":"\u5723"},"list":[{"index":88,"f":"\u8056"},{"index":89,"f":"\u8655"},{"index":90,"f":"\u982d"},{"index":91,"f":"\u5be7"},{"index":92,"f":"\u5c0d"},{"index":93,"f":"\u723e"}]},{"info":{"index":452,"j":"\u8fdb"},"list":[{"index":452,"f":"\u9032"},{"index":453,"f":"\u9060"},{"index":454,"f":"\u9055"},{"index":455,"f":"\u9023"},{"index":456,"f":"\u9072"},{"index":457,"f":"\u90f5"}]},{"info":{"index":680,"j":"\u8be4"},"list":[{"index":680,"f":"\u8acd"},{"index":681,"f":"\u8a72"},{"index":682,"f":"\u8a73"},{"index":683,"f":"\u8a6b"},{"index":684,"f":"\u8ae2"},{"index":685,"f":"\u8a61"}]},{"info":{"index":833,"j":"\u67e0"},"list":[{"index":833,"f":"\u6ab8"},{"index":834,"f":"\u6a89"},{"index":835,"f":"\u6894"},{"index":836,"f":"\u67f5"},{"index":837,"f":"\u6a19"},{"index":838,"f":"\u68e7"}]},{"info":{"index":1151,"j":"\u75b4"},"list":[{"index":1151,"f":"\u5c59"},{"index":1152,"f":"\u7670"},{"index":1153,"f":"\u75d9"},{"index":1154,"f":"\u76ba"},{"index":1155,"f":"\u76de"},{"index":1156,"f":"\u9e7d"}]},{"info":{"index":1297,"j":"\u9884"},"list":[{"index":1297,"f":"\u9810"},{"index":1298,"f":"\u9911"},{"index":1299,"f":"\u9913"},{"index":1300,"f":"\u9918"},{"index":1301,"f":"\u9912"},{"index":1302,"f":"\u9a6a"}]},{"info":{"index":1447,"j":"\u8138"},"list":[{"index":1447,"f":"\u81c9"},{"index":1448,"f":"\u826b"},{"index":1449,"f":"\u863f"},{"index":1450,"f":"\u87a2"},{"index":1451,"f":"\u71df"},{"index":1452,"f":"\u7e08"}]},{"info":{"index":1717,"j":"\u94fd"},"list":[{"index":1717,"f":"\u92f1"},{"index":1718,"f":"\u93c8"},{"index":1719,"f":"\u93d7"},{"index":1720,"f":"\u92b7"},{"index":1721,"f":"\u9396"},{"index":1722,"f":"\u92f0"}]},{"info":{"index":1899,"j":"\u9619"},"list":[{"index":1899,"f":"\u95d5"},{"index":1900,"f":"\u96db"},{"index":1901,"f":"\u9727"},{"index":1902,"f":"\u97d9"},{"index":1903,"f":"\u97de"},{"index":1904,"f":"\u97fb"}]},{"info":{"index":2180,"j":"\u9556"},"list":[{"index":2180,"f":"\u93e2"},{"index":2181,"f":"\u93dc"},{"index":2182,"f":"\u93cd"},{"index":2183,"f":"\u93de"},{"index":2184,"f":"\u93e1"},{"index":2185,"f":"\u93d1"}]},{"info":{"index":131,"j":"\u4e98"},"list":[{"index":131,"f":"\u4e99"},{"index":132,"f":"\u4e9e"},{"index":133,"f":"\u7522"},{"index":134,"f":"\u50f9"},{"index":135,"f":"\u773e"},{"index":136,"f":"\u512a"}]},{"info":{"index":236,"j":"\u8bb2"},"list":[{"index":236,"f":"\u8b1b"},{"index":237,"f":"\u8af1"},{"index":238,"f":"\u8b33"},{"index":239,"f":"\u8a4e"},{"index":240,"f":"\u8a1d"},{"index":241,"f":"\u8a25"}]},{"info":{"index":591,"j":"\u6cf7"},"list":[{"index":591,"f":"\u7027"},{"index":592,"f":"\u7018"},{"index":593,"f":"\u6ffc"},{"index":594,"f":"\u7009"},{"index":595,"f":"\u6f51"},{"index":596,"f":"\u6fa4"}]},{"info":{"index":784,"j":"\u57a9"},"list":[{"index":784,"f":"\u580a"},{"index":785,"f":"\u588a"},{"index":786,"f":"\u57e1"},{"index":787,"f":"\u584f"},{"index":788,"f":"\u5816"},{"index":789,"f":"\u8907"}]},{"info":{"index":944,"j":"\u836f"},"list":[{"index":944,"f":"\u85e5"},{"index":945,"f":"\u96d6"},{"index":946,"f":"\u8766"},{"index":947,"f":"\u8806"},{"index":948,"f":"\u8755"},{"index":949,"f":"\u87fb"}]},{"info":{"index":1292,"j":"\u987f"},"list":[{"index":1292,"f":"\u9813"},{"index":1293,"f":"\u980e"},{"index":1294,"f":"\u9812"},{"index":1295,"f":"\u980c"},{"index":1296,"f":"\u980f"},{"index":1297,"f":"\u9810"}]},{"info":{"index":1482,"j":"\u8c1d"},"list":[{"index":1482,"f":"\u8ade"},{"index":1483,"f":"\u8cd5"},{"index":1484,"f":"\u8cd1"},{"index":1485,"f":"\u8cda"},{"index":1486,"f":"\u8cd2"},{"index":1487,"f":"\u8e8d"}]},{"info":{"index":1832,"j":"\u7f1d"},"list":[{"index":1832,"f":"\u7e2b"},{"index":1833,"f":"\u7e1e"},{"index":1834,"f":"\u7e8f"},{"index":1835,"f":"\u7e2d"},{"index":1836,"f":"\u7e0a"},{"index":1837,"f":"\u7e11"}]},{"info":{"index":1885,"j":"\u9524"},"list":[{"index":1885,"f":"\u9318"},{"index":1886,"f":"\u9310"},{"index":1887,"f":"\u9326"},{"index":1888,"f":"\u9341"},{"index":1889,"f":"\u9308"},{"index":1890,"f":"\u9307"}]},{"info":{"index":2282,"j":"\u8e7f"},"list":[{"index":2282,"f":"\u8ea5"},{"index":2283,"f":"\u9454"},{"index":2284,"f":"\u9744"},{"index":2285,"f":"\u97dd"},{"index":2286,"f":"\u986b"},{"index":2287,"f":"\u9a65"}]}],
	isfirst = true,
	$d = function(id){return document.getElementById(id);};
	
//广告
LBShare.showAd({
	pos: 'top'
});

LBShare.updateData({
		title: "挑战繁体字",
		imgUrl: "http://yx.hcsat.cn/games/ftz/images/i01.jpg",
		desc: "快来看看你都认识多少个繁体字吧！"
});

function more(){
	LBShare.more();
}

//EVENTS
$("#start").on('touchend', function() {

	LBShare.showAd({
		pos: 'top',
		hide: true
	})

	$d('wbegin').style.display = "none";
	$d('main').style.display = "block";
	init();
});
$("#whshare").on('touchstart', function() {
	bshare();
	//$('#fx')[0].style.display = "block";
});
$("#fx").on('touchstart', function() {
	$('#fx')[0].style.display = "none";
});

//开始游戏
function init() {
	time();
	$('#selects_box').off('touchend');
	newpage();

	
}

//猜字倒计时
function time() {
	var time = $('.time').text();
	var second = parseInt(time);
	if (second <= 0) {
		$('#selects_box').off('touchend');
		allnumshow = allnum;
		fuzhibef();
		gameover();
	} else {
		second--;
		time = second + 's';
		$('.time span').text(second);
	}
	clock = setTimeout('time()', 1000);
}

function newpage() {
	shownum = 0;
	$('.time span').text(sec);
	$('.cell').removeClass('err');
	$('.cell').removeClass('cor');
	show = randArray(show);
	shows(shownum);
}

//随机汉子列表
function randArray(data) {
	var randomArr = [],copy = data.slice(0);
	for (var i = 0, l = copy.length; l--; i++) {
		randomArr[i] = copy.splice(Math.floor(Math.random() * l),1)[0];
	}
	return randomArr;
}

//显示当前页文字
function shows(shownum) {
	$(".font").text(show[shownum]['info']['j']);
	var p = Math.floor(Math.random() * 6);
	for (var i = 0; i < 6; i++) {
		$("#s_" + i).text(show[shownum]['list'][i]['f']);
	}
	$("#s_" + p).text(show[shownum]['list'][0]['f']);
	$("#s_0").text(show[shownum]['list'][p]['f']);
	isWin(shownum);
}

function update() {
	$('.time span').text(sec);
	$('.cell').removeClass('err');
	$('.cell').removeClass('cor');
	clock = setTimeout('time()', time_num);
	allnum++;
	if (shownum < 99) {
		shownum++;
		shows(shownum);
	} else {
		newpage();
	}
	
}

//是否答对
function isWin(shownum) {
	$("#selects_box").on('touchend', '.cell', function() {
		if ($(this).text() == show[shownum]['list'][0]['f']) {
			$(this).addClass('cor');
			$('#selects_box').off('touchend');
			clearInterval(clock);
			setTimeout("update()", time_num);
		} else {
			$(this).addClass('err');
			$('#selects_box').off('touchend');
			allnumshow = allnum;
			fuzhibef();
			setTimeout("gameover()", time_num);
		}
	});
}

function gameover() {
	clearInterval(clock);
	if (isfirst) {
		isfirst = false;
		$d('wend').style.display = "block";
		$d('main').style.display = "none";
		share();
	}
}

function fuzhibef() {
	if (allnumshow == 0) {
		$d('text_pic').className = 'lose';
		$d('text_top').innerHTML = '一个字都没认出来';
		$d('text_bottom').innerHTML = '你的小伙们都惊呆了，没文化真可怕';
	} else if (allnumshow <= 10 && allnumshow > 0) {
		$d('text_pic').className = 'lose';
		$d('text_top').innerHTML = '认识<span id="shownum">' + allnumshow + '</span>个繁体字';
		$d('text_bottom').innerHTML = '你语文是数学老师教的吧，再来一把展现你的实力';
	} else if (allnumshow > 10 && allnumshow <= 50) {
		$d('text_pic').className = 'win';
		$d('text_top').innerHTML = '认识<span id="shownum">' + allnumshow + '</span>个繁体字';
		$d('text_bottom').innerHTML = '水平杠杠的，去香港和台湾都平趟了';
	} else if (allnumshow > 50 && allnumshow <= 100) {
		$d('text_pic').className = 'win';
		$d('text_top').innerHTML = '认识<span id="shownum">' + allnumshow + '</span>个繁体字';
	} else if (allnumshow > 100) {
		$d('text_pic').className = 'win';
		$d('text_top').innerHTML = '认识<span id="shownum">' + allnumshow + '</span>个繁体字';
	}
}

function share() {
	var str = allnum>=3? '快来看看你都认识多少个繁体字吧！':'我认识'+allnum+'个繁体字，超过了'+(allnum-2)+'%的中国人，你敢来挑战吗！'
	LBShare.updateData({ //修改分享文案
		desc: str
	});
}

function bshare() {
	LBShare.callShare();
}

/* 手机屏幕高度 */
var isDesktop = navigator['userAgent'].match(/(ipad|iphone|ipod|android|windows phone)/i) ? false : true;

function scr() {
	var height = isDesktop ? 1000 : ((window.innerWidth > window.innerHeight ? window.innerWidth : window.innerHeight));
	height = height + 'px';
	$(".wrap").css('height', height);
}
setTimeout('scr()', 100);