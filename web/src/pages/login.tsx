import { Button, Divider, Link, TextField, Typography } from "@mui/material";
import Card from "@mui/material/Card";
import React from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useLocation } from "react-router-dom";
import Privacy from "../components/privacy";
import RentHomeLogo from "../components/rentHomeLogo";
import Social, { SocialAction, socialList } from "../components/social";
import "../styles/index.css";

interface ILoginInput {
  email: string;
  password: string;
}

export default function LoginPage() {
  const location = useLocation();
  let from = ((location.state as any)?.from?.pathname as string) || "/";

  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<ILoginInput>();
  const [open, setOpen] = React.useState(false);

  const onSubmit: SubmitHandler<ILoginInput> = (data) => {
    console.log(data);
  };

  return (
    <div className="Container-centered">
      <Card variant="outlined" className="Card">
        <RentHomeLogo />
        <Typography textAlign={"center"} fontWeight={"700"} fontSize={"18px"}>
          Sign in
        </Typography>
        <form onSubmit={handleSubmit(onSubmit)} className="Form">
          <TextField
            variant="outlined"
            placeholder="Email address"
            {...register("email", {
              required: {
                value: true,
                message: "Please enter a valid email address.",
              },
            })}
          />
          {errors.email && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.email.message}
            </span>
          )}
          <TextField
            variant="outlined"
            placeholder="Password"
            type={"password"}
            {...register("password", {
              required: {
                value: true,
                message: "Please enter a password.",
              },
            })}
          />
          {errors.password && (
            <span style={{ color: "red", fontSize: "14px" }}>
              {errors.password.message}
            </span>
          )}
          <Button
            variant="contained"
            size="large"
            sx={{ textTransform: "none", fontWeight: "600" }}
            type="submit"
          >
            Sign in
          </Button>
        </form>
        <Link color="primary" className="Link" href="/forgot-password">
          Forgot your password?
        </Link>
        <Divider sx={{ fontWeight: "700" }}>OR</Divider>
        {socialList.map((s) => (
          <Social
            key={s.name}
            type={s}
            action={SocialAction.signin}
            from={from}
          />
        ))}
        <Typography textAlign={"center"}>
          Not signed up?{" "}
          <Link className="Link" href={"/signup"}>
            Create an account.
          </Link>
        </Typography>
        <Divider />
        <Typography
          sx={{ textAlign: "center", fontSize: "12px", cursor: "pointer" }}
          onClick={() => setOpen(true)}
        >
          Personal Information Collection Statement
        </Typography>
      </Card>
      <Privacy open={open} handleClose={() => setOpen(false)} />
    </div>
  );
}
