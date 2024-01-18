// up.js

// const CHUNK_SiZE = 1024 * 1024 * 5; // 5MB
const CHUNK_SiZE = 1024 * 1024 * 1; // 1MB

const UPLOAD_URL = "/upload";

/**
 * @type {HTMLInputElement}
 */
const fileInput = document.getElementById("fileInput");

const fileInputBtn = document.getElementById("fileInputBtn");
const fileInputLabel = document.getElementById("fileInputLabel");
// const fileSubmit = document.getElementById("fileSubmit");
const fileCheckSum = document.getElementById("fileCheckSum");
const fileProgress = document.getElementById("progress");
const fileProgressSend = document.getElementById("progress-send");

let isUploading = false;

fileInputBtn.onclick = () => {
    fileInput.click();
}

fileInput.addEventListener("change", () => {
    if (fileInput.files.length > 0) {
        const nm = fileInput.files.length > 1 ? `${fileInput.files.length} files` : `${fileInput.files[0].name}`;
        fileInputLabel.innerText = "Selected: " + nm;
    }
});

/** current working direcotry */
const cwd = getCWD();

document.getElementById("cwd").innerHTML = `<a href="${cwd}">${cwd}</a>`;

/** gets the current working directory and encodes it */
function getCWD() {
    const qr = new URLSearchParams(window.location.search);
    const cwd = qr.get("cwd");

    if (!cwd) {
        cwd = "/browse/"
    }

    return decodeURIComponent(cwd);
}

async function uploadFiles() {
    isUploading = true;
    fileProgressSend.innerText = `send 0/${fileInput.files.length}`
    for (let i = 0; i < fileInput.files.length; i++) {
        let file = fileInput.files[i];
        const uuid = generateUUID();
        try {
            await upload(file, uuid);
        } catch {
            fileProgress.innerText = `current: could not upload file${fileInput.files.length > 1 ? "s" : ""}`;
            isUploading = false;
            return;
        }
        fileProgressSend.innerText = `send ${i + 1}/${fileInput.files.length}`;
        console.log("done | see the abobe err", file.name);
    }
    fileProgress.innerHTML = `current: done uploading files.<br>GO the page to see the contents <button onclick="window.location.href = cwd">Go To directoy</button>`;
    isUploading = false;
}

/**
 * uploads a sinlge file to the "/upload?cwd=fo" aka fo directory directory
 * 
 * @param {File} file 
 * @param {string} uuid 
 * @returns {Promise<void>}
 */
async function upload(file, uuid) {
    try {
        await fetch(`${UPLOAD_URL}?cwd=${encodeURIComponent(cwd)}&name=${encodeURIComponent(file.name)}&uuid=${uuid}&size=${file.size}&offset=0`, {
            method: 'POST',
        })
    } catch (err) {
        console.error("while uploading file", file.name, err)
        fileProgress.innerText = err;
        throw err;
        return
    }

    const chuckCount = Math.ceil(file.size / CHUNK_SiZE);
    for (let chuckId = 0; chuckId < chuckCount; chuckId++) {
        // calculation mistakes
        const offset = chuckId * CHUNK_SiZE;
        const readUntil = (chuckId * CHUNK_SiZE) + CHUNK_SiZE;
        const data = file.slice(offset, readUntil);

        // it's js so ... or skill issue
        if (data.size === 0) break;

        try {
            await fetch(`${UPLOAD_URL}?cwd=${encodeURIComponent(cwd)}&name=${encodeURIComponent(file.name)}&uuid=${uuid}&size=${file.size}&offset=${offset}`, {
                method: 'PATCH',
                body: data,
            });

        } catch (err) {
            console.error("while uploading file", file.name, err);
            fileProgress.innerText = err;
            throw err;
        }

        const msg = `curret: ${Math.round((chuckId / chuckCount) * 100)}% ${file.name}`;
        fileProgress.innerHTML = msg;
        console.log(msg);
    }

    // last request
    // rename the file
    try {
        let sum = "";
        if (fileCheckSum.checked) {
            console.log("calcualing checksum of:", file.name);
            fileProgress.innerHTML = `current: cheching 256sum ${file.name}`;
            sum = await calculateHashofFile(file);
        }
        await fetch(`${UPLOAD_URL}?cwd=${encodeURIComponent(cwd)}&name=${encodeURIComponent(file.name)}&uuid=${uuid}&size=${file.size}&offset=${file.size}&sha256=${sum}`, {
            method: 'PUT',
        })
    } catch (err) {
        console.error("while uploading file", file.name, err)
        fileProgress.innerText = err;
        throw err;
    }
}

/////////////////////////////////////////
// HELPER FUNCTIONS
/////////////////////////////////////////

/** Generate a random UUID */
function generateUUID() {
    const cryptoObj = window.crypto || window.msCrypto; // For cross-browser compatibility
    if (cryptoObj && cryptoObj.getRandomValues) {
        // Use a cryptographically strong random number generator if available
        const buffer = new Uint16Array(8);
        cryptoObj.getRandomValues(buffer);
        return `${toHex(buffer[0], 4)
            }${toHex(buffer[1], 4)}${toHex(buffer[2], 4)}${toHex(buffer[3], 4)}${toHex(buffer[4], 4)}${toHex(buffer[5], 4)}${toHex(buffer[6], 4)}${toHex(buffer[7], 4)} `;
        // }
    } else {
        // Fallback to a less secure method
        return 'xxxxxxxxxxxx4xxxyxxxxxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
            const r = Math.random() * 16 | 0;
            const v = c === 'x' ? r : (r & 0x3 | 0x8);
            return v.toString(16);
        });
    }
}

function toHex(value, width) {
    const hex = value.toString(16);
    return '0'.repeat(width - hex.length) + hex;
}
