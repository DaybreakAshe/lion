import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
import { useState } from "react";
import { Editor } from '@tinymce/tinymce-react';
import ConfirmButton from "../../components/ConfirmButton/ConfirmButton";
import { uploadFile } from "../../services/createBlog/createBlog.service";
const useStyles = makeStyles((_theme: Theme) => ({
    root: {
        width: "100%",
        height: '500px',
    },
    buttonBox: {
        display: "flex",
        justifyContent: "flex-end",
        marginTop: "20px",
    }
}))
const PublicBlog = () => {
    const classes = useStyles()
    const [content, setContent] = useState<any>(null);

    const onUpload = async (fileContent: any) => {
        if (!fileContent) return;
        const formData = new FormData();
        formData.append('image', fileContent);
        // const res = await fetch('https://148.100.77.194:8999/upload', {
        //     method: 'POST',
        //     body: formData
        // });
        console.log("file##", formData);
        const res= await uploadFile({
            picture: formData,
            busiType:''
        })
        console.log("res##",res);
    }

    const saveBlog = () => {
        //1.获取字符串content里的图片
        const imgReg = /<img.*?(?:>|\/>)/gi;
        // 2.提取src
        const srcReg = /src=[\'\"]?([^\'\"]*)[\'\"]?/i;
        const imgSrc = content.match(imgReg)?.map((img: any) => {
            return img.match(srcReg)[1];
        });
        //3.提取base64
        const base64Reg = /base64,([^\'\"]*)[\'\"]?/i;
        const imgBase64 = imgSrc.map((img: any) => {
            return img.match(base64Reg)[1];
        });
        //4.对base64字符串进行解码,转换为二进制图片数据
        const imgBinary = imgBase64?.map((img: any) => {
            return atob(img);
        });
        //5.利用二进制数据生成一个File对象以包装图片数据
        const imgBlob = imgBinary?.map((img: any) => {
            const imgBuffer = new ArrayBuffer(img.length);
            const imgUint8 = new Uint8Array(imgBuffer);
            for (let i = 0; i < img.length; i++) {
                imgUint8[i] = img.charCodeAt(i);
            }
            const imgBlob = new Blob([imgUint8], { type: 'image/png' });
            return new File([imgBlob], 'img.png', { type: 'image/png' });
        });
        
        console.log("imgBlob", imgBlob)
        //6.上传图片
        imgBlob?.map((img: any) => {
            onUpload(img)
        });
        //转json，传给后端
        // const contentJson = JSON.stringify(content)
        // console.log("###",contentJson,typeof contentJson)
    }
    const cancel = () => {
        console.log("cancel")
    }
    return (
        <>
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
                        }}
                    />
                    <ConfirmButton
                        loading={false}
                        value={"Save"}
                        handleClick={saveBlog}
                        option={{
                            width: "100px",
                            height: "42px",
                            color: "#fff",
                        }}
                    />
                </Box>

            </Box>
        </>
    )
}

export default PublicBlog;