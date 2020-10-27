const tokenName = "challenge/token";

export const createCookie = (cookieName, cookieValue, hourToExpire) => {
    const date = new Date();
    date.setTime(date.getTime() + hourToExpire * 60 * 60 * 1000);
    document.cookie = `${cookieName} = ${cookieValue}; expires = ${date.toGMTString()}`;
};

export const deleteCookie = (cookieName) => {
    document.cookie = cookieName + "=; expires=Thu, 01 Jan 1970 00:00:01 GMT;";
};

export const getCookie = (cookieName) => {
    const name = cookieName + "=";
    const decodedCookie = decodeURIComponent(document.cookie);
    const cookies = decodedCookie.split(";");
    for (let i = 0; i < cookies.length; i++) {
        let cookie = cookies[i];
        while (cookie.charAt(0) === " ") {
            cookie = cookie.substring(1);
        }
        if (cookie.indexOf(name) === 0) {
            return cookie.substring(name.length, cookie.length);
        }
    }
    return "";
};

export const getToken = () => getCookie(tokenName);

export const setToken = (token) => createCookie(tokenName, token, 0.5);

export const deleteToken = () => deleteCookie(tokenName);

export const isAuth = () => !!getToken();
