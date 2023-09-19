import { Box, Theme, useMediaQuery } from "@mui/material";
import { makeStyles } from '@mui/styles'
import logo from "../../../src/assets/images/home/logo.png"
import UserInfo from '../../components/userInfo/UserInfo'
import { Link } from "react-router-dom";
import Menu from "../Menu/Menu";
import Search from "../../components/Search/Search";
import PublishButton from "../../components/PublishButton/PublishButton";
import { getStoredValue } from "../../utils/storage";
import { useTheme } from '@mui/material/styles';
import MobileMenu from "../Menu/MobileMenu";
const useStyles = makeStyles((theme: Theme) => ({
    root: {
        width: "100%",
        boxShadow: "0 0 10px rgba(0,0,0,.1)",
        position: "fixed",
        top: "0",
        left: "0",
        backgroundColor: "#fff",
        zIndex: "1000",
        height: "65px",
        boxSizing: "border-box",
    },
    content: {
        width: "100%",
        height: "100%",
        padding: "0 20px",
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        boxSizing: "border-box",
        maxWidth: "1520px",
        margin: "0 auto",
        [`${theme.breakpoints.down('sm')}`]: {
            padding: "0 10px",
        }
    },
    logoBox: {
        display: "flex",
        alignItems: "center",
        textDecoration: "none",
        color: "black",
        marginRight: "30px",
        "& span": {
            fontSize: "20px",
            fontWeight: "bold",
            // fontFamily: "fantasy",
            [`${theme.breakpoints.down('sm')}`]: {
                fontSize: "14px",
            }
        }
    },
    menuBox: {
        display: "flex",
        alignItems: "center",
    },
    searchBox: {
        display: "flex",
        alignItems: "center",
    }
}))

const Header = () => {
    const theme = useTheme();
    const classes = useStyles()
    const isLogin = getStoredValue('access_token')
    const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
    const isSmall = useMediaQuery(theme.breakpoints.down('lg'));
    return (
        <Box className={classes.root}>
            <Box className={classes.content}>
                {!isMobile ?
                    <Box className={classes.menuBox}>
                        <Link to="/" className={classes.logoBox}>
                            <img
                                src={logo}
                                alt="logo"
                                style={{
                                    width: "60px",
                                    height: "60px",
                                    cursor: "pointer",
                                    marginRight: "10px",
                                }}
                            />
                            <span>Super Lion</span>
                        </Link>
                        <Menu />
                    </Box> :
                    <MobileMenu />
                }
                <Box className={classes.searchBox}>
                    {!isSmall && <Search />}
                    {isLogin && !isMobile && <PublishButton />}
                    <UserInfo />
                </Box>
            </Box>
        </Box>
    )
}

export default Header