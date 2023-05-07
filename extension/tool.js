document.addEventListener('DOMContentLoaded', function () {
    document.querySelector('button').addEventListener('click', () => {
        let id = parseInt(document.getElementById("trainerId").value);
        if (id > 65535 || id < 0) {
            alert("Trainer ID out of range")
        } else {
            chrome.storage.local.set({ ckkeytrainerId: id });
        }
    });

    chrome.storage.local.get(["ckkeytrainerId"]).then((result) => {
        document.querySelector('input').value = result.ckkeytrainerId
    });
});