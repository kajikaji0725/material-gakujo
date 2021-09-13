import React, { useContext } from "react";
import Box from "@mui/material/Box";
import { Button, Container, TextField, Typography } from "@mui/material";
import { makeStyles } from "@mui/styles";
import { ApiClientContext } from "../App";
import { Controller, SubmitHandler, useForm } from "react-hook-form";

const useStyles = makeStyles({
  paper: {
    width: "100%",
    display: "flex",
    flexDirection: "column",
  },
});

interface FormInput {
  gakujoUsername: string;
  gakujoPassword: string;
  email: string;
  username: string;
}
export function RegisterPage(): JSX.Element {
  const { control, handleSubmit } = useForm<FormInput>();
  const client = useContext(ApiClientContext);
  const classes = useStyles();
  const onSubmit: SubmitHandler<FormInput> = (data) => {
    client.register({
      gakujoUsername: data.gakujoUsername,
      gakujoPassword: data.gakujoPassword,
      email: data.email,
      username: data.username,
    });
  };

  return (
    <Container component="main" maxWidth="xs">
      <Box
        className={classes.paper}
        boxShadow={1}
        sx={{
          mt: "1em",
        }}
      >
        <Typography variant="h4" align="left">
          ユーザー登録
        </Typography>
        <Box
          component="form"
          onSubmit={handleSubmit(onSubmit)}
          sx={{
            width: "100%", // Fix IE 11 issue.
          }}
          display="flex"
          flexDirection="column"
          justifyContent="center"
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
                type="text"
                margin="normal"
              />
            )}
          />

          <Controller
            name="gakujoUsername"
            control={control}
            defaultValue=""
            render={({ field }) => (
              <TextField
                {...field}
                required
                label="学情の username(id)"
                type="text"
                margin="normal"
              />
            )}
          />

          <Controller
            name="gakujoPassword"
            control={control}
            defaultValue=""
            render={({ field }) => (
              <TextField
                {...field}
                required
                label="学情の password"
                type="password"
                margin="normal"
              />
            )}
          />

          <Controller
            name="email"
            control={control}
            defaultValue=""
            render={({ field }) => (
              <TextField
                {...field}
                required
                label="メールアドレス"
                type="email"
                fullWidth
                margin="normal"
              />
            )}
          />
          <Button
            type="submit"
            variant="contained"
            color="primary"
            sx={{
              mb: "1em",
              mt: "1em",
              width: "200px",
            }}
          >
            登録する
          </Button>
        </Box>
      </Box>
    </Container>
  );
}
