import { useEffect, useState } from "react";
import axios from "axios";
import { Form } from "react-bootstrap";
import { getAllVault } from "../api";

function DropdownList({ setVaultId }) {
  //display vault list here and select in which vault you want to add credential.

  const [vaultList, setVaultList] = useState();

  useEffect( () => { 
    async function getData() {
      const result = await getAllVault();
      console.log(result)
      setVaultList(result);
    }
    getData()

    // axios
    //   .get("http://localhost:4000/users")
    //   .then((res) => {
    //     console.log(res.data);
    //     setVaultList(res.data);
    //   })
    //   .catch((err) => {
    //     console.log("ERROR in getting Users list:", err);
    //   });
    console.log(vaultList);
  }, []);

  return (
    <Form.Select
      size="sm"
      onChange={(e) => {
        setVaultId(e.target.value);
        console.log("Vault Uid:", e.target.value);
      }}
    >
      <option key={0}>Select Vault</option>
      {vaultList != null ? (
        vaultList.map((item) => {
          return (
            <option key={item.id} value={item.id}>
              {item.name}
            </option>
          );
        })
      ) : (
        <p>No Vault Created! Create Vault...</p>
      )}
    </Form.Select>
  );
}

export default DropdownList;
