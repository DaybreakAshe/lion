import { configureStore, AnyAction } from '@reduxjs/toolkit';
import {
    persistReducer,
    persistStore,
    FLUSH,
    REHYDRATE,
    PAUSE,
    PERSIST,
    PURGE,
    REGISTER,
} from 'redux-persist';
import storage from 'redux-persist/lib/storage';

const initialState = {
    id: null,
    username: null,
    nickname: null,
    avatar: null,
    email: null,
}

function reducer(state = initialState, action: AnyAction) {
    switch (action.type) {
        case 'SET_ID':
            return { ...state, id: action.payload };
        case 'DELETE_ID':
            return { ...state, id: null };
        case 'SET_USERNAME':
            return { ...state, username: action.payload };
        case 'DELETE_USERNAME':
            return { ...state, username: null };
        case 'SET_NICKNAME':
            return { ...state, nickname: action.payload };
        case 'DELETE_NICKNAME':
            return { ...state, nickname: null };
        case 'SET_AVATAR':
            return { ...state, avatar: action.payload };
        case 'DELETE_AVATAR':
            return { ...state, avatar: null };
        case 'SET_EMAIL':
            return { ...state, email: action.payload };
        case 'DELETE_EMAIL':
            return { ...state, email: null };
        default:
            return state;
    }
}

const persistConfig = {
    key: 'root',
    storage,
};

const persistedReducer = persistReducer(persistConfig, reducer);
//store
const store = configureStore({
    reducer: persistedReducer,
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware({
            serializableCheck: {
                ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
            },
        }),
});
const persistor = persistStore(store);

export { store, persistor };