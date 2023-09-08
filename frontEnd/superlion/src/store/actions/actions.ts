import { createAction } from "@reduxjs/toolkit";


export const setUsername = createAction<string>("SET_USERNAME");
export const deleteName = createAction<string>("DELETE_USERNAME");

export const setNickname = createAction<string>("SET_NICKNAME");
export const deleteNickname = createAction<string>("DELETE_NICKNAME");

export const setAvatar = createAction<string>("SET_AVATAR");
export const deleteAvatar = createAction<string>("DELETE_AVATAR");

export const setEmail = createAction<string>("SET_EMAIL");
export const deleteEmail = createAction<string>("DELETE_EMAIL");