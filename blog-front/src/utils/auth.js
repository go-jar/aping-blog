import {LoginKey} from '../const/login.js'

export function GetToken() {
  return localStorage.getItem(LoginKey.ACCESS_TOKEN)
}

export function SetToken(token, username) {
    localStorage.setItem(LoginKey.ACCESS_TOKEN, token);
}

export function RemoveToken() {
    localStorage.setItem(LoginKey.ACCESS_TOKEN, '');
}
