#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/proc_fs.h>
#include <linux/uaccess.h>
#include <linux/random.h>

#define PROC_NAME "rngen"

static ssize_t random_number_gen(struct file *file, char __user *buf, size_t count, loff_t *offset);

static struct proc_ops proc_fops = {
    .proc_read = random_number_gen,
};

// function for generating and returning a random number, used with profcs
// proc filesystem can be read multiple times by user-space -> offset keeps track -> no need for 'for' loop
static ssize_t random_number_gen(struct file *file, char __user *buf, size_t count, loff_t *offset) {
    
    unsigned int random_number;
    char buffer[32];
    ssize_t len;

    if (*offset > 0) {
      return 0;
    }

    get_random_bytes(&random_number, sizeof(random_number));
    // converts to string for profcs
    len = snprintf(buffer, sizeof(buffer), "%u\n", random_number);
    
    printk(KERN_INFO "rngen: Generated random number: %u\n", random_number);

    // check if all data is read
    if (*offset >= len) {
        return 0; // EOF
    }

    // copy data from kernel space (buffer) to user space (buf)
    if (copy_to_user(buf, buffer + *offset, len - *offset)) {
        return -EFAULT;
    }
    
    printk(KERN_INFO "rngen: Successfully returned random number to user\n");

    // update after one read 
    *offset += len;
    return len;
}

static int __init rn_init(void) {
    proc_create(PROC_NAME, 0, NULL, &proc_fops);
    printk(KERN_INFO "rngen: Loaded random number module.\n");
    return 0;
}

static void __exit rn_exit(void) {
    remove_proc_entry(PROC_NAME, NULL);
    printk(KERN_INFO "rngen: Exiting random number module.\n");
}

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("A module that generates random numbers and exposes them.");
MODULE_AUTHOR("Elena Krzina");

module_init(rn_init);
module_exit(rn_exit);
