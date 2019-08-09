const USER_KEY = '__user__';
const TOKEN_KEY = '__token__';
const USER_TAG = '__tag__';


export function saveUser(user) {
    localStorage.setItem(USER_KEY, JSON.stringify(user))
    return user
}

export function removeUser() {
    localStorage.removeItem(USER_KEY)
    localStorage.removeItem(TOKEN_KEY)
}

export function loadUser(def) {
    let user = localStorage.getItem(USER_KEY)
    if (!user) {
        return def
    }
    if (user === "undefined") {
        return def
    }
    return JSON.parse(user)
}

export function saveToken(token) {
    localStorage.setItem(TOKEN_KEY, token)
    return token
}

export function loadToken(def) {
    let token = localStorage.getItem(TOKEN_KEY)
    if (!token) {
        return def
    }
    return token
}

export function saveUserTag(tags) {
    localStorage.setItem(USER_TAG, JSON.stringify(tags))
    return tags
}

export function loadUserTag(def) {
    let tags = localStorage.getItem(USER_TAG)
    if (!tags) {
        return def
    }
    if (tags === "undefined") {
        return def
    }
    return JSON.parse(tags)
}