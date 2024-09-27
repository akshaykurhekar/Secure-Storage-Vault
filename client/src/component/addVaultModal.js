import { useState } from "react";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import axios from "axios";
import {createVaultApi} from "../api"

function AddVault() {
  const [show, setShow] = useState(false);
  const [name, setName] = useState();
  const [desc, setDesc] = useState();
  const [pass, setPass] = useState();

  const handleClose = () => {
    setShow(false);
  };
  const handleShow = () => {
    setShow(true);
  };

  const createVault = async () => {
    const data = {
      "Name": name,
      "Password": pass,
      "Desc":desc
    };

    try {
      // const res = await axios.post("http://localhost:5000/create/vault", data);
      const res = createVaultApi(data);
      
      console.log("Vault res", res);
      // setCredentials(res.data); // Uncomment this if setCredentials is defined
      setShow(false);
    } catch (err) {
      console.log("ERROR in creating vault:", err);
    }
  };

  return (
    <>
      <Button variant="warning" style={{ margin: "5px" }} onClick={handleShow}>
        Create Vault
      </Button>

      <Modal
        show={show}
        onHide={handleClose}
        backdrop="static"
        keyboard={false}
      >
        <Modal.Header closeButton>
          <Modal.Title>create Vault</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          Name:{" "}
          <input
            style={{ margin: "10px" }}
            type="text"
            onChange={(e) => {
              setName(e.target.value);
            }}
            placeholder="Enter Vault name.."
          ></input>
          <br></br>
          Description:{" "}
          <input
            style={{ margin: "10px" }}
            type="text"
            onChange={(e) => {
              setDesc(e.target.value);
            }}
            placeholder="Enter Vault desc..!"
          ></input>
          <br></br>
          Password:{" "}
          <input
            style={{ margin: "10px" }}
            type="text"
            onChange={(e) => {
              setPass(e.target.value);
            }}
            placeholder="Create Vault Password."
          ></input>
        </Modal.Body>
        <Modal.Footer>
          <Button size="sm" variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button size="sm" variant="success" onClick={createVault}>
            Save
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
}

export default AddVault;
