
//存储
export const storeValue = (key: string, token: string) => {
    localStorage.setItem(key, token)
}
//获取
export const getStoredValue = (key: string): string | null => {
    const token = localStorage.getItem(key)
    return token
}
//删除
export const removeStoredValue = (key: string) => {
    localStorage.removeItem(key)
}
//清空
export const clearStoredValue = () => {
    localStorage.clear()
}