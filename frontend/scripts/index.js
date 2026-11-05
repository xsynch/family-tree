const welcomeModal = document.getElementById("welcome-modal")


window.onload = (event) => {
    setTimeout(() => welcomeModal.classList.add("welcome-modal-visible"), 2000)
};