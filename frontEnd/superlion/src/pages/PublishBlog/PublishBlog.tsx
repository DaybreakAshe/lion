import { Box, Theme } from "@mui/material";
import { makeStyles } from "@mui/styles";
import { useState, useCallback } from "react";
import { Editor } from "@tinymce/tinymce-react";
import ConfirmButton from "src/components/button/ConfirmButton";
import {
  uploadFile,
  createBlog,
} from "../../services/createBlog/createBlog.service";
import SetBlogInfo from "./SetBlogInfo";
import { useNavigate } from "react-router-dom";
import { enqueueSnackbar } from "notistack";
const useStyles = makeStyles((_theme: Theme) => ({
  root: {
    width: "100%",
  },
  buttonBox: {
    display: "flex",
    justifyContent: "flex-end",
    marginTop: "20px",
  },
}));
const PublicBlog = () => {
  const classes = useStyles();
  const navigate = useNavigate();
  const [loading, setLoading] = useState<boolean>(false);
  const [content, setContent] = useState<any>(null);
  const [blogContent, setBlogContent] = useState<any>(null); //替换过base64的content

  //返回base64的数组
  const handleBase64 = useCallback((): string[] => {
    const imgReg = /<img.*?(?:>|\/>)/gi;
    const srcReg = /src=['"]?([^'"]*)['"]?/i;
    const imgSrc = content.match(imgReg)?.map((img: any) => {
      return img.match(srcReg)[1];
    });
    return imgSrc;
  }, [content]);

  //上传图片
  const onUpload = async (fileContent: any) => {
    if (!fileContent) return;
    setLoading(true);
    const res = await uploadFile({
      picture: fileContent,
      busiType: "test",
    });
    return res?.data?.fileUrl;
  };

  const saveBlog = async () => {
    if (!content) {
      enqueueSnackbar("Please enter content", { variant: "warning" });
      return;
    }
    const base64Reg = /base64,([^'"]*)['"]?/i;
    const imgBase64 = handleBase64()?.map((img: any) => {
      const bytes = atob(img.match(base64Reg)[1]);
      const arr = new Uint8Array(bytes.length);
      for (let i = 0; i < bytes.length; i++) {
        arr[i] = bytes.charCodeAt(i);
      }
      const blob = new Blob([arr], { type: "image/jpeg" });
      const file = new File([blob], "image.jpg");
      return file;
    });
    if (!imgBase64) {
      setLoading(false);
      setBlogContent(content);
      return;
    }
    const uploadPromise = imgBase64.map((file) => {
      return onUpload(file);
    });
    try {
      const res = await Promise.all(uploadPromise);
      let newContent = content;
      res.map((url: string, index: number) => {
        newContent = newContent.replace(handleBase64()[index], url);
        if (index === res.length - 1) {
          setLoading(false);
          update(newContent);
        }
      });
    } catch (err) {
      enqueueSnackbar("Upload failed", { variant: "error" });
    }
  };

  const cancel = () => {
    navigate(-1);
  };

  const update = async (content: string) => {
    // "MD" | "HTML";
    const param = {
      title: "test-1",
      category: "food",
      contentType: "MD",
      markdownContent: "QWE",
      htmlContent: content,
    };
    const res = await createBlog(param);
    if (res?.code === 200) {
      enqueueSnackbar("发布成功", { variant: "success" });
    } else {
      enqueueSnackbar("发布失败", { variant: "error" });
    }
  };

  return (
    <>
      <Box className={classes.root}>
        <SetBlogInfo />
        <Editor
          apiKey="mv9vikpudtaga4ks85kphmm3zmb5ydoa7vgatwchzq2ag705"
          init={{
            height: 500,
            menubar: true,
            plugins: [
              "advlist",
              "autolink",
              "lists",
              "link",
              "image",
              "charmap",
              "preview",
              "anchor",
              "searchreplace",
              "visualblocks",
              "code",
              "fullscreen",
              "insertdatetime",
              "media",
              "table",
              "code",
              "help",
              "wordcount",
            ],
            toolbar:
              "undo redo | blocks | " +
              "bold italic forecolor | alignleft aligncenter " +
              "alignright alignjustify | bullist numlist outdent indent | " +
              "removeformat | link | help",
            content_style:
              "body { font-family:Helvetica,Arial,sans-serif; font-size:14px }",
          }}
          onEditorChange={(newValue) => {
            console.log("newValue", newValue);
            setContent(newValue);
          }}
          value={content}
        />

        <Box className={classes.buttonBox}>
          <ConfirmButton
            loading={false}
            value={"取消"}
            handleClick={cancel}
            option={{
              width: "100px",
              height: "42px",
              background: "#fff",
              marginRight: "20px",
              marginBottom: "20px",
            }}
          />
          <ConfirmButton
            loading={false}
            value={"保存到草稿"}
            handleClick={saveBlog}
            option={{
              width: "100px",
              height: "42px",
              color: "#fff",
              marginRight: "20px",
              marginBottom: "20px",
              isDisabled: true,
            }}
          />
          <ConfirmButton
            loading={loading}
            value={"发布"}
            handleClick={saveBlog}
            option={{
              width: "100px",
              height: "42px",
              color: "#fff",
              marginBottom: "20px",
            }}
          />
        </Box>
      </Box>
    </>
  );
};

export default PublicBlog;
