// function b64(e) {
//     var t = "";
//     var n = new Uint8Array(e);
//     var r = n.byteLength;
//     for (var i = 0; i < r; i++) {
//         t += String.fromCharCode(n[i])
//     }
//     return window.btoa(t)
// }

window.onload = () => {

    /* connect to a websocket */
    let url = `ws://${document.domain}:${location.port}/viewer`
    let ws = new WebSocket(url);

    /* allow buttons to send control signals */
    let buttons = document.getElementsByClassName("direction");
    console.log(buttons);
    for (let button of buttons) {
        console.log(button);
        button.onclick = () => {
            console.log(button);
            ws.send(JSON.stringify({
                "direction": button.dataset.direction,
            }))
        }
    }

    /* change images when a new one is recieved */
    let div = document.getElementById("sentdata");
    let img = document.getElementById('img');
    ws.onmessage = (event) => {
        img.src = 'data:image/jpg;base64,' + event.data;
    }
}