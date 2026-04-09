const responseEl = document.getElementById('response');

function getBaseUrl() {
    return window.location.origin;
}

function showResponse(data) {
    responseEl.textContent = JSON.stringify(data, null, 2);
}

async function callApi(endpoint, method = 'GET', body = null) {
    const idInstance = document.getElementById('idInstance').value.trim();
    const apiToken = document.getElementById('apiToken').value.trim();

    if (!idInstance || !apiToken) {
        alert('Пожалуйста, заполните idInstance и ApiTokenInstance');
        return;
    }
    
    let url = method === 'GET'
	? `${getBaseUrl()}/api${endpoint}?idInstance=${idInstance}&apiTokenInstance=${apiToken}`
	: `${getBaseUrl()}/api${endpoint}`;


    
    const options = {
        method: method,
        headers: { 'Content-Type': 'application/json' }
    };

    if (body) {
        options.body = JSON.stringify({
            idInstance,
            apiTokenInstance: apiToken,
            ...body
        });
    }

    try {
        const res = await fetch(url, options);
        const data = await res.json();
        showResponse(data);
    } catch (err) {
        showResponse({ error: err.message });
    }
}

function getSettings() {
    callApi('/settings');
}

function getStateInstance() {
    callApi('/state');
}

function sendMessage() {
    const chatId = document.getElementById('chatIdMsg').value.trim();
    const message = document.getElementById('messageText').value.trim();

    if (!chatId || !message) {
        alert('Заполните chatId и текст сообщения');
        return;
    }

    callApi('/send-message', 'POST', { chatId, message });
}

function sendFileByUrl() {
    const chatId = document.getElementById('chatIdFile').value.trim();
    const urlFile = document.getElementById('urlFile').value.trim();
    const fileName = document.getElementById('fileName').value.trim();
    const caption = document.getElementById('caption').value.trim();

    if (!chatId || !urlFile || !fileName) {
        alert('Заполните chatId, URL файла и имя файла');
        return;
    }

    callApi('/send-file', 'POST', { chatId, urlFile, fileName, caption });
}
