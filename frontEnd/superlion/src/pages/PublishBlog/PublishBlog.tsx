import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
import { useState, useCallback } from "react";
import { Editor } from '@tinymce/tinymce-react';
import ConfirmButton from "../../components/confirmButton/ConfirmButton.tsx";
import { uploadFile, createBlog } from "../../services/createBlog/createBlog.service";
import SnackbarMessage from '../../components/Snackbar/Snackbar.tsx'
import SetBlogInfo from "./SetBlogInfo";
import { useNavigate } from "react-router-dom";
const useStyles = makeStyles((_theme: Theme) => ({
    root: {
        width: "100%",
    },
    buttonBox: {
        display: "flex",
        justifyContent: "flex-end",
        marginTop: "20px",
    }
}))
const PublicBlog = () => {
    const classes = useStyles()
    const navigate = useNavigate();
    const [alertMessage, setAlertMessage] = useState('');
    const [isOpen, setIsOpen] = useState<boolean>(false);
    const [severity, setSeverity] = useState<'error' | 'warning' | 'info' | 'success'>('info');
    const [loading, setLoading] = useState<boolean>(false);
    const [content, setContent] = useState<any>(null);
    const [blogContent, setBlogContent] = useState<any>(null);  //替换过base64的content

    //返回base64的数组
    const handleBase64 = useCallback((): string[] => {
        const imgReg = /<img.*?(?:>|\/>)/gi;
        const srcReg = /src=['"]?([^'"]*)['"]?/i;
        const imgSrc = content.match(imgReg)?.map((img: any) => {
            return img.match(srcReg)[1];
        });
        return imgSrc;
    }, [content])

    //上传图片
    const onUpload = async (fileContent: any) => {
        if (!fileContent) return;
        setLoading(true);
        const res = await uploadFile({
            picture: fileContent,
            busiType: 'test'
        })
        return res?.data?.fileUrl;
    }

    const saveBlog = async () => {
        if (!content) {
            setAlertMessage('Please enter content');
            setSeverity('error');
            setIsOpen(!isOpen);
            return;
        }
        const base64Reg = /base64,([^'"]*)['"]?/i;
        const imgBase64 = handleBase64()?.map((img: any) => {
            const bytes = atob(img.match(base64Reg)[1]);
            const arr = new Uint8Array(bytes.length);
            for (let i = 0; i < bytes.length; i++) {
                arr[i] = bytes.charCodeAt(i);
            }
            const blob = new Blob([arr], { type: 'image/jpeg' });
            const file = new File([blob], 'image.jpg');
            return file;
        });
        if (!imgBase64) {
            setLoading(false);
            setBlogContent(content);
            return;
        }
        const uploadPromise = imgBase64.map(file => {
            return onUpload(file);
        });
        try {
            const res = await Promise.all(uploadPromise);
            let newContent = content;
            res.map((url: string, index: number) => {
                newContent = newContent.replace(handleBase64()[index], url);
                if (index === res.length - 1) {
                    console.log(newContent)
                    setLoading(false);
                    setBlogContent(newContent)
                }
            })
        } catch (err) {
            setAlertMessage('Upload failed');
            setSeverity('error');
            setIsOpen(true);
        }
    }
    const cancel = () => {
        navigate(-1);
    }
    return (
        <>
            <SnackbarMessage message={alertMessage} severity={severity} duration={5000} isOpen={isOpen} />
            <Box className={classes.root}>
                <Editor
                    apiKey='mv9vikpudtaga4ks85kphmm3zmb5ydoa7vgatwchzq2ag705'
                    init={{
                        height: 500,
                        menubar: true,
                        plugins: [
                            'advlist', 'autolink', 'lists', 'link', 'image', 'charmap', 'preview',
                            'anchor', 'searchreplace', 'visualblocks', 'code', 'fullscreen',
                            'insertdatetime', 'media', 'table', 'code', 'help', 'wordcount',
                        ],
                        toolbar: 'undo redo | blocks | ' +
                            'bold italic forecolor | alignleft aligncenter ' +
                            'alignright alignjustify | bullist numlist outdent indent | ' +
                            'removeformat | link | help',
                        content_style: 'body { font-family:Helvetica,Arial,sans-serif; font-size:14px }'
                    }}
                    onEditorChange={(newValue) => {
                        setContent(newValue);
                    }}
                    value={content}
                />
                <SetBlogInfo />
                <Box className={classes.buttonBox}>
                    <ConfirmButton
                        loading={false}
                        value={"Cancel"}
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
                        loading={loading}
                        value={"Save"}
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
    )
}

export default PublicBlog;