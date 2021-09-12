import React, { useContext } from "react";
import Box from "@mui/material/Box";
import { Button, Container, TextField } from "@mui/material";
import { makeStyles } from "@mui/styles";
import { ApiClientContext } from "../App";
import { Controller, SubmitHandler, useForm } from "react-hook-form";

const useStyles = makeStyles({
  paper: {
    width: "100%",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  form: {
    alignItems: "center",
  },
});

interface FormInput {
  username: string;
  password: string;
}
export function SignInPage(): JSX.Element {
  const { control, handleSubmit } = useForm<FormInput>();
  const client = useContext(ApiClientContext);
  const classes = useStyles();
  const onSubmit: SubmitHandler<FormInput> = (data) => {
    client.login(data.username, data.password);
  };

  return (
    <Container component="main" maxWidth="xs">
      <Box className={classes.paper}>
        <Box
          component="form"
          onSubmit={handleSubmit(onSubmit)}
          m={2}
          sx={{
            width: "100%", // Fix IE 11 issue.
          }}
          className={classes.form}
          textAlign="center"
        >
          <Controller
            name="username"
            control={control}
            defaultValue=""
            render={({ field }) => (
              <TextField
                {...field}
                required
                label="username"
                type="username"
                fullWidth
                margin="normal"
              />
            )}
          />

          <Controller
            name="password"
            control={control}
            defaultValue=""
            render={({ field }) => (
              <TextField
                {...field}
                required
                label="password"
                type="password"
                fullWidth
                margin="normal"
              />
            )}
          />

          <Button type="submit" variant="contained" color="primary">
            login
          </Button>
        </Box>
      </Box>
    </Container>
  );
}
