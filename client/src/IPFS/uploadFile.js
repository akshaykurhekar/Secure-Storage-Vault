import { useState } from "react";

function UploadFile({ setFileResponse, fileName }) {
  const [selectedFile, setSelectedFile] = useState();
  const changeHandler = (event) => {
    setSelectedFile(event.target.files);
  };

  const handleSubmission = async () => {
    const token = process.env.JWT_Token;
    console.log("token:", token);

    try {
      const formData = new FormData();
      Array.from(selectedFile).forEach((file) => {
        formData.append("file", file);
      });
      const metadata = JSON.stringify({
        name: fileName,
      });
      formData.append("pinataMetadata", metadata);

      const options = JSON.stringify({
        cidVersion: 0,
      });
      formData.append("pinataOptions", options);

      const res = await fetch(
        "https://api.pinata.cloud/pinning/pinFileToIPFS",
        {
          method: "POST",
          headers: {
            Authorization: `Bearer ${token}`,
          },
          body: formData,
        }
      );
      const resData = await res.json();
      setFileResponse(resData);
      console.log("Success", resData);
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div style={{ margin: "10px" }}>
      <label className="form-label"> Choose File : </label>
      <input type="file" onChange={changeHandler} accept=".png, .jpg, .mp3" />
      <button onClick={handleSubmission}>Submit</button>
    </div>
  );
}

export default UploadFile;
