const USER_KEY = '__user_lock__';
const TOKEN_KEY = '__token_lock__';


export function saveUser(user) {
    sessionStorage.setItem(USER_KEY, JSON.stringify(user))
    return user
}

export function removeUser() {
    sessionStorage.removeItem(USER_KEY)
    sessionStorage.removeItem(TOKEN_KEY)
}

export function loadUser(def) {
    let user = sessionStorage.getItem(USER_KEY)
    if (!user) {
        return def
    }
    if (user === "undefined") {
        return def
    }
    return JSON.parse(user)
}

export function saveToken(token) {
    sessionStorage.setItem(TOKEN_KEY, token)
    return token
}

export function loadToken(def) {
    let token = sessionStorage.getItem(TOKEN_KEY)
    if (!token) {
        return def
    }
    return token
}