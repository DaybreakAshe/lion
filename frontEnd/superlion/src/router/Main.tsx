import { Box, Theme } from "@mui/material";
import { makeStyles } from "@mui/styles";
import { Routes, Route } from "react-router-dom";
import HomePage from "../pages/HomePage/HomePage";
import Blog from "../pages/Blog/Blog";
import About from "../pages/About/About";
import Contacts from "../pages/Contacts/Contacts";
import PublishBlog from "../pages/PublishBlog/PublishBlog";
import Profile from "src/pages/Profile/Profile";
import MyBlog from "src/pages/MyBlog/MyBlog";
import homeBack from "../../src/assets/images/home/home2.jpg";

const useStyles = makeStyles((_theme: Theme) => ({
  root: {
    width: "100%",
    backgroundColor: "#F9F9F9",
    backgroundImage: `url(${homeBack})`,
    backgroundSize: "cover",
    backgroundPosition: "center",
    backgroundRepeat: "no-repeat",
  },
  content: {
    width: "100%",
    maxWidth: "1520px",
    minHeight: "calc(100vh - 65px)",
    boxSizing: "border-box",
    margin: "0 auto",
    marginTop: "65px",
    padding: "20px",
  },
}));

const Main = () => {
  const classes = useStyles();
  return (
    <Box className={classes.root}>
      <Box className={classes.content}>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/blog" element={<Blog />} />
          <Route path="/about" element={<About />} />
          <Route path="/contacts" element={<Contacts />} />
          <Route path="/publishBlog" element={<PublishBlog />} />
          <Route path="/profile" element={<Profile />} />
          <Route path="/myBlog" element={<MyBlog />} />
          <Route path="/*" element={<></>} />
        </Routes>
      </Box>
    </Box>
  );
};

export default Main;
