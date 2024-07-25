import { Box, Theme } from "@mui/material";
import { makeStyles } from "@mui/styles";
import MenuIcon from "@mui/icons-material/Menu";
import { useState, useEffect } from "react";
import MenuItem from "./MenuItem.tsx";
import { Link } from "react-router-dom";
import logo from "../../../src/assets/images/home/logo.png";
import { getStoredValue } from "../../utils/storage.ts";
const useStyles = makeStyles((theme: Theme) => ({
  root: {
    position: "relative",
  },
  menuIcon: {},
  menuCard: {
    width: "200px",
    height: "100vh",
    zIndex: 100,
    boxShadow: "0 0 3px rgba(0,0,0,.2)",
    backgroundColor: "#fff",
    position: "absolute",
    top: "-19px",
    transition: "all .3s",
    padding: "30px 15px",
    boxSizing: "border-box",
    display: "flex",
    flexDirection: "column",
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
      [`${theme.breakpoints.down("sm")}`]: {
        fontSize: "15px",
      },
    },
  },
  menuListBox: {
    display: "flex",
    flexDirection: "column",
    marginTop: "30px",
    alignItems: "flex-start",
    padding: "0 10px",
  },
}));

const MobileMenu = () => {
  const classes = useStyles();
  const isLogin = getStoredValue("access_token");
  const [open, setOpen] = useState(false);
  useEffect(() => {
    document.addEventListener("click", () => {
      setOpen(false);
    });
  }, []);
  return (
    <Box className={classes.root}>
      <MenuIcon
        className={classes.menuIcon}
        onClick={(e: any) => {
          setOpen(true);
          e.stopPropagation();
        }}
      />
      <Box
        className={classes.menuCard}
        style={{ left: open ? "-10px" : "-215px" }}
      >
        <Link to="/" className={classes.logoBox}>
          <img
            src={logo}
            alt="logo"
            style={{
              width: "40px",
              height: "40px",
              cursor: "pointer",
              marginRight: "10px",
            }}
          />
          <span>Super Lion</span>
        </Link>
        <Box className={classes.menuListBox}>
          <MenuItem title="博客" path="/blog" isMobile={true} />
          {isLogin && (
            <MenuItem title="发布博客" path="/publishBlog" isMobile />
          )}
          {isLogin && <MenuItem title="我的博客" path="/myBlog" isMobile />}
          {isLogin && <MenuItem title="我的草稿" path="/myDraft" isMobile />}
          <MenuItem title="关于" path="/about" isMobile />
        </Box>
      </Box>
    </Box>
  );
};

export default MobileMenu;
