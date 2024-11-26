const signInBtn = document.getElementById("signIn");
const signUpBtn = document.getElementById("signUp");
const fistForm = document.getElementById("form1");
const secondForm = document.getElementById("form2");
const btnLogin = document.getElementById("btn_login")
const container = document.querySelector(".container");

signInBtn.addEventListener("click", () => {
    container.classList.remove("right-panel-active");
});

signUpBtn.addEventListener("click", () => {
    container.classList.add("right-panel-active");
});

btnLogin.addEventListener("click",login)

function login() {
    let accountAll = [{
        email:"xintian339@gmail.com",
        pwd:339,
    }]
    let email = document.getElementById("email_login").value
    let pwd = document.getElementById("pwd_login").value

    let account = accountAll.filter(function (e) {
        return e.email == email
    })[0];

    if (!account) {
        alert('账号不存在');
    }else {
        if (email == account.email  && pwd == account.pwd) {
            alert('登陆成功');
            window.location.href="/api/zh_blog/"
        }else {
            alert('密码错误')
        }
    }


}
fistForm.addEventListener("submit", (e) => e.preventDefault());
secondForm.addEventListener("submit", (e) => e.preventDefault());

