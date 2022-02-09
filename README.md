# Git User Manager (GUM)

It is the project that enables to switch between Git users via terminal. It just uses Git commands and does not directly access Git files (like config file) for any operation.

## Installation

- Clone this repo to your PC:

```git
git clone https://github.com/mtyuksel/gum.git
```

- Open your zshrc/bashrc file with any editor:

```
code ~/.zshrc
```

- Add following with the correct directory at the and of the zshrc/bashrc file:
```
export PATH=$PATH:<THE_DIRECTORY_YOU_CLONED_GUM_INTO>/gum
```

- And now, you can reach GUM in any directory. Check with following (you will see default profile names):

![image](https://user-images.githubusercontent.com/60136172/153279599-cf39838b-8bcb-446c-914f-6613cf7c17b8.png)

## Get Started

- Add your profile info to config.yaml

![image](https://user-images.githubusercontent.com/60136172/153280373-8fe5bb4d-4071-4cf5-925f-d462824531b9.png)

- You can set your user profiles with following command:
```
gum set <profile-name>
```
![image](https://user-images.githubusercontent.com/60136172/153279918-46d21fa6-2482-4b89-a003-6e16dbf4f8fa.png)

- You can check your user profile names with following command:
```
gum show
```
![image](https://user-images.githubusercontent.com/60136172/153279599-cf39838b-8bcb-446c-914f-6613cf7c17b8.png)

