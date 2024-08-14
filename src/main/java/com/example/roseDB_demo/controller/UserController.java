package com.example.roseDB_demo.controller;



import com.example.roseDB_demo.model.User;
import com.example.roseDB_demo.service.UserService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api")
@CrossOrigin(origins = "*")

public class UserController {

    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    @PostMapping("/user")
    public ResponseEntity<String> createUser(@RequestBody User user) {
        return userService.createUser(user);
    }

    @GetMapping("/user")
    public ResponseEntity<String> getUser(@RequestParam String email) {
        return userService.getUser(email);
    }
}
