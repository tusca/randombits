# DFU Restore — MacBook Pro (Apple Silicon: M1, M2, M3, M4)

A DFU (Device Firmware Update) restore rewrites the Mac's firmware and internal storage at a level **below** macOS and FileVault. It is the reliable way to wipe an Apple silicon MacBook Pro when:

- The FileVault recovery key is lost and not escrowed in MDM/ABM.
- The Mac is stuck in a boot loop or bricked after a failed update.
- You need a clean factory state regardless of account passwords or disk encryption.

> **Important:** A DFU **Restore** erases everything on the target Mac, including the internal SSD and all user data. A DFU **Revive** attempts to fix firmware without erasing user data — use it first if data preservation matters.

---

## About recovery keys and ABM

- **Apple Business Manager (ABM) does NOT store FileVault recovery keys.** ABM only holds device assignment (which MDM the Mac auto-enrolls into) and, for supervised devices, Activation Lock bypass codes.
- FileVault Personal Recovery Keys (PRKs) are escrowed only in the MDM (e.g., Jamf Pro) — if deleted there, they are gone.
- DFU Restore bypasses FileVault because it reformats the storage controller itself; no key is required.

---

## Prerequisites

| Item | Notes |
|---|---|
| **Host Mac** | A second Mac running macOS 12.4 or later (macOS 13+ recommended). |
| **Apple Configurator 2** | Free on the Mac App Store. Install on the host Mac. |
| **USB-C cable** | Must be **charge + data** capable. Apple's supplied USB-C charge cable works. Thunderbolt 3/4 cables work. **Charge-only cables will silently fail** — this is the #1 cause of DFU problems. |
| **Power** | Both Macs plugged into AC power. |
| **Target Mac** | MacBook Pro with Apple silicon (M1 / M1 Pro / M1 Max / M2 / M2 Pro / M2 Max / M3 / M3 Pro / M3 Max / M4 / M4 Pro / M4 Max). |

---

## Which port to use on the target MacBook Pro

On **all Apple silicon MacBook Pro models (M1–M4, 13" / 14" / 16")**, use:

> **The Thunderbolt / USB 4 port closest to the display, on the LEFT side of the laptop.**

- 13" MacBook Pro (M1, M2): left-side port closest to the screen hinge.
- 14" / 16" MacBook Pro (M1 Pro/Max, M2 Pro/Max, M3 family, M4 family): left-side port closest to the screen hinge (the MagSafe port is separate — do not use it for DFU; use a Thunderbolt port).

On the **host Mac**, any Thunderbolt / USB-C port is fine.

---

## Step 1 — Prepare both Macs

1. On the host Mac, install **Apple Configurator 2** from the Mac App Store and launch it at least once (accept permissions prompts).
2. Plug the host Mac into power.
3. On the target MacBook Pro: **shut it down fully** (Apple menu → Shut Down, or hold the power button for ~10 seconds if unresponsive).
4. Disconnect the target from power.
5. Connect the USB-C cable:
   - One end into the **correct DFU port** on the target (left side, closest to display).
   - Other end into any USB-C / Thunderbolt port on the host.
6. Reconnect the target to power. (Power can flow via the MagSafe port or through the Thunderbolt cable — either works.)

---

## Step 2 — Enter DFU mode (MacBook Pro, Apple silicon)

The target's screen **stays completely black** throughout. No Apple logo, no chime. That is correct — if you see an Apple logo, you went too far; shut down and retry.

1. Make sure the target is shut down and connected as described above.
2. Press and hold **all four** of the following keys **simultaneously** for about **10 seconds**:
   - **Right Shift**
   - **Left Option** (⌥)
   - **Left Control** (⌃)
   - **Power** button
3. After ~10 seconds, **release the three modifier keys** (Right Shift, Left Option, Left Control) — but **keep holding the Power button** for about another **10 seconds**.
4. On the host Mac, Apple Configurator 2 should display a device labeled **"DFU"**.

If the DFU icon does not appear within ~30 seconds:

- Unplug everything, shut the target down, and try again.
- Swap the USB-C cable for a known-good Apple cable.
- Confirm you are using the left-side port closest to the display on the target.
- Confirm Apple Configurator 2 is open and focused on the host.

---

## Step 3 — Revive (preserve data) or Restore (erase)

### Option A — Revive (try this first if you want to preserve data)

Revive reinstalls the firmware and the latest recoveryOS **without** erasing the user's data volume. Use this for boot/firmware failures when data should be kept.

1. In Apple Configurator 2, right-click the DFU device.
2. Choose **Advanced → Revive Device**.
3. Click **Revive** to confirm.
4. Wait 10–20 minutes. The target may reboot several times. Do not disconnect.
5. When finished, the target boots normally (or to recoveryOS) with user data intact.

### Option B — Restore (erases everything, required for unknown FileVault key)

Restore reinstalls firmware **and reformats the internal SSD**, erasing all data including FileVault-encrypted volumes. This is the step required when the recovery key is unknown.

1. In Apple Configurator 2, right-click the DFU device.
2. Choose **Restore** (or **Actions → Restore**).
3. Confirm when prompted. **All data on the target will be erased.**
4. Wait 15–45 minutes. The target will reboot multiple times. Do not disconnect the cable, do not close the lid, do not let either Mac sleep.
5. When complete, the target boots to **Hello / Setup Assistant**.

---

## Step 4 — After the restore

If the MacBook Pro is assigned to your organization in **Apple Business Manager** and mapped to an MDM server (e.g., Jamf Pro):

1. In Setup Assistant, select language and region.
2. Connect to Wi-Fi (or Ethernet via adapter).
3. **Automated Device Enrollment** should trigger automatically — the Mac contacts Apple, learns which MDM to enroll in, and downloads the enrollment profile.
4. Continue through Setup Assistant and enrollment as configured by your MDM (local account creation, FileVault re-enablement, etc.).

If auto-enrollment does not trigger:

- Verify the device's serial number is still assigned to your MDM server in ABM.
- Verify the MDM server's Automated Device Enrollment token is valid (not expired) in ABM.
- Check the MDM for a PreStage Enrollment / ADE configuration that applies to this device.

---

## Troubleshooting

| Symptom | Likely cause / fix |
|---|---|
| DFU icon never appears in Configurator 2 | Cable is charge-only — swap to a known-good USB-C data cable or Thunderbolt cable. |
| Apple logo appears on target during key combo | Held keys too long or wrong sequence — shut down and retry from step 2. |
| Configurator shows device but Restore fails mid-way | Network issue on host Mac (Configurator downloads firmware from Apple). Ensure stable internet and retry. |
| Restore fails with error 4000-series | Usually cable/port — try a different Apple-branded cable, retry from cold. |
| Target boots to Activation Lock screen after restore | Device is locked to an Apple ID. If the Mac is in ABM and supervised, clear Activation Lock via ABM (Devices → select → Disable Activation Lock). |
| Setup Assistant does not show MDM enrollment | Confirm ABM assignment and MDM ADE token; sometimes takes a few minutes after Wi-Fi connects. |

---

## Key safety notes

- Do **not** disconnect the cable or power during Revive/Restore — interrupting can brick firmware and require another DFU attempt.
- DFU Restore is **destructive and irreversible** — double-check you have the right target Mac before confirming.
- The host Mac downloads the IPSW firmware from Apple during Restore (~15 GB). Ensure sufficient disk space on the host and a stable internet connection.
- Apple Silicon DFU does not require a DFU-specific IPSW file by default — Configurator 2 fetches the latest signed firmware automatically. You can supply a specific IPSW via Option-click on Restore if needed.

---

## References

- Apple Configurator User Guide — "Revive or restore a Mac with Apple silicon": https://support.apple.com/guide/apple-configurator-2/welcome/mac
- Apple Support — "If you can't turn on your Mac": https://support.apple.com/HT204267
- Apple Business Manager User Guide: https://support.apple.com/guide/apple-business-manager/welcome/web
