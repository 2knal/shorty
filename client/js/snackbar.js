function showSnackbar(msg) {
    let x = document.getElementById("snackbar");
    x.className = "show";
    x.innerText = msg;
    setTimeout(function(){ x.className = x.className.replace("show", ""); }, 3000);
}