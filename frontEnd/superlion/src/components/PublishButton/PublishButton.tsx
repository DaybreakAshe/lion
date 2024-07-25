import { Box } from "@mui/material";
import { makeStyles } from "@mui/styles";
import { Link } from "react-router-dom";
const useStyles = makeStyles(() => ({
  root: {
    position: "relative",
  },
  publishButton: {
    textDecoration: "none",
    background: "#1a73e8",
    color: "#fff",
    width: "85px",
    height: "36px",
    borderRadius: "8px",
    marginRight: "30px",
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    fontSize: "14px",
    cursor: "pointer",
  },
  publishCard: {
    width: "130px",
    position: "absolute",
    right: "20px",
    background: "#fff",
    top: "40px",
    borderRadius: "8px",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    boxShadow: "0px 0px 10px rgba(0,0,0,0.1)",
    transition: "all 0.3s ease",
    overflow: "hidden",
  },
}));

const PublishButton = () => {
  const classes = useStyles();
  return (
    <>
      <Box className={classes.root}>
        <Box
          className={classes.publishButton}
          component={Link}
          to="/publishBlog"
        >
          发布博客
        </Box>
      </Box>
    </>
  );
};

export default PublishButton;
