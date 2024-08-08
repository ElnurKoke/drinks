function getParameterByName(name) {
    const url = new URL(window.location.href);
    return url.searchParams.get(name);
}
const id = getParameterByName('id');
console.log(id);
document.getElementById('container').innerHTML = generateHTMLbyID(id); 

function generateHTMLbyID(id){
    return `
        <p>ID напитка:${id}</p>
        <h3>Name: нет скороо</h3>
        <p>Состав: пусто скороо</p>
        <ul>
            <li>1</li>
            <li>2</li>
            <li>3</li>
            <li>4</li>
            <li>soon</li>
        </ul>
        <p>Картинка: пусто</p>
        <p>Приготовление: нет</p>
    `
}