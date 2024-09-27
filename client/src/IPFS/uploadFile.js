import { useState } from "react";

function UploadFile({ setFileResponse, fileName }) {
  const [selectedFile, setSelectedFile] = useState();
  const changeHandler = (event) => {
    setSelectedFile(event.target.files);
  };

  const handleSubmission = async () => {
    const token = 
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiJlMjgwNmYyMy1lMWJiLTRlYjEtYTBmNS1kYzhmZTM4NzM3YTgiLCJlbWFpbCI6ImFrc2hheS5rdXJoZWthcjEwMTQ2NjJAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsInBpbl9wb2xpY3kiOnsicmVnaW9ucyI6W3siaWQiOiJGUkExIiwiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjF9XSwidmVyc2lvbiI6MX0sIm1mYV9lbmFibGVkIjpmYWxzZSwic3RhdHVzIjoiQUNUSVZFIn0sImF1dGhlbnRpY2F0aW9uVHlwZSI6InNjb3BlZEtleSIsInNjb3BlZEtleUtleSI6ImNjYjdmMDVmMmJiYzVlOTc0ZTc3Iiwic2NvcGVkS2V5U2VjcmV0IjoiYTkzMDhkNDU1YjZkN2RhMzM3NWEyYjIzY2I0OTVmMWUxNzJmNmJiMjdjMDQ0YjI4YWFkYjFjZGYzMGFmMGY2YiIsImlhdCI6MTcxODYwNTI5N30.3V5ruShYqMhibSuJxHeB2UFDggbRt5YrJFoz0hYBU-0";
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
