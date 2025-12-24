// from chatGPT


function shaOk() {
  return crypto && crypto.subtle && crypto.subtle.digest;
}

/**
 *
 * @param {ArrayBuffer} buffer
 */
async function sha256(buffer) {
    // Calculate the hash using the Web Crypto API
    const hashBuffer = await crypto.subtle.digest('SHA-256', buffer);

    // Convert the hash buffer to a hex string
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    return hashArray.map(byte => byte.toString(16).padStart(2, '0')).join('');
}

/**
 *
 * @param {File} file
 */
function readFileAsArrayBuffer(file) {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();

        reader.onload = () => {
            resolve(reader.result);
        };

        reader.onerror = reject;

        reader.readAsArrayBuffer(file);
    });
}

/**
 * calculates checksum of the given File
 *
 * @param {File} file
 */
async function calculateHashofFile(file) {
    if (!file) {
        console.warn('No file selected.');
        return "";
    }

    try {
        const fileBuffer = await readFileAsArrayBuffer(file);
        const hash = await sha256(fileBuffer);
        return hash;
    } catch (error) {
        console.error('Error calculating hash:', error);
    }

    return "";
}
