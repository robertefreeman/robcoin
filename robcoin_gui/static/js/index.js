var items = {};
var totalCoins = 0;

window.feed = function(callback) {
  $.ajax({
    type: "GET",
    dataType: "json",
    headers: {
      Accept: "application/json",
      "Access-Control-Allow-Origin": "*"

    },
    url: "http://api.coin.robert.dtcntr.net/hashes",
    success: function(data) {
      var mem = data.hashes;
      var tick = {
        plot0: parseInt(mem)
      };
      feed2();
      callback(JSON.stringify(tick));
    }
  });
};


window.feed2 = function(callback) {
  $.ajax({
    type: "GET",
    dataType: "json",
    headers: {
      Accept: "application/json",
      "Access-Control-Allow-Origin": "*"

    },
    url: "http://api.coin.robert.dtcntr.net/coins",
    success: function(data) {
      var mem = data.coins;
      var tick = {
        plot0: parseInt(mem)
      };
      totalCoins = parseInt(mem);
      callback(JSON.stringify(tick));
    }
  });
};


/*
window.feed = function(callback) {
  $.getJSON('http://ucp.robert.dtcntr.net:30026', function(data) {
    console.log(data);
    items = data.Ops.toString();
  });

  callback(JSON.stringify(20));
};
*/

var myConfigHashes = {
  //chart styling
  type: 'line',
  globals: {
    fontFamily: 'Roboto',
  },
  backgroundColor: '#fff',
  title: {
    backgroundColor: '#1565C0',
    text: 'Hashes per second',
    color: '#fff',
    height: '30x',
  },
  plotarea: {
    marginTop: '80px'
  },
  crosshairX: {
    lineWidth: 4,
    lineStyle: 'dashed',
    lineColor: '#424242',
    marker: {
      visible: true,
      size: 9
    },
    plotLabel: {
      backgroundColor: '#fff',
      borderColor: '#e3e3e3',
      borderRadius: 5,
      padding: 15,
      fontSize: 15,
      shadow: true,
      shadowAlpha: 0.2,
      shadowBlur: 5,
      shadowDistance: 4,
    },
    scaleLabel: {
      backgroundColor: '#424242',
      padding: 5
    }
  },
  scaleY: {
    guide: {
      visible: false
    }
  },
  tooltip: {
    visible: false
  },
  //real-time feed
  refresh: {
    type: 'feed',
    transport: 'js',
    url: 'feed()',
    interval: 500
  },
  plot: {
    shadow: 1,
    shadowColor: '#eee',
    shadowDistance: '10px',
    lineWidth: 5,
    hoverState: {
      visible: false
    },
    marker: {
      visible: false
    },
    aspect: 'spline'
  },
  series: [{
    values: [],
    lineColor: '#2196F3',
    text: 'Blue Line'
  }]
};

var myConfigCoins = {
  type: "gauge",
  globals: {
    fontSize: 25
  },
  plotarea: {
    marginTop: 80
  },
  plot: {
    size: '100%',
    valueBox: {
      placement: 'center',
      text: '%v<br>Total Coins Found', //default
      fontSize: 35,
      /* rules: [{
          rule: '%v >= 90',
          text: '%v<br>'
        },
        {
          rule: '%v < 90 && %v > 50',
          text: '%v<br>'
        },
        {
          rule: '%v < 50 && %v > 30',
          text: '%v<br>'
        },
        {
          rule: '%v <  20',
          text: '%v<br>'
        }
      ] */
    }
  },
  tooltip: {
    borderRadius: 5
  },
  scaleR: {
    aperture: 270,
    minValue: 0,
    maxValue: 200,
    step: 20,
    center: {
      visible: false
    },
    tick: {
      visible: false
    },
    item: {
      offsetR: 0,
      rules: [{
        rule: '%i == 9',
        offsetX: 15
      }]
    },
    labels: ['0', '20', '40', '60', '80', '100', '120', '140', '160', '180', '200'],
    ring: {
      size: 10,
      rules: [{
          rule: '%v <= 30',
          backgroundColor: '#E53935'
        },
        {
          rule: '%v > 30 && %v < 50',
          backgroundColor: '#EF5350'
        },
        {
          rule: '%v >= 50 && %v < 90',
          backgroundColor: '#FFA726'
        },
        {
          rule: '%v >= 90',
          backgroundColor: '#29B6F6'
        }
      ]
    }
  },
  refresh: {
    type: "feed",
    transport: "js",
    url: "feed2()",
    interval: 200,
    resetTimeout: 1000
  },
  series: [{
    values: [10], // starting value
    backgroundColor: 'black',
    indicator: [10, 10, 10, 10, 0.75],
    animation: {
      effect: 2,
      method: 1,
      sequence: 4,
      speed: 20
    },
  }]
};

zingchart.render({
  id: 'myChartHashes',
  data: myConfigHashes,
  height: 400,
  width: 500
});

zingchart.render({
  id: 'myChartCoin',
  data: myConfigCoins,
  height: 400,
  width: 500
});
