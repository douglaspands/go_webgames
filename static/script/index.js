(async() => {
    const inputEmulator = document.getElementById('inputEmulator');
    const inputRom = document.getElementById('inputRom');
    const button = document.querySelector('button');
    button.disabled = true;
    inputEmulator.addEventListener('change', async (event) => {
        const response = await fetch(`/games?console=${encodeURIComponent(event.target.value)}`);
        const content = await response.json();
        inputRom.innerHTML = '<option selected disabled>Escolha um game...</option>';
        content.data.forEach(game => {
            const option = document.createElement('option');
            option.value = game;
            option.textContent = game;
            inputRom.appendChild(option);
        });
        button.disabled = false;
    });
})();