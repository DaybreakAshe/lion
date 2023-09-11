import { Box, Theme, Divider, Dialog, DialogTitle, DialogContent, Button, Avatar } from "@mui/material";
import { useState, useEffect, useCallback } from 'react'
import { makeStyles } from '@mui/styles'
import CloseIcon from '@mui/icons-material/Close'
import CircularProgress from '@mui/material/CircularProgress';
import google_ico from '../../assets/images/login/ico-google.svg'
import { getStoredValue, storeValue } from '../../utils/storage'
import { useSelector, useDispatch } from 'react-redux'
import PermIdentityIcon from '@mui/icons-material/PermIdentity';
import { getUserInfo } from '../../services/login/login.service'
import { setUsername, setNickname, setAvatar, setEmail } from '../../store/actions/actions'
const useStyles = makeStyles((theme: Theme) => ({
    dialogContent: {
        padding: "25px 10px",
        boxSizing: "border-box",
        [theme.breakpoints.down('sm')]: {
            width: "100%",
            height: "100%",
            padding: "20px 0",
        },
        display: "flex",
        flexDirection: "column",
        justifyContent: "center"
    },
    loginWay: {
        width: "450px",
        height: "44px",
        border: "1px solid #555555",
        margin: "8px 0",
        borderRadius: "24px",
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
        padding: "0 24px",
        cursor: "pointer",
        background: "#fff",
        boxSizing: "border-box",
        [theme.breakpoints.down('md')]: {
            width: "100%",
            maxWidth: "450px",
        },
        "&:hover": {
            color: "black",
            background: "#eeeeee",
        },
        "&:active": {
            color: "black",
            background: "#dadada",
        },
    },
    loginWayText: {
        fontWeight: 400,
        fontSize: '16px',
        color: '#555555',
    },
    loginButton: {
        background: "#FFB948 !important",
        color: "#fff !important",
        width: "88px !important",
        height: "44px !important",
        borderRadius: "24px !important"
    },
    hidden: {
        display: "none",
    },
    content: {
        width: "40px",
        height: "40px",
        background: "#F2F4F8",
        borderRadius: "50%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        position: "relative",
        cursor: "pointer",
    },
    userInfoCard: {
        width: "280px",
        background: "#FFFFFF",
        border: "1px solid #EEEEEE",
        boxShadow: "0px 4px 20px rgba(0, 0, 0, 0.1)",
        borderRadius: "8px",
        position: "absolute",
        top: "55px",
        right: "-10px",
        zIndex: 9999,
        [theme.breakpoints.down('md')]: {
        }
    },
    avatarBox: {
        height: "168px",
        background: "#FBEEEC",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        flexDirection: "column",
    },
    infoBox: {
        width: "100%",
        height: "56px",
        display: "flex",
        alignItems: "center",
        boxSizing: "border-box",
        paddingLeft: "24px",
        "&:hover": {
            background: "#F2F4F8"
        }
    },
    infoBoxIco: {
        width: "18px",
        height: "20px",
        marginRight: "12px"
    },
    infoBoxText: {
        fontWeight: 900,
        fontSize: '16px',
        color: '#7D849B'
    },
    nameText: {
        marginTop: "12px",
        fontWeight: 800,
        fontSize: '16px',
        color: '#000000'
    }
}))

const UserInfo = () => {
    const classes = useStyles()
    const dispatch = useDispatch();
    const [open, setOpen] = useState(false)
    const isLogin = getStoredValue('access_token')
    const [loading, setLoading] = useState<boolean>(false)
    /* eslint-disable no-restricted-globals */
    let url = location.href;
    const avatar = useSelector((state: any) => state.avatar)
    const userName = useSelector((state: any) => state.username)
    const [isShow, setIsShow] = useState<boolean>(false)
    const pathSegments = window.location.origin
    //获取地址栏url
    const redirectUri = pathSegments === 'http://localhost:3000' ? 'http://localhost:3000' : 'https://superlion.vercel.app'
    const youtubeLogin = async () => {
        setLoading(true)
        window.location.href = `https://accounts.google.com/o/oauth2/v2/auth?scope=https://www.googleapis.com/auth/userinfo.email&include_granted_scopes=true&response_type=token&state=3EAB37D9D5310BFE&redirect_uri=${redirectUri}&client_id=32041706814-n36purujenfckur3831hkjgipbc4plia.apps.googleusercontent.com`
    }
    const parseUrl = (url: string) => {
        if (url.indexOf("#") > -1) {
            url = url.split("#")[1];
        }
        const arr = url.split("&");
        const params: any = {};
        for (let i = 0; i < arr.length; i++) {
            const data = arr[i].split("=");
            params[data[0]] = data[1];
        }
        return params;
    }
    const handleUserInfo = async (info: any) => {
        console.log("传的参数##", info)
        const res = await getUserInfo(info)
        console.log("结果##", res)
        if (res) {
            storeValue('access_token', res?.access_token || '')
        }
    }
    useEffect(() => {
        if (url && !isLogin) {
            const res = parseUrl(url)
            if (res.access_token) {
                const param = {
                    accessToken: res.access_token
                }
                handleUserInfo(param)
            }
        }
    }, [isLogin, url])
    return (
        <>
            {
                isLogin ?
                    <Box className={classes.content}>
                        <Avatar src={avatar}
                            onClick={(e) => {
                                e.stopPropagation();
                                e.nativeEvent.stopImmediatePropagation();
                                setIsShow(!isShow)
                            }}
                        />
                        <Box
                            className={classes.userInfoCard}
                            style={{
                                display: isShow ? "block" : "none"
                            }}
                        >
                            <Box className={classes.avatarBox} onClick={(e) => {
                                e.stopPropagation();
                                e.nativeEvent.stopImmediatePropagation();
                            }}>
                                <Avatar src={avatar} style={{ width: "72px", height: "72px" }} />
                                <span className={classes.nameText}>{userName}</span>
                            </Box>
                            <Box>
                                <Box
                                    className={classes.infoBox}
                                    onClick={() => { }}
                                >
                                    <PermIdentityIcon className={classes.infoBoxIco} />
                                    <span className={classes.infoBoxText}>My Profile</span>
                                </Box>
                            </Box>
                        </Box>
                    </Box>
                    :
                    <Button onClick={() => setOpen(true)} className={classes.loginButton}>
                        log in
                    </Button>
            }

            <Dialog open={open} onClose={() => setOpen(false)}>
                <DialogTitle style={{
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "space-between",
                }}>
                    <span style={{
                        fontWeight: "800",
                        fontSize: "17px",
                    }}>log in</span>
                    <CloseIcon onClick={() => setOpen(false)} style={{ cursor: "pointer" }} />
                </DialogTitle>
                <DialogContent>
                    <Divider />
                    <Box className={classes.dialogContent}>
                        <Box className={classes.loginWay} onClick={() => youtubeLogin()}>
                            <img src={google_ico} alt="google_ico" />
                            {loading ?
                                <CircularProgress style={{ width: "20px", height: "20px", color: "#FF7161" }} /> : <span className={classes.loginWayText}>Continue with Google</span>}
                            <div></div>
                        </Box>
                    </Box>
                </DialogContent>
            </Dialog>
        </>
    )
}

export default UserInfo