(async() => {
    const inputConsole = document.getElementById('inputConsole');
    const inputGame = document.getElementById('inputGame');
    const button = document.querySelector('button');
    button.disabled = true;
    inputConsole.addEventListener('change', async (event) => {
        const response = await fetch(`/games?console=${encodeURIComponent(event.target.value)}`);
        const content = await response.json();
        inputGame.innerHTML = '<option selected disabled>Escolha um game...</option>';
        content.data.forEach(game => {
            const option = document.createElement('option');
            option.value = game;
            option.textContent = game;
            inputGame.appendChild(option);
        });
        button.disabled = false;
    });
})();