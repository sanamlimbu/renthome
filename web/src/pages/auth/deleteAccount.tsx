import { EmailOutlined } from "@mui/icons-material";
import {
  Alert,
  Button,
  Card,
  InputAdornment,
  TextField,
  Typography,
} from "@mui/material";
import { useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import RentHomeLogo from "../../components/rentHomeLogo";
import { API_ADDRESS } from "../../config";
import { getTokenFromLocalStorage } from "../../helpers/auth";
import { ErrorResponse } from "../../types/types";
import "../styles/index.css";

interface IDeleteAccountInput {
  confirm_email: string;
}

export default function DeleteAccountPage() {
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<IDeleteAccountInput>();
  const navigate = useNavigate();
  const [error, setError] = useState("");

  const onSubmit: SubmitHandler<IDeleteAccountInput> = async (input) => {
    try {
      const res = await fetch(`${API_ADDRESS}/api/auth/email-update`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${getTokenFromLocalStorage()}`,
        },
        body: JSON.stringify(input),
      });
      if (res.ok) {
        const data = res.json();
        console.log(data);
        navigate("/me");
      } else {
        const data: ErrorResponse = await res.json();
        setError(data.message);
      }
    } catch (error) {}
  };

  return (
    <div className="Container-centered">
      <Card variant="outlined" className="Card">
        <RentHomeLogo />
        <Typography textAlign={"center"} fontWeight={"700"} fontSize={"18px"}>
          Update your email address
        </Typography>
        <form onSubmit={handleSubmit(onSubmit)} className="Form">
          {error && <Alert severity="warning">{error}</Alert>}
          <TextField
            variant="outlined"
            placeholder="New email address"
            {...register("confirm_email", {
              required: {
                value: true,
                message: "Please enter a valid email address.",
              },
            })}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <EmailOutlined />
                </InputAdornment>
              ),
            }}
          />
          {errors.confirm_email && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.confirm_email.message}
            </span>
          )}
          <Button
            variant="contained"
            size="large"
            sx={{ textTransform: "none", fontWeight: "600" }}
            type="submit"
          >
            Next
          </Button>
        </form>
        <div style={{ paddingTop: "14px", paddingBottom: "14px" }}>
          <Button
            sx={{ textTransform: "none", fontWeight: "700" }}
            fullWidth
            onClick={() => navigate("/me")}
          >
            Cancel
          </Button>
        </div>
      </Card>
    </div>
  );
}
