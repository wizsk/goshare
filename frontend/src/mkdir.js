// mkdir.js

/** it is dependent on the 'const cwd' which is in the 'up.js' file so import this after importing up.js */
async function mkdir() {
    let input = prompt("Enter the name for the new directory");
    if (!input) {
        alert("No name provided");
        return;
    }

    const url = `/mkdir?cwd=${encodeURIComponent(cwd)}&name=${encodeURIComponent(input)}`;
    console.log("req", url);
    const res = await fetch(url, {
        method: 'POST',
    });

    if (res.ok && res.redirected) {
        window.location.href = `/upload?cwd=${encodeURIComponent(cwd[cwd.length - 1] === "/" ? cwd : cwd + "/" + input)}`;
        return;
    }
    alert("Please input a valid fileName");
    mkdir();
}