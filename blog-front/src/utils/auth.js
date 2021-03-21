import {LoginKey} from '../const/login.js'

export function getToken() {
  return localStorage.getItem(LoginKey.ACCESS_TOKEN)
}

export function setToken(token, username) {
    localStorage.setItem(LoginKey.ACCESS_TOKEN, token);
}

export function removeToken() {
    localStorage.setItem(LoginKey.ACCESS_TOKEN, '');
}
