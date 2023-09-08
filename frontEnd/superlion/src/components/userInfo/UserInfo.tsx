import { Box, Theme, TextField, Divider, Dialog, DialogTitle, DialogContent, DialogContentText, DialogActions, Button } from "@mui/material";
import { useState } from 'react'
import { makeStyles } from '@mui/styles'
import CloseIcon from '@mui/icons-material/Close'
import CircularProgress from '@mui/material/CircularProgress';
import google_ico from '../../assets/images/login/ico-google.svg'
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
        justifyContent:"center"
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
    loginButton:{
        background:"#FFB948 !important",
        color:"#fff !important",
        width:"88px !important",
        height:"44px !important",
        borderRadius:"24px !important"
    }
}))

const UserInfo = () => {
    const classes = useStyles()
    const [open, setOpen] = useState(false)
    const [loading,setLoading] = useState<boolean>(false)
    const youtubeLogin = async () => {
        setLoading(true)
        window.location.href =`https://accounts.google.com/o/oauth2/v2/auth?scope=https://www.googleapis.com/auth/userinfo.email&include_granted_scopes=true&response_type=token&state=3EAB37D9D5310BFE&redirect_uri=https://superlion.vercel.app&client_id=32041706814-n36purujenfckur3831hkjgipbc4plia.apps.googleusercontent.com`
    }
    return (
        <>
            <Button onClick={() => setOpen(true)} className={classes.loginButton}>
                log in
            </Button>
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