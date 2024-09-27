import { useState } from "react";
import { Form } from "react-bootstrap";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import DropdownList from "./dropdown";
import UploadFile from "../IPFS/uploadFile";
import axios from "axios";
import { createCredential } from "../api";

function AddCredential() {
  const [show, setShow] = useState(false);
  const [cname, setCName] = useState("");
  const [type, setType] = useState(null);
  const [vaultId, setVaultId] = useState();
  const [fileResponse, setFileResponse] = useState(null);

  const handleClose = () => {
    setShow(false);
    setType(null);
  };
  const handleShow = () => {
    setShow(true);
    setType(null);
  };
  let AA = null;

  const selectInput = () => {
    switch (type) {
      case "text":
        AA = (
          <div style={{ padding: "5px", marginTop: "10px" }}>
            <textarea
              style={{ width: "100%", height: "90px" }}
              placeholder="Enter your credential content!"
            ></textarea>
          </div>
        );
        break;
      case "file":
        AA = <UploadFile setFileResponse={setFileResponse} fileName={cname} />;
        break;
      default:
        AA = <div style={{ margin: "5px" }}>not selected</div>;
        break;
    }
    return AA;
  };

  const saveCredential = async () => {
    //   {
    //     "IpfsHash": "QmfCSPUrHfeCt7ZgDztzXvddjZX15BGeZxXbqBCDvWMoHj",
    //     "PinSize": 15580,
    //     "Timestamp": "2024-07-15T02:11:07.413Z"
    // }
    console.log("CID:", fileResponse.IpfsHash);
    const data = {
      "VId": parseInt(vaultId),
      "Credential": { "Name": cname, "Cid": fileResponse.IpfsHash },
    };

    try {
      // const res = await axios.post("http://localhost:4000/addCredential", data);
      const result = await createCredential(data);
      console.log("Credential Created db res:", result);
      // setCredentials(res.data); // Uncomment this if setCredentials is defined
      setShow(false);
    } catch (err) {
      console.log("ERROR in creating vault:", err);
    }
  };

  return (
    <>
      <Button variant="info" style={{ margin: "5px" }} onClick={handleShow}>
        Add Credential in Vault
      </Button>

      <Modal
        show={show}
        onHide={handleClose}
        backdrop="static"
        keyboard={false}
      >
        <Modal.Header closeButton>
          <Modal.Title>Add Credential</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <div style={{ padding: "10px", color: "green" }}>
            <label>Credential Name: </label>
            <input
              type="text"
              placeholder="Enter Name"
              onChange={(e) => {
                setCName(e.target.value);
              }}
            />
          </div>
          <DropdownList setVaultId={setVaultId} />
          <Form.Select
            size="sm"
            style={{ marginTop: "10px" }}
            id="dropdown-basic-button"
            variant="light"
            title="Select Type"
            onChange={(e) => {
              setType(e.target.value);
              console.log("selected value", e.target.value);
            }}
          >
            <option value="0">Select Type </option>
            <option value="text">{`Text - .txt`}</option>
            <option value="file">{`File - .png/.mp3/.mp4`} </option>
            <option value="folder">{`Folder`} </option>
          </Form.Select>
          {selectInput()}
        </Modal.Body>
        <Modal.Footer>
          <Button size="sm" variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button size="sm" variant="success" onClick={saveCredential}>
            Save
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
}

export default AddCredential;
