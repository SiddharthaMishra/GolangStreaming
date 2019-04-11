// function b64(e) {
//     var t = "";
//     var n = new Uint8Array(e);
//     var r = n.byteLength;
//     for (var i = 0; i < r; i++) {
//         t += String.fromCharCode(n[i])
//     }
//     return window.btoa(t)
// }
// var move;
// $(document).ready(function () {
//     let namespace = '/client';
//     let namespace = "";
//     var socket = io.connect(location.protocol + '//' + document.domain + ':' + location.port + namespace);

//     move = function (arg) {
//         socket.emit("move", {
//             "direction": arg
//         })
//     };

//     socket.on('image2Client', function (data) {
//         $("#sensordata").html("sensor value: " + data.sensor)
//         $("#img").attr("src", "data:image/png;base64," + b64(data.image));
//     });
// });

window.onload = () => {
    let div = document.getElementById("sentdata");
    let img = document.getElementById('img');


    let url = `ws://${document.domain}:${location.port}/viewer`
    let ws = new WebSocket(url);
    ws.onmessage = (event) => {
        img.src = 'data:image/png;base64,' + event.data;
        console.log(event.data);
    }
}