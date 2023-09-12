import { createAction } from "@reduxjs/toolkit";

export const setId = createAction<string>("SET_ID");
export const deleteId = createAction<string>("DELETE_ID");

export const setUserName = createAction<string>("SET_USERNAME");
export const deleteUserName = createAction<string>("DELETE_USERNAME");

export const setNickName = createAction<string>("SET_NICKNAME");
export const deleteNickName = createAction<string>("DELETE_NICKNAME");

export const setAvatar = createAction<string>("SET_AVATAR");
export const deleteAvatar = createAction<string>("DELETE_AVATAR");

export const setEmail = createAction<string>("SET_EMAIL");
export const deleteEmail = createAction<string>("DELETE_EMAIL");